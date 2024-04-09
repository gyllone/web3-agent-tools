import platform

from typing import Any, Mapping, Optional
from pydantic import BaseModel

from libs.schema import Schema
from libs.converter import ValueConverter


class ToolSchema(BaseModel):
    """The schema for a tool. It defines the input arguments and the return type of the tool."""

    name: str
    """The unique name of the tool that clearly communicates its purpose."""
    description: str
    """Used to tell the model how/when/why to use the tool."""
    args_schema: Optional[Schema] = None
    """Pydantic model class to validate and parse the tool's input arguments."""
    result_schema: Optional[Schema] = None
    """Pydantic model class to validate and parse the tool's output."""
    return_direct: bool = False
    """Whether to return the tool's output directly. Setting this to True means that after the tool is called, the 
    AgentExecutor will stop looping."""

    metadata: Optional[dict[str, Any]] = None
    """Optional metadata associated with the tool. Defaults to None
    This metadata will be associated with each call to this tool,
    and passed as arguments to the handlers defined in `callbacks`.
    You can use these to eg identify a specific instance of a tool with its use case."""

    def run_tool(self, path: str, args: Optional[Mapping[str, Any]]) -> Optional[Mapping[str, Any]]:
        try:
            os_type = platform.system()
            if os_type == "Windows":
                from ctypes import windll
                lib = windll.LoadLibrary(path)
            else:
                from ctypes import cdll
                lib = cdll.LoadLibrary(path)
        except BaseException as e:
            raise ValueError(f"Failed to load tool at {path}: {e}")

        args_converter = ValueConverter(self.args_schema) if self.args_schema else None
        result_converter = ValueConverter(self.result_schema) if self.result_schema else None

        c_func = getattr(lib, self.name)
        c_func.argtypes = args_converter.get_arg_types() if args_converter else []
        c_func.restype = result_converter.get_structure_type() if result_converter else None

        c_release_func = getattr(lib, f"{self.name}_release")
        if c_release_func and c_func.restype:
            c_release_func.argtypes = [c_func.restype]
            c_release_func.restype = None

        c_args = args_converter.py_object_to_c_values(args) if args_converter else []
        c_res = c_func(*c_args)

        if result_converter:
            py_res = result_converter.c_struct_to_py_object(c_res)
            if c_release_func and c_res:
                c_release_func(c_res)
            return py_res
        else:
            return None
