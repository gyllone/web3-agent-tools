import inspect
import typing

from typing import Type, List, Any, Dict, Optional, get_origin, get_args
from ctypes import c_bool, c_longlong, c_double, c_char_p, Structure, c_size_t, cast, byref, POINTER
from pydantic import BaseModel


CType = Type[c_bool | c_longlong | c_double | c_char_p | Any]
PyType = Type[bool | int | float | str | BaseModel | List[Any] | Dict[str, Any] | Optional[Any]]

registered_c_types = {}


def struct_type_convert_py_to_c(model_type: Type[BaseModel]) -> Type[Structure]:
    name = f"c_structure_{model_type.__name__}"
    structure = registered_c_types.get(name)
    if not structure:
        structure = type(
            name,
            (Structure,),
            {
                "_fields_": [
                    (name, type_convert_py_to_c(field.annotation))
                    for name, field in model_type.__fields__.items()
                ]
            },
        )
        registered_c_types[name] = structure
    return structure  # type: ignore


def wrap_list_type_to_c_struct(c_type: CType) -> Type[Structure]:
    name = f"c_list_{c_type.__name__}"
    clist = registered_c_types.get(name)
    if not clist:
        clist = type(
            name,
            (Structure,),
            {
                "_fields_": [
                    ("len", c_size_t),
                    ("values", POINTER(c_type)),
                ]
            },
        )
        registered_c_types[name] = clist
    return clist  # type: ignore


def wrap_optional_type_to_c_struct(c_type: CType) -> Type[Structure]:
    name = f"c_optional_{c_type.__name__}"
    c_optional = registered_c_types.get(name)
    if not c_optional:
        c_optional = type(
            name,
            (Structure,),
            {
                "_fields_": [
                    ("is_some", c_bool),
                    ("value", c_type),
                ]
            },
        )
        registered_c_types[name] = c_optional
    return c_optional  # type: ignore


def wrap_dict_type_to_c_struct(c_type: CType) -> Type[Structure]:
    name = f"c_dict_{c_type.__name__}"
    cdict = registered_c_types.get(name)
    if not cdict:
        cdict = type(
            name,
            (Structure,),
            {
                "_fields_": [
                    ("len", c_size_t),
                    ("keys", POINTER(c_char_p)),
                    ("values", POINTER(c_type)),
                ]
            },
        )
        registered_c_types[name] = cdict
    return cdict  # type: ignore


def type_convert_py_to_c(py_type: PyType) -> CType:
    if inspect.isclass(py_type):
        if py_type is bool:
            return c_bool
        elif py_type is int:
            return c_longlong
        elif py_type is float:
            return c_double
        elif py_type is str:
            return c_char_p
        elif issubclass(py_type, BaseModel):
            return POINTER(struct_type_convert_py_to_c(py_type))
        else:
            raise TypeError(f"Unsupported python class type {py_type}")

    origin = get_origin(py_type)
    args = get_args(py_type)
    if origin is list:
        c_type = type_convert_py_to_c(args[0])
        return POINTER(wrap_list_type_to_c_struct(c_type))
    elif origin is dict:
        if args[0] is str:
            c_type = type_convert_py_to_c(args[1])
            return POINTER(wrap_dict_type_to_c_struct(c_type))
        else:
            raise TypeError(f"Unsupported key type {args[0]}")
    elif origin is typing.Union:
        if len(args) == 2 and args[1] is type(None):
            c_type = type_convert_py_to_c(args[0])
            return POINTER(wrap_optional_type_to_c_struct(c_type))
        else:
            raise TypeError(f"Unsupported type {py_type}")
    else:
        raise TypeError(f"Unsupported type {py_type}")


def arg_types_convert_py_to_c(model_type: Type[BaseModel]) -> List:
    return [
        type_convert_py_to_c(field.annotation)
        for name, field in model_type.__fields__.items()
    ]


def value_convert_c_to_py(py_type: PyType, c_value: Any) -> Any:
    if inspect.isclass(py_type):
        if py_type in (bool, int, float):
            return c_value
        elif py_type is str:
            return str(c_value.decode("utf-8"))
        elif issubclass(py_type, BaseModel):
            return struct_value_convert_c_to_py(py_type, c_value.contents)
        else:
            raise TypeError(f"Unsupported python class type {py_type}")

    origin = get_origin(py_type)
    args = get_args(py_type)
    if origin is list:
        return list_value_convert_c_to_py(args[0], c_value.contents)
    elif origin is dict:
        if args[0] is str:
            return dict_value_convert_c_to_py(args[1], c_value.contents)
        else:
            raise TypeError(f"Unsupported key type {args[0]}")
    elif origin is typing.Union:
        if len(args) == 2 and args[1] is type(None):
            return optional_value_convert_c_to_py(args[0], c_value.contents)
        else:
            raise TypeError(f"Unsupported type {py_type}")
    else:
        raise TypeError(f"Unsupported type {py_type}")


