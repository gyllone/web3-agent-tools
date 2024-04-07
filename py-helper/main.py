from typing import Optional, TypedDict
from pydantic import BaseModel, Field

from libs.tool import ToolSchema


class V(BaseModel):
    a: str
    b: int


class TVLQueryArgs(BaseModel):
    name: Optional[str] = Field(None, description="Protocol or project name")
    blockchain: Optional[str] = Field(None, description="Blockchain name")
    t: Optional[V] = Field(None, description="test")


class TVLQueryResult(BaseModel):
    status: bool = Field(description="if query is successful or not")
    error: str = Field(description="error message if status is false")
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

    args = TVLQueryArgs(name="lido", blockchain="ethereum")
    resp = schema.run_tool("../go-tools/outputs/defillama.so", args)
    if resp is not None:
        print(resp.json(indent=2))
