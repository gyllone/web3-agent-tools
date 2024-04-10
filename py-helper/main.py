import json
from typing import Optional, List, Dict
from pydantic import BaseModel, Field
from libs.schema import Schema
from libs.tool import ToolSchema


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

#     args = Input(
#         foo="foo",
#         bar=[
#             {"x": 1, "y": 2},
#             {"x": 3, "y": 4},
#         ],
#         baz=Param(
#             foo=True,
#             bar=[1.0, 2.0],
#             baz={"x": 1, "y": 2},
#         ),
#     )
#     resp = schema.run_tool("../go-tools/outputs/test.so", args.dict(by_alias=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
