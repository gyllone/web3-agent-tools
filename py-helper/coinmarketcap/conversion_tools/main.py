import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


class PriceInput(BaseModel):
    amount: float
    id: Optional[str]
    symbol: Optional[str]
    time: Optional[str]
    convert: Optional[str]
    convert_id: Optional[str]


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Quote(BaseModel):
    price: float
    last_updated: str


class Price(BaseModel):
    id: int
    symbol: str
    name: str
    amount: float
    last_updated: str
    quote: dict[str, Quote]


class PriceOutput(Result):
    value: Optional[Price] = Field(description="111")


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_price_conversion",
        description="This is a test",
        args_schema=ParamSchema.from_model_type(PriceInput),
        result_schema=ParamSchema.from_model_type(PriceOutput),
    )

    print("\n===============Running Tool===============\n")

    args = PriceInput(amount=10.27, id="1", convert="USD")

    resp = schema.run_tool("../../../go-tools/output/conversion_tools.so",
                           **args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
