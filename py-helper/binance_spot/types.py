import json

from typing import Optional
from pydantic import BaseModel, Field

from libs.tool import ToolSchema
from libs.schema import ParamSchema


class KLine(BaseModel):
    open_time: float
    open_price: float
    high_price: float
    low_price: float
    close_price: float
    volume: float
    close_time: str
    quote_asset_volume: float
    number_of_trades: int
    taker_buy_base_asset_volume: float
    taker_buy_quote_asset_volume: float


class PriceChange24h(BaseModel):
    symbol: str
    price_change: float
    price_change_pct: float
    weighted_avg_price: float
    last_price: float
    open_price: float
    high_price: float
    low_price: float
    volume: float
    quote_volume: float
    open_time: float
    close_time: float
    count: int


class Input(BaseModel):
    foo: Optional[str] = Field(None, description="foo")
    bar: list[dict[str, int]] = Field(description="bar")
    baz: Param = Field(description="baz")
