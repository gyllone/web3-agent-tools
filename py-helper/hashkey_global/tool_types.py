

from pydantic import BaseModel, Field
from typing import Optional

class Balance(BaseModel):
    Asset: str=Field(description="virtual asset name, like BTC, ETH")
    Total: str=Field(description="virtual asset total balance")
    Free: str=Field(description="virtual asset available amount")


class GetTradeAccountBalanceArgs(BaseModel):
    ApiKey: str=Field(description="api key of hashkey global exchange")
    ApiSignKey: str=Field(description="api sign key of hashkey global exchange")

class GetTradeAccountBalanceResult(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    result: Optional[list[Balance]] = Field(None, description="list of user's virtual asset balance in spot trade account of hashkey global exchange")

class SpotOrder(BaseModel):
    orderId: str = Field(description="orderid")
    SymbolName: str = Field(description="orderid")
    TransactTime: str = Field(description="orderid")
    Price: str = Field(description="orderid")
    Status: str = Field(description="orderid")
    OrigQty: str = Field(description="orderid")
    ExecutedQty: str = Field(description="orderid")

class CreateSpotLimitOrderArgs(BaseModel):
    ApiKey: str=Field(description="api key of hashkey global exchange")
    Secret: str=Field(description="api sign key of hashkey global exchange")
    Symbol:str=Field(description="trading pair of crypto, like BTCUSDT means BTC as base currency and USDT as quote currency")
    Side: str=Field(description="'BUY' or 'SELL' in plain text")
    Price: str=Field(description="specific value at which a trader is willing to buy or sell a particular cryptocurrency pair.")
    Quantity: str=Field(description="quantity of base asset of symbol, e.g. quantity of BTCUSDT means quantity BTC to trade")
    

class CreateSpotMarketOrderArgs(BaseModel):
    ApiKey: str=Field(description="api key of hashkey global exchange")
    Secret: str=Field(description="api sign key of hashkey global exchange")
    Symbol: str=Field(description="trading pair of crypto, like BTCUSDT means BTC as base currency and USDT as quote currency")
    Side: str=Field(description="'BUY' or 'SELL' in plain text")
    Quantity: str=Field(description="quantity of base asset of symbol, e.g. quantity of BTCUSDT means quantity BTC to trade")

class CreateSpotMarketOrderResult(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    result: Optional[list[SpotOrder]] = Field(None, description="spot order")

class Kline(BaseModel):
    t: str=Field(description="open timestamp of Timestamp in RFC3339 format")
    s: str=Field(description="symbol")
    o: str = Field(description="opening price")
    c: str = Field(description="closing price")
    h: str = Field(description="highest price")
    l: str = Field(description="lowest price")
    v: str = Field(description="traded volume")
    

class GetKlineArgs(BaseModel):
    Symbol: str=Field(description="trading pair of crypto, like BTCUSDT means BTC as base currency and USDT as quote currency")
    Interval: str=Field(description='''time interval of candlestick chart interval.
                         m for minutes; h for hours; d for days; w for weeks; M for months; 
                        available values are [3m,5m,15m,30m,1h,2h,4h,6h,8h,12h,1d,1w,1M]''') 
    StartTime: Optional[str]=Field(description="start time of kline chart in RFC3339 format, format like 2023-04-05T17:45:30+08:00")
    EndTime: Optional[str]=Field(description="end time of kline chart in RFC3339 format, e.g. 2023-04-05T17:45:30+08:00")
    Limit: Optional[str]=Field(description="Return the number of bars, the maximum value and defaut value is 1000")

class GetKlineResult(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    result: Optional[list[Kline]] = Field(None, description="list of kline bars")

class Price(BaseModel):
    symbol:str = Field(description="trading pair of crypto, like BTCUSDT")
    price: str = Field(description="price")

class GetLatestPriceArgs(BaseModel):
    Symbol: str=Field(description="trading pair of crypto, like BTCUSDT means BTC as base currency and USDT as quote currency")

class GetLatestPriceResult(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    result: Optional[list[Price]]=Field(None, description="latest price")
    