import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.tool import ToolSchema, ParamSchema


class Input(BaseModel):
    start: Optional[int]
    limit: Optional[int]
    sort: Optional[str]
    include_metals: Optional[bool]


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Fiat(BaseModel):
    id: int
    name: str
    sign: str
    symbol: str


class Output(Result):
    value: Optional[list[Fiat]]


if __name__ == '__main__':
    schema = ToolSchema(
        name="query_id_map",
        description="This is a test",
        args_schema=ParamSchema.from_model_type(Input),
        result_schema=ParamSchema.from_model_type(Output),
    )

    print("\n===============Running Tool===============\n")

    args = Input(start=5, limit=2, sort="name", include_metals=True)

    resp = schema.run_tool("../../../go-tools/output/fiat.so", **args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
