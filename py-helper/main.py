import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


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
    sort: Optional[str] = Field(description="sort")
    symbol: Optional[str] = Field(description="symbol")
    aux: Optional[str] = Field(description="aux")
    start: Optional[int] = Field(description="start")
    limit: Optional[int] = Field(description="limit")


class Result(BaseModel):
    is_fail: bool = Field(description="is_fail")
    error_message: str = Field(description="error_message")


class IdMapOutput(Result):
    data: list[dict[str, str]] = Field(description="idMaps")


class QuoteInput(BaseModel):
    id: Optional[str]
    slug: Optional[str]
    convert: Optional[str]
    convert_id: Optional[str]
    aux: Optional[str]
    skip_invalid: Optional[bool]


class QuoteOutput(Result):
    data: list[list[float]] = Field(description="quotes")


class MetaInput(BaseModel):
    id: Optional[str]
    slug: Optional[str]
    address: Optional[str]
    aux: Optional[str]
    skip_invalid: Optional[bool]


class MetaOutput(Result):
    data: list[dict[str, str]] = Field(description="metas")


class ListingInput(BaseModel):
    start: Optional[int]
    limit: Optional[int]
    price_min: Optional[int]
    price_max: Optional[int]
    market_cap_min: Optional[int]
    market_cap_max: Optional[int]
    volume_24h_min: Optional[int]
    volume_24h_max: Optional[int]
    circulating_supply_min: Optional[int]
    circulating_supply_max: Optional[int]
    percent_change_24h_min: Optional[int]
    percent_change_24h_max: Optional[int]
    convert: Optional[str]
    convert_id: Optional[str]
    sort: Optional[str]
    sort_dir: Optional[str]
    cryptocurrency_type: Optional[str]
    tag: Optional[str]
    aux: Optional[str]


class Market(BaseModel):
    metadata: dict[str, str]
    quotes: dict[str, dict[str, str]]


class ListingOutput(Result):
    data: list[Market] = Field(description="market data")


if __name__ == '__main__':
    schema = ToolSchema(
        # name="query_quotes",
        name="query_listings",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(QuoteInput),
        # result_schema=ParamSchema.from_model_type(QuoteOutput),
        # args_schema=ParamSchema.from_model_type(ListingInput),
        # result_schema=ParamSchema.from_model_type(ListingOutput),
        args_schema=ParamSchema.from_model_type(Input),
        result_schema=ParamSchema.from_model_type(Output),
    )

    print("\n===============Running Tool===============\n")

    # args = QuoteInput(id="1,3,5")
    args = ListingInput(limit=5, convert="ETH")
    resp = schema.run_tool("../go-tools/output/crycur.so", args.dict(by_alias=True))

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
