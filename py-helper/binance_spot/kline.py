import json

from libs.tool import ToolSchema
from libs.schema import ParamSchema

from binance_spot.types import QueryKlines, KlineOutput


if __name__ == '__main__':
    schema = ToolSchema(
        name="k_lines",
        description="Query Klines for a symbol. Klines are uniquely identified by their open time",
        args_schema=ParamSchema.from_model_type(QueryKlines),
        result_schema=ParamSchema.from_model_type(KlineOutput),
        metadata={
            "output_tag": "kline"
        }
    )
    print(schema.json(by_alias=True, exclude_none=True))

    # print("\n===============Running Tool===============\n")
    #
    # args = QueryKlines(
    #     symbol="ETHUSDT",
    # ).dict(by_alias=True)
    # print(args)
    # resp = schema.run_tool("../../go-tools/outputs/binance_spot.so", **args)
    # if resp is not None:
    #     print(json.dumps(resp, indent=2))
