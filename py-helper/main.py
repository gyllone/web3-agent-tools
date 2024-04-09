import json
from typing import Optional
from pydantic import BaseModel, Field

from libs.tool import ToolSchema
from libs.schema import Schema


class V(BaseModel):
    a: str
    b: int


class A(BaseModel):
    status: bool = Field(False, description="status")
    value: float = Field(0, description="value")
    name: str = Field("", description="name")


class B(BaseModel):
    a: Optional[int] = Field(None, description="test")
    b: "B"
    c: dict[str, list[A]] = Field(description="test")


class InputInternalA(BaseModel):
    status: int = Field(description="status")
    name: str = Field(description="name")


class InputInternal(BaseModel):
    status: bool = Field(description="status")
    value: float = Field(description="value")
    name: str = Field(description="name")
    a: InputInternalA = Field(description="a")


class Output(BaseModel):
    status: bool = Field(description="status")
    value: float = Field(description="value")
    name: str = Field(description="name")


class Input(BaseModel):
    input: InputInternal = Field(description="input")


if __name__ == '__main__':
    schema = ToolSchema(
        name="test",
        description="This is a test",
        args_schema=Schema.from_model_type(Input),
        result_schema=Schema.from_model_type(Output),
    )

    print("\n===============Running Tool===============\n")

    args = Input(
        input=InputInternal(
            status=True,
            value=1.0,
            name="test",
            a=InputInternalA(
                status=200,
                name="test"
            ),
        )
    )
    resp = schema.run_tool("../go-tools/outputs/test.so", args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
