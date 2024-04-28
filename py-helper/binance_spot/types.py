from typing import Optional
from pydantic import BaseModel, Field


class QueryWithdrawHistory(BaseModel):
    api_key: str
    secret_key: str
    coin: Optional[str] = Field(None, description="coin name in uppercase, if not provided, query all coins")
    withdraw_id: Optional[str] = Field(None, description="withdraw id")
    status: Optional[int] = Field(None, description="withdraw status: 0=pending, 1=success, 2=failed")
    offset: Optional[int] = Field(None, description="offset")
    limit: Optional[int] = Field(None, description="limit")
    start_time: Optional[int] = Field(None, description="start time in milliseconds")
    end_time: Optional[int] = Field(None, description="end time in milliseconds")


class WithdrawHistory(BaseModel):
    id: str
    amount: float
    tx_fee: float
    coin: str
    status: int
    address: str
    tx_id: str
    apply_time: str
    network: str
    transfer_type: int
    info: str
    confirmations: int
    wallet_type: int


class WithdrawHistoryOutput(BaseModel):
    status: bool
    error: str
    result: Optional[list[WithdrawHistory]]


class QueryFundingAsset(BaseModel):
    api_key: str
    secret_key: str
    asset: Optional[str] = Field(None, description="coin name in uppercase, if not provided, query all coins")


class FundingAsset(BaseModel):
    asset: str
    free: float
    locked: float
    freeze: float
    withdrawing: float
    btc_valuation: float


class FundingAssetsOutput(BaseModel):
    status: bool
    error: str
    result: Optional[list[FundingAsset]]


class QueryAccountInfo(BaseModel):
    api_key: str
    secret_key: str
    recv_window: Optional[int] = Field(None, description="specify the request must be processed within a certain "
                                                         "number of milliseconds or be rejected by the server, The "
                                                         "value cannot be greater than 60000")


class Balance(BaseModel):
    asset: str
    free: float
    locked: float


class AccountInfo(BaseModel):
    can_trade: bool
    can_withdraw: bool
    can_deposit: bool


class QueryKlines(BaseModel):
    symbol: str = Field(description="trading symbol, combined with base and quote token in uppercase, e.g. BTCUSDT")
    interval: str = Field("1d", description="time interval, supported values: 1m,3m,5m,15m,30m,1h,2h,4h,6h,8h.12h,"
                                            "1d,3d,1w,1M")
    limit: int = Field(500, description="number of klines, max 1000")
    start_time: Optional[str] = Field(None, description="start time in RFC3339 format, e.g. 2023-01-01T00:00:00Z")
    end_time: Optional[str] = Field(None, description="end time in RFC3339 format, e.g. 2023-01-01T00:00:00Z")


class KLine(BaseModel):
    t: str = Field(description="open time")
    s: str = Field(description="symbol")
    o: float = Field(description="open price")
    c: float = Field(description="close price")
    h: float = Field(description="high price")
    l: float = Field(description="low price")
    v: float = Field(description="volume")
    q: float = Field(description="quote asset volume")
    tb: float = Field(description="taker buy base asset volume")
    tq: float = Field(description="taker buy quote asset volume")
    n: int = Field(description="number of trades")


class KlineOutput(BaseModel):
    status: bool
    error: str
    result: Optional[list[KLine]]


class QueryPriceChange24h(BaseModel):
    bases: list[str] = Field([], description='list of base tokens in uppercase, max number is 100. E.g. ["BTC", '
                                             '"ETH"]. If not provided, query all tokens')
    quote: str = Field("USDT", description="quote token in uppercase, e.g. USDT")
    descending: bool = Field(True, description="sort by price change percentage in descending order")
    limit: int = Field(10, description="number of price changes, max 100")


class PriceChange(BaseModel):
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
    open_time: str
    close_time: str
    count: int


class PriceChangeOutput(BaseModel):
    status: bool
    error: str
    result: Optional[list[PriceChange]]


