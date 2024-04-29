import json
from typing import Optional

from pydantic import BaseModel
from pydantic import Field

from libs.schema import ParamSchema
from libs.tool import ToolSchema


class User(BaseModel):
    id: str
    name: str
    username: str


class UserInput(BaseModel):
    username: str


class UserOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[list[User]]


class Domain(BaseModel):
    id: str
    name: str
    description: str


class ContextAnnotation(BaseModel):
    domain: Domain
    entity: Domain


class Tweet(BaseModel):
    id: str
    text: str
    context_annotations: list[ContextAnnotation]
    create_at: str


class TimelineInfo(BaseModel):
    tweet: Tweet
    author: User


class TimelineInfoInput(BaseModel):
    id: str


class TimelineInfoOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[list[TimelineInfo]]


if __name__ == '__main__':
    schema = ToolSchema(
        # name="query_user_ids",
        # name="query_user_timeline",
        name="query_user_mention_timeline",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(UserInput),
        # result_schema=ParamSchema.from_model_type(UserOutput),
        args_schema=ParamSchema.from_model_type(TimelineInfoInput),
        result_schema=ParamSchema.from_model_type(TimelineInfoOutput),
        metadata={
            "annotation": "*querying from twitter*\n"
        }
    )

    print("\n===============Running Tool===============\n")

    # args = UserInput(username="G2NiKo")
    args = TimelineInfoInput(id="3351760203")

    resp = schema.run_tool("../../go-tools/output/twitter.so", **args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
