import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.schema import Schema
from libs.tool import ToolSchema


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


class IdMapInput(BaseModel):
    listing_status: Optional[str] = Field(description="listing_status")
    start: Optional[int] = Field(description="start")
    limit: Optional[int] = Field(description="limit")
    sort: Optional[str] = Field(description="sort")
    symbol: Optional[str] = Field(description="symbol")
    aux: Optional[str] = Field(description="aux")


class Result(BaseModel):
    is_fail: bool = Field(description="is_fail")
    error_message: str = Field(description="error_message")


class IdMapOutput(Result):
    id_maps: list[dict[str, str]] = Field(description="idMaps")


class QuoteInput(BaseModel):
    ids: Optional[str]
    slug: Optional[str]
    convert: Optional[str]
    convert_id: Optional[str]
    aux: Optional[str]
    skip_invalid: Optional[bool]


class QuoteOutput(Result):
    quotes: list[list[float]] = Field(description="quotes")


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_quotes",
        description="This is a test",
        args_schema=Schema.from_model_type(QuoteInput),
        result_schema=Schema.from_model_type(QuoteOutput),
    )

    print("\n===============Running Tool===============\n")

    args = QuoteInput(ids="1,2", convert="USD", skip_invalid=False)
    resp = schema.run_tool("../go-tools/output/idmaps.so", args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
