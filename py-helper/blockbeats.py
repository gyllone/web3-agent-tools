from typing import Optional, List
from pydantic import BaseModel, Field

from libs.tool import ToolSchema
from libs.schema import ParamSchema
import json


class BlockBeatsNewsArgs(BaseModel):
    date: Optional[str] = Field(None, description="news date in format %Y-%m-%d, like 2023-01-01")
    limit: Optional[int] = Field(None, description="limitation of size")


class NewsItem(BaseModel):
    title: str = Field(description="title of news")
    content: str = Field(description="content of news")
    timestamp: str = Field(description="timestamp of news")


class BlockBeatsNewsResult(BaseModel):
    status: bool = Field(description="if the request is success")
    error: str = Field(description="error message")
    its: Optional[list[NewsItem]] = Field(None, description="list of news")


if __name__ == '__main__':
    schema = ToolSchema(
        name="get_blockbeats_news",
        description="query from blockbeats news",
        args_schema=ParamSchema.from_model_type(BlockBeatsNewsArgs),
        result_schema=ParamSchema.from_model_type(BlockBeatsNewsResult),
        metadata={
            "annotation": "*querying from blockbeats*\n"
        }
    )
    # print tool schema
    print(schema.json(indent=2, by_alias=False, exclude_none=True))

    print("\n===============Running Tool===============\n")

    args = BlockBeatsNewsArgs(
        # date="2024-04-10",
        limit=2
    )

    resp = schema.run_tool("../go-tools/flash_news/outputs/flash_news.so", **args.dict(by_alias=True, exclude_none=True))
    if resp is not None:
        print(json.dumps(resp, indent=2, ensure_ascii=False))
