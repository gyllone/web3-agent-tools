import json

from ctypes import cdll
from typing import Optional, Type, Dict, Any, Tuple
from pydantic import BaseModel, create_model
from pydantic.fields import FieldInfo, Undefined
from datamodel_code_generator.model import DataModelFieldBase
from datamodel_code_generator.parser.jsonschema import JsonSchemaParser

from .converter import (arg_types_convert_py_to_c, type_convert_py_to_c, arg_values_convert_py_to_c,
                        value_convert_c_to_py)


class ToolSchema(BaseModel):
    """The schema for a tool. It defines the input arguments and the return type of the tool."""

    name: str
    """The unique name of the tool that clearly communicates its purpose."""
    description: str
    """Used to tell the model how/when/why to use the tool."""
    args_schema: Optional[Dict[str, Any]] = None
    """Pydantic model class to validate and parse the tool's input arguments."""
    result_schema: Optional[Dict[str, Any]] = None
    """Pydantic model class to validate and parse the tool's output."""
    return_direct: bool = False
    """Whether to return the tool's output directly. Setting this to True means that after the tool is called, the 
    AgentExecutor will stop looping."""

    metadata: Optional[Dict[str, Any]] = None
    """Optional metadata associated with the tool. Defaults to None
    This metadata will be associated with each call to this tool,
    and passed as arguments to the handlers defined in `callbacks`.
    You can use these to eg identify a specific instance of a tool with its use case."""

    @property
    def args_schema_model(self) -> Optional[Type[BaseModel]]:
        """Returns the Pydantic model class to validate and parse the tool's input arguments."""
        return self.create_model_type_from_schema(self.args_schema) if self.args_schema else None

    @property
    def result_schema_model(self) -> Optional[Type[BaseModel]]:
        """Returns the Pydantic model class to validate and parse the tool's output."""
        return self.create_model_type_from_schema(self.result_schema) if self.result_schema else None

    @staticmethod
    def create_model_type_from_schema(_schema: Dict[str, Any]) -> Optional[Type[BaseModel]]:
        parser = JsonSchemaParser(
            json.dumps(_schema),
            validation=True,
            use_schema_description=True,
            use_field_description=True,
            reuse_model=True,
            collapse_root_models=True,
        )
        parser.parse()

        def _build_field(field: DataModelFieldBase) -> Tuple[str, FieldInfo]:
            return (
                field.type_hint,
                FieldInfo(
                    default=field.default if field.has_default else Undefined,
                    alias=field.alias,
                    description=field.docstring,
                )
            )

        if len(parser.results) == 0:
            return None
        elif len(parser.results) == 1:
            result = parser.results[0]
            return create_model(
                result.name,
                **{field.name: _build_field(field) for field in result.fields}
            )
        else:
            raise ValueError("too many models, only support 1")

    def run_tool(self, path: str, args: Optional[BaseModel]) -> Optional[BaseModel]:
        try:
            lib = cdll.LoadLibrary(path)
        except BaseException as e:
            raise ValueError(f"Failed to load tool at {path}: {e}")

        args_schema = self.args_schema_model
        result_schema = self.result_schema_model

        c_func = getattr(lib, self.name)
        c_func.argtypes = arg_types_convert_py_to_c(args_schema) if args_schema else []
        c_func.restype = type_convert_py_to_c(result_schema) if result_schema else None

        c_release_func = getattr(lib, f"{self.name}_release")
        c_release_func.argtypes = [c_func.restype]
        c_release_func.restype = None

        c_args = arg_values_convert_py_to_c(args_schema, args) if args_schema else []
        c_res = c_func(*c_args)

        if result_schema:
            if not c_res:
                raise ValueError(f"Tool {self.name} returned nil pointer")
            else:
                py_res = value_convert_c_to_py(result_schema, c_res)
                c_release_func(c_res)
                return py_res
        else:
            return None