class QueryRollingPriceChange(BaseModel):
    symbols: list[str] = Field(description='list of trading symbols in uppercase, max number is 100. E.g. ["BTCUSDT",'
                                           '"ETHUSDT"]')
    descending: bool = Field(True, description="sort by price change percentage in descending order")
    window_size: str = Field("1d", description="window size, supported values: 1m,2m....59m for minutes; 1h,2h....23h "
                                               "for hours, 1d...7d for days. Note: units cannot be combined (e.g. "
                                               "1d2h is not allowed).")


class LatestPrice(BaseModel):
    symbol: str
    price: float


class CreateOrderRequest(BaseModel):
    api_key: str
    secret_key: str
    symbol: str = Field(description="trading symbol, combined with base and quote token in uppercase, e.g. BTCUSDT")
    side: str = Field(description="trade side, supported values: BUY, SELL")
    order_type: str = Field(description="order type, supported values: LIMIT, MARKET, STOP_LOSS, STOP_LOSS_LIMIT, "
                                        "TAKE_PROFIT, TAKE_PROFIT_LIMIT, LIMIT_MAKER")
    quantity: float
    time_in_force: Optional[str] = Field(None, description="How long an order will be active before expiration, "
                                                           "which is mandatory used with LIMIT, STOP_LOSS_LIMIT and "
                                                           "TAKE_PROFIT_LIMIT orders. Supported values: GTC, IOC, FOK")
    price: Optional[float] = Field(None, description="used with LIMIT, STOP_LOSS_LIMIT, TAKE_PROFIT_LIMIT and "
                                                     "LIMIT_MAKER orders")
    stop_price: Optional[float] = Field(None, description="used with STOP_LOSS, STOP_LOSS_LIMIT, TAKE_PROFIT, "
                                                          "and TAKE_PROFIT_LIMIT orders")
    recv_window: Optional[int] = Field(None, description="specify the request must be processed within a certain "
                                                         "number of milliseconds or be rejected by the server, The "
                                                         "value cannot be greater than 60000")


class OrderResponse(BaseModel):
    symbol: str
    order_id: int = Field(description="order id")
    transact_time: str = Field(description="order processed time")
    working_time: str = Field(description="order matched time")
    price: float = Field(description="order price")
    orig_qty: float = Field(description="order original quantity")
    executed_qty: float = Field(description="order executed quantity")
    cumulative_quote_qty: float = Field(description="order cumulative quote quantity")
    status: str
    time_in_force: str = Field(description="how long the order will be active before expiration")
    order_type: str
    side: str


class OrderResponseOutput(BaseModel):
    status: bool
    error: str
    result: Optional[OrderResponse]


class CancelOrderRequest(BaseModel):
    api_key: str
    secret_key: str
    symbol: str = Field(description="trading symbol, combined with base and quote token in uppercase, e.g. BTCUSDT")
    order_id: int = Field(description="order id")
    cancel_restrictions: Optional[str] = Field(None, description="cancel restrictions, supported values: ONLY_NEW, "
                                                                 "ONLY_PARTIALLY_FILLED")
    recv_window: Optional[int] = Field(None, description="specify the request must be processed within a certain "
                                                         "number of milliseconds or be rejected by the server, The "
                                                         "value cannot be greater than 60000")


class CancelOrderOutput(BaseModel):
    status: bool
    error: str
    result: Optional[int]


class GetOpenOrdersRequest(BaseModel):
    api_key: str
    secret_key: str
    symbol: Optional[str] = Field(description="symbol of the order, if not provided, query all open orders")
    recv_window: Optional[int] = Field(None, description="specify the request must be processed within a certain "
                                                         "number of milliseconds or be rejected by the server, The "
                                                         "value cannot be greater than 60000")


class Order(BaseModel):
    symbol: str
    order_id: int
    price: float
    orig_qty: float
    executed_qty: float
    cumulative_quote_qty: float
    status: str
    time_in_force: str
    order_type: str
    side: str
    stop_price: float
    timestamp: str
    update_time: str
    is_working: bool
    working_time: str
    orig_quote_order_qty: float


class GetOpenOrdersOutput(BaseModel):
    status: bool
    error: str
    result: Optional[list[Order]]
