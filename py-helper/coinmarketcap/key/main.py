import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


class InfoInput(BaseModel):
    pass


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Plan(BaseModel):
    credit_limit_monthly: int
    credit_limit_monthly_reset: str
    credit_limit_monthly_reset_UTC: str
    rate_limit_minute: int


class CurrentMinute(BaseModel):
    requests_made: int
    requests_left: int


class CurrentDay(BaseModel):
    credits_used: int


class CurrentMonth(BaseModel):
    credits_used: int
    credits_left: int


class Usage(BaseModel):
    current_minute: CurrentMinute
    current_day: CurrentDay
    current_month: CurrentMonth


class Info(BaseModel):
    plan: Plan
    usage: Usage


class InfoOutput(Result):
    value: Optional[Info] = Field(description="info")


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_info",
        description="This is a test",
        args_schema=ParamSchema.from_model_type(InfoInput),
        result_schema=ParamSchema.from_model_type(InfoOutput),
    )

    print("\n===============Running Tool===============\n")

    args = InfoInput()

    resp = schema.run_tool("../../../go-tools/output/key.so", **args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
