import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


class QuoteInput(BaseModel):
    convert: Optional[str]
    convert_id: Optional[str]


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Quote(BaseModel):
    total_market_cap: float
    total_volume_24h: float
    total_volume_24h_reported: float
    altcoin_volume_24h: float
    altcoin_volume_24h_reported: float
    altcoin_market_cap: float
    defi_volume_24h: float
    defi_volume_24h_reported: float
    defi_24h_percentage_change: float
    defi_market_cap: float
    stablecoin_volume_24h: float
    stablecoin_volume_24h_reported: float
    stablecoin_24h_percentage_change: float
    stablecoin_market_cap: float
    derivatives_volume_24h: float
    derivatives_volume_24h_reported: float
    derivatives_24h_percentage_change: float
    total_market_cap_yesterday: float
    total_volume_24h_yesterday: float
    total_market_cap_yesterday_percentage_change: float
    total_volume_24h_yesterday_percentage_change: float
    last_updated: str


class Metric(BaseModel):
    active_cryptocurrencies: int
    total_cryptocurrencies: int
    active_market_pairs: int
    active_exchanges: int
    total_exchanges: int
    eth_dominance: float
    btc_dominance: float
    eth_dominance_yesterday: float
    btc_dominance_yesterday: float
    eth_dominance_24h_percentage_change: float
    btc_dominance_24h_percentage_change: float
    defi_volume_24h: float
    defi_volume_24h_reported: float
    defi_market_cap: float
    defi_24h_percentage_change: float
    stablecoin_volume_24h: float
    stablecoin_volume_24h_reported: float
    stablecoin_market_cap: float
    stablecoin_24h_percentage_change: float
    derivatives_volume_24h: float
    derivatives_volume_24h_reported: float
    derivatives_24h_percentage_change: float
    quote: dict[str, Quote]
    last_updated: str


class QuoteOutput(Result):
    value: Optional[Metric] = Field(description="111")


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_quotes_latest",
        description="This is a test",
        args_schema=ParamSchema.from_model_type(QuoteInput),
        result_schema=ParamSchema.from_model_type(QuoteOutput),
    )

    print("\n===============Running Tool===============\n")

    args = QuoteInput(convert="USD")

    resp = schema.run_tool("../../../go-tools/output/global_metrics.so", args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
