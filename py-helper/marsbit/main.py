import json
from typing import Optional

from pydantic import BaseModel
from pydantic import Field

from libs.schema import ParamSchema
from libs.tool import ToolSchema


class News(BaseModel):
    news_type: str
    title: str
    content: str
    synopsis: str
    publish_time: str


class NewsInput(BaseModel):
    query_time: Optional[int]
    page_size: Optional[int]
    lang: Optional[str]


class NewsOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: list[News]


class SearchedNewsObjInput(BaseModel):
    q: str
    page: Optional[int]
    page_size: Optional[int]


class SearchedNews(BaseModel):
    title: str
    content: str
    synopsis: str
    publish_time: str


class SearchedNewsObj(BaseModel):
    News: list[SearchedNews]
    Lives: list[SearchedNews]
    ExcellentNews: list[SearchedNews]


class SearchedNewsObjOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[SearchedNewsObj]


if __name__ == '__main__':
    schema = ToolSchema(
        # name="query_news",
        name="query_multisearch",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(NewsInput),
        # result_schema=ParamSchema.from_model_type(NewsOutput),
        args_schema=ParamSchema.from_model_type(SearchedNewsObjInput),
        result_schema=ParamSchema.from_model_type(SearchedNewsObjOutput),
        metadata={
            "annotation": "*querying from odaily*\n"
        }
    )

    print("\n===============Running Tool===============\n")

    # args = NewsInput(page_size=5, lang="en")
    args = SearchedNewsObjInput(q="bit")
    resp = schema.run_tool("../../go-tools/output/marsbit.so", **args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
