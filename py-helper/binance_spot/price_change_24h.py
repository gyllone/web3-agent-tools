import json

from libs.tool import ToolSchema
from libs.schema import ParamSchema

from binance_spot.types import QueryPriceChange24h, PriceChangeOutput


if __name__ == '__main__':
    schema = ToolSchema(
        name="price_change_24h_statistics",
        description="Query 24 hour rolling window price change statistics.",
        args_schema=ParamSchema.from_model_type(QueryPriceChange24h),
        result_schema=ParamSchema.from_model_type(PriceChangeOutput),
    )
    print(schema.json(by_alias=True, exclude_none=True))

    print("\n===============Running Tool===============\n")

    args = QueryPriceChange24h().dict(by_alias=True)
    resp = schema.run_tool("../../go-tools/outputs/binance_spot.so", **args)
    if resp is not None:
        print(json.dumps(resp, indent=2))
