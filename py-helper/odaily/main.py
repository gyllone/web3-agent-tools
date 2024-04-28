import json
from typing import Optional

from pydantic import BaseModel
from pydantic import Field

from libs.schema import ParamSchema
from libs.tool import ToolSchema


class News(BaseModel):
    id: int
    is_top: int
    title: str
    description: str
    cover: str
    news_url: str
    extraction_tags: str
    updated_at: str


class NewsflashesInput(BaseModel):
    per_page: Optional[int]
    is_import: Optional[bool]


class NewsflashesOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: list[News]


class PostInput(BaseModel):
    type_str: Optional[str]


class Post(BaseModel):
    id: int
    title: str
    summary: str
    updated_at: str


class PostOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[list[Post]]


if __name__ == '__main__':
    schema = ToolSchema(
        # name="query_newsflashes",
        name="query_post_list",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(NewsflashesInput),
        # result_schema=ParamSchema.from_model_type(NewsflashesOutput),
        args_schema=ParamSchema.from_model_type(PostInput),
        result_schema=ParamSchema.from_model_type(PostOutput),
        metadata={
            "annotation": "*querying from odaily*\n"
        }
    )

    print("\n===============Running Tool===============\n")

    # args = NewsflashesInput(per_page=2)
    args = PostInput()
    resp = schema.run_tool("../../go-tools/output/odaily.so", **args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
