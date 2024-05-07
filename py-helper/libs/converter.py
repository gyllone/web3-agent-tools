from typing import Type, Any, Optional, Tuple, Mapping
from ctypes import c_bool, c_longlong, c_double, c_char_p, c_size_t, Structure, POINTER, cast

from .schema import Object, Property, ParamSchema


CType = Type[c_bool | c_longlong | c_double | c_char_p | Structure]
PyValue = bool | int | float | str | list[Any] | dict[str, Any] | Mapping[str, Any] | Optional[Any]
CValue = c_bool | c_longlong | c_double | c_char_p | Structure


class TypeConverter(object):
    scope: str
    schema: ParamSchema
    c_types_cache: dict[str, Type[Structure]] = {}

    def __init__(self, scope: str, schema: ParamSchema):
        self.scope = scope
        self.schema = schema

    def get_arg_types(self) -> list[CType]:
        return [item for _, item in self._get_c_fields(self.schema)]

    def get_structure_type(self) -> Type[Structure]:
        c_fields = self._get_c_fields(self.schema)
        return self._wrap_struct(f"{self.scope}_c_root_structure", c_fields)

    def _get_c_fields(self, obj: Object) -> list[Tuple[str, CType]]:
        c_types = []
        required = obj.required or []
        properties = obj.properties or {}
        for key, prop in properties.items():
            c_type = self._get_c_type(prop)
            # optional = not required and default is None
            if key not in required and prop.default is None:
                c_type = self._wrap_optional(c_type)
            c_types.append((key, c_type))
        return c_types

    def _get_c_type(self, prop: Property) -> CType:
        if prop.allOf:
            ref_type = prop.allOf[0].get_ref_type()
        else:
            ref_type = prop.get_ref_type()
        if ref_type is not None:
            objects = self.schema.definitions or {}
            obj = objects.get(ref_type)
            if not obj:
                raise TypeError(f"Ref type {ref_type} not found in objects")
            c_fields = self._get_c_fields(obj)
            return self._wrap_struct(ref_type, c_fields)

        py_type = prop.parse_type()
        if not py_type:
            raise TypeError("Must define type")
        if py_type is str:
            return c_char_p
        elif py_type is int:
            return c_longlong
        elif py_type is float:
            return c_double
        elif py_type is bool:
            return c_bool
        elif py_type is list:
            if not prop.items:
                raise TypeError("Must define item of list")
            inner_c_type = self._get_c_type(prop.items)
            return self._wrap_list(inner_c_type)
        elif py_type is dict:
            if not prop.additionalProperties:
                raise TypeError("Must define additional property of dict")
            inner_c_type = self._get_c_type(prop.additionalProperties)
            return self._wrap_dict(inner_c_type)

    def _wrap_struct(self, name: str, c_fields: list[Tuple[str, CType]]) -> Type[Structure]:
        name = f"{self.scope}_c_structure_{name}"
        structure = self.c_types_cache.get(name)
        if not structure:
            structure = type(
                name,
                (Structure,),
                {
                    # "_pack_": 1,
                    "_fields_": c_fields,
                },
            )
            self.c_types_cache[name] = structure  # type: ignore
        return structure  # type: ignore

    def _wrap_list(self, c_type: CType) -> Type[Structure]:
        name = f"{self.scope}_c_list_{c_type.__name__}"
        clist = self.c_types_cache.get(name)
        if not clist:
            clist = type(
                name,
                (Structure,),
                {
                    # "_pack_": 1,
                    "_fields_": [
                        ("len", c_size_t),
                        ("values", POINTER(c_type)),
                    ]
                },
            )
            self.c_types_cache[name] = clist  # type: ignore
        return clist  # type: ignore

    def _wrap_dict(self, c_type: CType) -> Type[Structure]:
        name = f"{self.scope}_c_dict_{c_type.__name__}"
        cdict = self.c_types_cache.get(name)
        if not cdict:
            cdict = type(
                name,
                (Structure,),
                {
                    # "_pack_": 1,
                    "_fields_": [
                        ("len", c_size_t),
                        ("keys", POINTER(c_char_p)),
                        ("values", POINTER(c_type)),
                    ]
                },
            )
            self.c_types_cache[name] = cdict  # type: ignore
        return cdict  # type: ignore

    def _wrap_optional(self, c_type: CType) -> Type[Structure]:
        name = f"{self.scope}_c_optional_{c_type.__name__}"
        c_optional = self.c_types_cache.get(name)
        if not c_optional:
            c_optional = type(
                name,
                (Structure,),
                {
                    # "_pack_": 1,
                    "_fields_": [
                        ("is_some", c_bool),
                        ("value", c_type),
                    ]
                },
            )
            self.c_types_cache[name] = c_optional  # type: ignore
        return c_optional  # type: ignore