def struct_value_convert_c_to_py(model_type: Type[BaseModel], c_struct: Structure) -> BaseModel:
    return model_type(
        **{
            name: value_convert_c_to_py(field.annotation, getattr(c_struct, name))
            for name, field in model_type.__fields__.items()
        }
    )


def list_value_convert_c_to_py(py_type: PyType, c_list: Structure) -> List[Any]:
    c_type = type_convert_py_to_c(py_type)
    values_pointer = cast(c_list.values, POINTER(c_type * c_list.len))
    return [
        value_convert_c_to_py(py_type, values_pointer.contents[i])
        for i in range(int(c_list.len))
    ]


def dict_value_convert_c_to_py(py_type: PyType, c_dict: Structure) -> Dict[str, Any]:
    c_type = type_convert_py_to_c(py_type)
    keys_pointer = cast(c_dict.keys, POINTER(c_char_p * c_dict.len))
    values_pointer = cast(c_dict.values, POINTER(c_type * c_dict.len))
    return {
        keys_pointer.contents[i].decode("utf-8"):
            value_convert_c_to_py(py_type, values_pointer.contents[i])
        for i in range(int(c_dict.len))
    }


def optional_value_convert_c_to_py(c_type: CType, c_struct: Structure) -> Optional[Any]:
    return value_convert_c_to_py(c_type, c_struct.value) if c_struct.is_some else None


def value_convert_py_to_c(py_type: PyType, py_value: Any) -> Any:
    if inspect.isclass(py_type):
        if py_type is bool:
            return c_bool(py_value)
        elif py_type is int:
            return c_longlong(py_value)
        elif py_type is float:
            return c_double(py_value)
        elif py_type is str:
            return c_char_p(py_value.encode("utf-8"))
        elif issubclass(py_type, BaseModel):
            return byref(struct_value_convert_py_to_c(py_type, py_value))
        else:
            raise TypeError(f"Unsupported python class type {py_type}")

    origin = get_origin(py_type)
    args = get_args(py_type)
    if origin is list:
        return byref(list_value_convert_py_to_c(args[0], py_value))
    elif origin is dict:
        if args[0] is str:
            return byref(dict_value_convert_py_to_c(args[1], py_value))
        else:
            raise TypeError(f"Unsupported key type {args[0]}")
    elif origin is typing.Union:
        if len(args) == 2 and args[1] is type(None):
            return byref(optional_value_convert_py_to_c(args[0], py_value))
        else:
            raise TypeError(f"Unsupported type {py_type}")
    else:
        raise TypeError(f"Unsupported type {py_type}")


def arg_values_convert_py_to_c(model_type: Type[BaseModel], model_value: BaseModel) -> List:
    return [
        value_convert_py_to_c(field.annotation, getattr(model_value, name))
        for name, field in model_type.__fields__.items()
    ]


def struct_value_convert_py_to_c(model_type: Type[BaseModel], model_value: BaseModel) -> Structure:
    c_struct = struct_type_convert_py_to_c(model_type)
    return c_struct(
        **{
            name: value_convert_py_to_c(field.annotation, getattr(model_value, name))
            for name, field in model_type.__fields__.items()
        }
    )


def list_value_convert_py_to_c(py_type: PyType, py_list: List[Any]) -> Structure:
    c_type = type_convert_py_to_c(py_type)
    c_list = wrap_list_type_to_c_struct(c_type)
    _len = len(py_list)
    return c_list(
        len=_len,
        values=(c_type * _len)(*[value_convert_py_to_c(py_type, value) for value in py_list])
    )


def dict_value_convert_py_to_c(py_type: PyType, py_dict: Dict[str, Any]) -> Structure:
    c_type = type_convert_py_to_c(py_type)
    c_dict = wrap_dict_type_to_c_struct(c_type)
    _len = len(py_dict)
    return c_dict(
        len=_len,
        keys=(c_char_p * _len)(*[c_char_p(key.encode("utf-8")) for key in py_dict.keys()]),
        values=(c_type * _len)(*[value_convert_py_to_c(py_type, value) for value in py_dict.values()])
    )


def optional_value_convert_py_to_c(py_type: PyType, py_value: Optional[Any]) -> Structure:
    c_type = type_convert_py_to_c(py_type)
    c_struct = wrap_optional_type_to_c_struct(c_type)
    if py_value is None:
        return c_struct(is_some=False)
    else:
        return c_struct(is_some=True, value=value_convert_py_to_c(py_type, py_value))
