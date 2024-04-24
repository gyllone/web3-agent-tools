import json

from libs.tool import ToolSchema
from libs.schema import ParamSchema

from binance_spot.types import CreateOrderRequest, OrderResponseOutput


if __name__ == '__main__':
    schema = ToolSchema(
        name="create_order",
        description="Send in a new order",
        args_schema=ParamSchema.from_model_type(CreateOrderRequest),
        result_schema=ParamSchema.from_model_type(OrderResponseOutput),
    )
    print(schema.json(by_alias=True, exclude_none=True))

    # print("\n===============Running Tool===============\n")
    #
    # args = CreateOrderRequest(
    #     api_key="Mk2ZrUlBth4H3H01FCF77Iv30WSjXKlDKV0fc1m6ZlSBMeOAUsi2NvAqgUEsPGYm",
    #     secret_key="CGyRwxe8aUsfpBKWfNSz5MtN1sX8e3vvz9dlNE52mDvzOHvi11TIGZXGXJUPhZve",
    #     symbol="BTCUSDT",
    #     side="SELL",
    #     order_type="LIMIT",
    #     price=80000,
    #     quantity=0.001,
    #     time_in_force="GTC",
    # ).dict(by_alias=True)
    # print(args)
    # resp = schema.run_tool("../../go-tools/outputs/binance_spot.so", **args)
    # if resp is not None:
    #     print(json.dumps(resp, indent=2))