class ValueConverter(TypeConverter):
    def py_args_to_c_values(self, **kwargs) -> list[CValue]:
        c_values = self._py_values_to_c_values(self.schema, kwargs)
        return list(c_values.values())

    def c_struct_to_py_object(self, structure: Structure) -> Mapping[str, Any]:
        return self._c_struct_to_py_object(self.schema, structure)

    def _c_value_to_py_value(self, prop: Property, c_value: CValue) -> PyValue:
        if prop.allOf:
            ref_type = prop.allOf[0].get_ref_type()
        else:
            ref_type = prop.get_ref_type()
        if ref_type is not None:
            objects = self.schema.definitions or {}
            obj = objects.get(ref_type)
            if not obj:
                raise TypeError(f"Ref type {ref_type} not found in objects")
            return self._c_struct_to_py_object(obj, c_value)

        py_type = prop.parse_type()
        if not py_type:
            raise TypeError("Must define type")
        if py_type is str:
            if c_value is None:
                return ""
            elif isinstance(c_value, bytes):
                return c_value.decode("utf-8")
            else:
                raise TypeError(f"Expected c_char_p, got {type(c_value)}")
        elif py_type in (int, float, bool):
            return py_type(c_value)
        elif py_type is list:
            if not prop.items:
                raise TypeError("Must define item of list")
            return self._c_list_to_py_list(prop.items, c_value)
        elif py_type is dict:
            if not prop.additionalProperties:
                raise TypeError("Must define additional property of dict")
            return self._c_dict_to_py_dict(prop.additionalProperties, c_value)

    def _py_value_to_c_value(self, prop: Property, py_value: PyValue) -> CValue:
        if prop.allOf:
            ref_type = prop.allOf[0].get_ref_type()
        else:
            ref_type = prop.get_ref_type()
        if ref_type is not None:
            objects = self.schema.definitions or {}
            obj = objects.get(ref_type)
            if not obj:
                raise TypeError(f"Ref type {ref_type} not found in objects")
            return self._py_object_to_c_struct(ref_type, obj, py_value)

        py_type = prop.parse_type()
        if not py_type:
            raise TypeError("Must define type")
        if py_type is str:
            return c_char_p(py_value.encode("utf-8"))
        elif py_type is float:
            return c_double(py_value)
        elif py_type is int:
            return c_longlong(py_value)
        elif py_type is bool:
            return c_bool(py_value)
        elif py_type is list:
            if not prop.items:
                raise TypeError("Must define item of list")
            return self._py_list_to_c_list(prop.items, py_value)
        elif py_type is dict:
            if not prop.additionalProperties:
                raise TypeError("Must define additional property of dict")
            return self._py_dict_to_c_dict(prop.additionalProperties, py_value)

    def _c_struct_to_py_object(self, obj: Object, c_struct: Structure) -> Mapping[str, Any]:
        _object = {}
        required = obj.required or []
        properties = obj.properties or {}
        for key, prop in properties.items():
            if key not in required:
                if prop.default is None:
                    # optional
                    c_value = getattr(c_struct, key)
                    _object[key] = self._c_optional_to_py_optional(prop, c_value)
                else:
                    _object[key] = prop.default
            else:
                c_value = getattr(c_struct, key)
                _object[key] = self._c_value_to_py_value(prop, c_value)
        return _object

    def _py_values_to_c_values(self, obj: Object, py_values: Mapping[str, PyValue]) -> Mapping[str, CValue]:
        c_values = {}
        required = obj.required or []
        properties = obj.properties or {}
        for key, prop in properties.items():
            py_value = py_values.get(key)
            if key not in required:
                if prop.default is None:
                    # optional
                    c_values[key] = self._py_optional_to_c_optional(prop, py_value)
                else:
                    c_values[key] = self._py_value_to_c_value(prop, py_value or prop.default)
            else:
                if py_value is None:
                    raise ValueError(f"Missing required field {key}")
                c_values[key] = self._py_value_to_c_value(prop, py_value)

        return c_values

    def _py_object_to_c_struct(self, name: str, obj: Object, py_object: Mapping[str, Any]) -> Structure:
        c_fields = self._get_c_fields(obj)
        c_struct = self._wrap_struct(name, c_fields)
        structure = self._py_values_to_c_values(obj, py_object)
        return c_struct(**structure)

    def _c_list_to_py_list(self, prop: Property, c_list: Structure) -> list[Any]:
        c_type = self._get_c_type(prop)
        values_pointer = cast(c_list.values, POINTER(c_type * c_list.len))
        return [
            self._c_value_to_py_value(prop, values_pointer.contents[i])
            for i in range(int(c_list.len))
        ]

    def _py_list_to_c_list(self, prop: Property, py_list: list[Any]) -> Structure:
        c_type = self._get_c_type(prop)
        c_list = self._wrap_list(c_type)
        _len = len(py_list)
        return c_list(
            len=_len,
            values=(c_type * _len)(*[self._py_value_to_c_value(prop, py_value) for py_value in py_list])
        )

    def _c_dict_to_py_dict(self, prop: Property, c_dict: Structure) -> dict[str, Any]:
        c_type = self._get_c_type(prop)
        keys_pointer = cast(c_dict.keys, POINTER(c_char_p * c_dict.len))
        values_pointer = cast(c_dict.values, POINTER(c_type * c_dict.len))
        return {
            keys_pointer.contents[i].decode("utf-8"):
                self._c_value_to_py_value(prop, values_pointer.contents[i])
            for i in range(int(c_dict.len))
        }

    def _py_dict_to_c_dict(self, prop: Property, py_dict: dict[str, Any]) -> Structure:
        c_type = self._get_c_type(prop)
        c_dict = self._wrap_dict(c_type)
        _len = len(py_dict)
        return c_dict(
            len=_len,
            keys=(c_char_p * _len)(*[c_char_p(key.encode("utf-8")) for key in py_dict.keys()]),
            values=(c_type * _len)(*[self._py_value_to_c_value(prop, py_value) for py_value in py_dict.values()])
        )

    def _c_optional_to_py_optional(self, prop: Property, c_optional: Structure) -> Optional[Any]:
        return self._c_value_to_py_value(prop, c_optional.value) if c_optional.is_some else None

    def _py_optional_to_c_optional(self, prop: Property, py_optional: Optional[Any]) -> Structure:
        c_type = self._get_c_type(prop)
        c_optional = self._wrap_optional(c_type)
        if py_optional is None:
            return c_optional(is_some=False)
        else:
            return c_optional(is_some=True, value=self._py_value_to_c_value(prop, py_optional))
