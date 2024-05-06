import json
from typing import Optional

from pydantic import BaseModel
from pydantic import Field

from libs.schema import ParamSchema
from libs.tool import ToolSchema


class TimelineInfo(BaseModel):
    id: str
    text: str
    create_at: str


class TimelineInfoInput(BaseModel):
    usernames: str
    max_results_per_user: int


class TimelineInfoOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[list[TimelineInfo]]


class TimelineInfosOutput(BaseModel):
    status: bool = Field(description="status")
    error: str = Field(description="error")
    value: Optional[dict[str, TimelineInfoOutput]]


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_users_timeline",
        # name="query_user_mention_timeline",
        description="This is a test",
        args_schema=ParamSchema.from_model_type(TimelineInfoInput),
        result_schema=ParamSchema.from_model_type(TimelineInfosOutput),
        metadata={
            "annotation": "*querying from twitter*\n"
        }
    )

    print("\n===============Running Tool===============\n")

    # args = UserInput(username="G2NiKo")
    args = TimelineInfoInput(usernames="G2NiKo,elonmusk", max_results_per_user=5)

    resp = schema.run_tool("../../go-tools/output/twitter.so", **args.dict(by_alias=True))
    if resp is not None:
        print(json.dumps(resp, indent=2))
