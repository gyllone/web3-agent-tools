import json

from libs.tool import ToolSchema
from libs.schema import ParamSchema

from binance_spot.types import QueryRollingPriceChange, PriceChangeOutput


if __name__ == '__main__':
    schema = ToolSchema(
        name="rolling_window_price_change_statistics",
        description="Query price change statistics in a rolling window for specified symbols",
        args_schema=ParamSchema.from_model_type(QueryRollingPriceChange),
        result_schema=ParamSchema.from_model_type(PriceChangeOutput),
    )
    # print(schema.json(by_alias=True, exclude_none=True))

    print("\n===============Running Tool===============\n")

    args = QueryRollingPriceChange(symbols=["BTCUSDT"]).dict(by_alias=True)
    resp = schema.run_tool("../../go-tools/outputs/binance_spot.so", **args)
    if resp is not None:
        print(json.dumps(resp, indent=2))
