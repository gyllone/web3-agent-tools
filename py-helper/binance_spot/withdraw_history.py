import json

from libs.tool import ToolSchema
from libs.schema import ParamSchema

from binance_spot.types import QueryWithdrawHistory, WithdrawHistoryOutput


if __name__ == '__main__':
    schema = ToolSchema(
        name="withdraw_history",
        description="Fetch withdraw history",
        args_schema=ParamSchema.from_model_type(QueryWithdrawHistory),
        result_schema=ParamSchema.from_model_type(WithdrawHistoryOutput),
    )
    print(schema.json(by_alias=True, exclude_none=True))

    # print("\n===============Running Tool===============\n")
    #
    # args = QueryWithdrawHistory(
    #     api_key="Mk2ZrUlBth4H3H01FCF77Iv30WSjXKlDKV0fc1m6ZlSBMeOAUsi2NvAqgUEsPGYm",
    #     secret_key="CGyRwxe8aUsfpBKWfNSz5MtN1sX8e3vvz9dlNE52mDvzOHvi11TIGZXGXJUPhZve",
    # ).dict(by_alias=True)
    # resp = schema.run_tool("../../go-tools/outputs/binance_spot.so", **args)
    # if resp is not None:
    #     print(json.dumps(resp, indent=2))
