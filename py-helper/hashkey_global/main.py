from typing import Optional, List
from pydantic import BaseModel, Field

from libs.tool import ToolSchema
from libs.schema import ParamSchema
import json
from .tool_types import *

project = 'hashkey_global'
ApiKey="gnb0CeMO43AJgkF6OwIaES7bvF0SoJm59mEu2VxsSSjHJgyX3jwYIEcSm4jDYkvd"
Secret="Hy0Y5SkAB5SK28PxniiTxhUsfMImXoh96YcnZcKMtEDznA3oogHyuI9ytLq1oyVd"

schemas = [
    {
        "schema": ToolSchema(
            name="get_spot_trade_account_balance",
            description="get hashkey spot trade account balance",
            args_schema=ParamSchema.from_model_type(GetTradeAccountBalanceArgs),
            result_schema=ParamSchema.from_model_type(GetTradeAccountBalanceResult),
            metadata={
                "annotation": "*querying balance from hashkey...*\n"
            }
        ),
        "case": GetTradeAccountBalanceArgs(
            ApiKey=ApiKey,
            ApiSignKey=Secret
        )
    },
    {    
        "schema": ToolSchema(
            name="create_spot_market_order",
            description="create spot order to hashkey global exchange by market price",
            args_schema=ParamSchema.from_model_type(CreateSpotMarketOrderArgs),
            result_schema=ParamSchema.from_model_type(CreateSpotMarketOrderResult),
            metadata={
                "annotation": "*creating spot order with market price to hashkey...*\n"
            }
        ),
        "case":None
        # "case": CreateSpotMarketOrderArgs(
        #     ApiKey=ApiKey,
        #     Secret=Secret,
        #     Symbol="BTCUSDT",
        #     Side="BUY",
        #     Quantity="0.0001"
        # )
    }, {
        "schema": ToolSchema(
            name="get_kline",
            description="query kline data of specific syboml from hashkey global exchange",
            args_schema=ParamSchema.from_model_type(GetKlineArgs),
            result_schema=ParamSchema.from_model_type(GetKlineResult),
            metadata={
               "output_tag":"kline"
            }
        ),
        "case":GetKlineArgs(
            Symbol="BTCUSDT",
            Interval="15m",
            Limit=5
        )
    }

]


def write_json(jsonStr: str, fileName: str="schema.json"):
    with open(fileName, 'w', encoding='utf-8') as file:
        file.write(jsonStr)

if __name__ == '__main__':
    
    jsonStr = ',\n'.join([schema["schema"].json(indent=2, exclude_none=True) for schema in schemas])
    print(f'[{jsonStr}]')
    write_json(f'[{jsonStr}]', f'{project}/schema.json')
    print("\n===============Running Tool===============\n")

    for schema in schemas:
        if schema["case"] is None:
            continue
        resp = schema["schema"].run_tool(f"../go-tools/{project}/outputs/{project}.so", schema["case"].dict(by_alias=True, exclude_none=True))
        if resp is not None:
            print(json.dumps(resp, indent=2, ensure_ascii=False))
