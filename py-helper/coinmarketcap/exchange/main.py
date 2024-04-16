import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


class IdMapInput(BaseModel):
    start: Optional[int]
    limit: Optional[int]
    listing_status: Optional[str]
    slug: Optional[str]
    sort: Optional[str]
    aux: Optional[str]
    crypto_id: Optional[str]


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Exchange(BaseModel):
    id: int
    name: str
    slug: str
    is_active: int
    is_listed: int
    is_redistributable: int
    first_historical_data: str
    last_historical_data: str


class IdMapOutput(Result):
    value: list[Exchange]


class MetadataInput(BaseModel):
    id: Optional[str]
    slug: Optional[str]
    aux: Optional[str]


class Metadata(BaseModel):
    id: int
    name: str
    slug: str
    description: str
    notice: str
    fiats: list[str]
    urls: str
    date_launched: str
    maker_fee: float
    taker_fee: float
    spot_volume_usd: float
    spot_volume_last_updated: str
    weekly_visits: int


class MetadataOutput(Result):
    value: dict[str, Metadata]


class AssetInput(BaseModel):
    id: str


class Platform(BaseModel):
    crypto_id: int
    symbol: str
    name: str


class Currency(BaseModel):
    crypto_id: int
    price_usd: float
    symbol: str
    name: str


class Asset(BaseModel):
    wallet_address: str
    balance: float
    platform: Platform
    currency: Currency


class AssetOutput(Result):
    value: list[Asset]


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_exchange_assets",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(IdMapInput),
        # result_schema=ParamSchema.from_model_type(IdMapOutput),
        # args_schema=ParamSchema.from_model_type(MetadataInput),
        # result_schema=ParamSchema.from_model_type(MetadataOutput),
        args_schema=ParamSchema.from_model_type(AssetInput),
        result_schema=ParamSchema.from_model_type(AssetOutput),
    )

    print("\n===============Running Tool===============\n")

    # args = IdMapInput(start=10, limit=20, slug="binance", sort="volume_24h")
    # args = MetadataInput(slug="binance,gdax")
    args = AssetInput(id="270")

    resp = schema.run_tool("../../../go-tools/output/exchange.so", args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
