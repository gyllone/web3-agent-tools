from typing import Optional, List, Dict
from pydantic import BaseModel, Field

from libs.tool import ToolSchema


# class FooArgs(BaseModel):
#     foo: float = Field(description="Foo")
#     bar: List["Bar"] = Field(description="Bar")
#
#     class Bar(BaseModel):
#         foo1: "Baz" = Field(description="Foo1")
#         bar1: int = Field(description="Bar1")
#
#         class Baz(BaseModel):
#             foo2: bool = Field(description="Foo2")
#             bar2: Dict[str, int] = Field(description="Bar2")
#             baz2: Optional[str] = Field(description="Baz2")
#
#
# class FooResp(BaseModel):
#     foo: str
#     bar: List[str]
#     baz: Dict[str, float]


class TVLQueryArgs(BaseModel):
    protocol: Optional[str] = Field(None, description="Protocol or project name")
    blockchain: Optional[str] = Field(None, description="Blockchain name")


class TVLQueryResult(BaseModel):
    tvl: float = Field(description="Total value locked in USD")


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_tvl",
        description="query tvl data from defillama",
        args_schema=TVLQueryArgs.schema(),
        result_schema=TVLQueryResult.schema(),
    )
    # print tool schema
    print(schema.json(indent=2))

    print("\n===============Running Tool===============\n")

    args = TVLQueryArgs(
        protocol="lido",
        blockchain="ethereum",
    )
    try:
        resp = schema.run_tool("../go-tools/outputs/defillama.so", args)
        if resp is not None:
            print(resp.json(indent=2))
    except Exception as e:
        print(e)
