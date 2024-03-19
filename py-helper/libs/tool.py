import json

from ctypes import cdll
from typing import Optional, Type, Dict, Any, Tuple
from pydantic import BaseModel, Field, create_model
from datamodel_code_generator.model import DataModelFieldBase
from datamodel_code_generator.parser.jsonschema import JsonSchemaParser

from .converter import (arg_types_convert_py_to_c, struct_type_convert_py_to_c, arg_values_convert_py_to_c,
                        struct_value_convert_c_to_py)


class ToolSchema(BaseModel):
    """The schema for a tool. It defines the input arguments and the return type of the tool."""

    id: str
    """The unique identifier of this tool."""
    name: str
    """The unique name of the tool that clearly communicates its purpose."""
    description: str
    """Used to tell the model how/when/why to use the tool."""
    args_schema: Optional[Dict[str, Any]] = None
    """Pydantic model class to validate and parse the tool's input arguments."""
    return_schema: Optional[Dict[str, Any]] = None
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
        if self.args_schema is None:
            return None
        return self.create_mode_type_from_schema(self.args_schema)

    @property
    def return_schema_model(self) -> Optional[Type[BaseModel]]:
        """Returns the Pydantic model class to validate and parse the tool's output."""
        if self.return_schema is None:
            return None
        return self.create_mode_type_from_schema(self.return_schema)

    @staticmethod
    def create_mode_type_from_schema(_schema: Dict[str, Any]) -> Optional[Type[BaseModel]]:
        parser = JsonSchemaParser(
            json.dumps(_schema),
            validation=True,
            use_schema_description=True,
            use_field_description=True,
        )
        parser.parse()

        def _build_field(field: DataModelFieldBase) -> Tuple:
            if field.has_default:
                return field.data_type.type_hint, Field(field.default, alias=field.alias, description=field.docstring)
            else:
                return field.data_type.type_hint, Field(alias=field.alias, description=field.docstring)

        if len(parser.results) > 0:
            result = parser.results[0]
            return create_model(
                result.name,
                **{field.name: _build_field(field) for field in result.fields}
            )
        else:
            return None

    def run_tool(self, path: str, args: Optional[BaseModel]) -> Optional[BaseModel]:
        try:
            lib = cdll.LoadLibrary(path)
        except Exception as e:
            raise ValueError(f"Failed to load tool at {path}: {e}")

        args_schema = self.args_schema_model
        return_schema = self.return_schema_model

        func = getattr(lib, self.name)
        func.argtypes = arg_types_convert_py_to_c(args_schema) if args_schema else []
        func.restype = struct_type_convert_py_to_c(return_schema) if return_schema else None

        c_args = arg_values_convert_py_to_c(args_schema, args) if args_schema else []
        c_resp = func(*c_args)
        return struct_value_convert_c_to_py(return_schema, c_resp) if return_schema else None
