import json
from typing import Optional
from pydantic import BaseModel, Field

from libs.tool import ToolSchema
from libs.schema import Schema


class Param(BaseModel):
    foo: bool = Field(description="foo")
    bar: list[float] = Field(description="bar")
    baz: dict[str, int] = Field(description="baz")


class Input(BaseModel):
    foo: Optional[str] = Field(None, description="foo")
    bar: list[dict[str, int]] = Field(description="bar")
    baz: Param = Field(description="baz")


class Output(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    result: Optional[Param] = Field(None, description="result")


if __name__ == '__main__':
    schema = ToolSchema(
        name="test",
        description="This is a test",
        args_schema=Schema.from_model_type(Input),
        result_schema=Schema.from_model_type(Output),
    )

    print("\n===============Running Tool===============\n")

    args = Input(
        foo="foo",
        bar=[
            {"x": 1, "y": 2},
            {"x": 3, "y": 4},
        ],
        baz=Param(
            foo=True,
            bar=[1.0, 2.0],
            baz={"x": 1, "y": 2},
        ),
    )
    resp = schema.run_tool("../go-tools/outputs/test.so", args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
