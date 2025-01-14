import json
from typing import Optional

from pydantic import BaseModel, Field

from libs.schema import ParamSchema
from libs.tool import ToolSchema


class Result(BaseModel):
    status: bool = Field(description="is_fail")
    error: str = Field(description="error_message")


class Platform(BaseModel):
    id: int
    name: str
    symbol: str
    slug: str
    token_address: str


class IdMapInput(BaseModel):
    listing_status: Optional[str] = Field(description="listing_status")
    sort: Optional[str] = Field(description="sort")
    symbol: Optional[str] = Field(description="symbol")
    aux: Optional[str] = Field(description="aux")
    start: Optional[int] = Field(description="start")
    limit: Optional[int] = Field(description="limit")


class Cryptocurrency(BaseModel):
    id: int
    rank: int
    name: str
    symbol: str
    slug: str
    is_active: int
    first_historical_data: str
    last_historical_data: str
    platform: Optional[Platform]


class IdMapOutput(Result):
    value: list[Cryptocurrency] = Field(description="idMaps")


class Quote(BaseModel):
    last_updated: str
    price: float
    volume_24h: float
    percent_change_1h: float
    percent_change_24h: float
    percent_change_7d: float
    percent_change_30d: float
    percent_change_60d: float
    percent_change_90d: float
    market_cap: float
    market_cap_dominance: float
    fully_diluted_market_cap: float
    tvl: float


class Tag(BaseModel):
    slug: str
    name: str
    category: str


class QuoteData(BaseModel):
    id: int
    name: str
    symbol: str
    slug: str
    num_market_pairs: int
    date_added: str
    tags: list[Tag]
    max_supply: float
    circulating_supply: float
    total_supply: float
    is_active: int
    infinite_supply: bool
    platform: Optional[Platform]
    cmc_rank: int
    is_fiat: int
    self_reported_circulating_supply: float
    self_reported_market_cap: float
    tvl_ratio: float
    last_updated: str
    quote: dict[str, Quote]


class QuoteInput(BaseModel):
    id: Optional[str]
    slug: Optional[str]
    convert: Optional[str]
    convert_id: Optional[str]
    aux: Optional[str]
    skip_invalid: Optional[bool]


class QuoteOutput(Result):
    value: dict[str, QuoteData] = Field(description="quotes")


class MetaInput(BaseModel):
    id: Optional[str]
    slug: Optional[str]
    address: Optional[str]
    aux: Optional[str]
    skip_invalid: Optional[bool]


class URLs(BaseModel):
    website: list[str]
    twitter: list[str]
    message_board: list[str]
    chat: list[str]
    facebook: list[str]
    explorer: list[str]
    reddit: list[str]
    technical_doc: list[str]
    source_code: list[str]
    announcement: list[str]


class Metadata(BaseModel):
    id: int
    name: str
    symbol: str
    category: str
    description: str
    slug: str
    logo: str
    subreddit: str
    notice: str
    tags: list[str]
    tag_names: list[str]
    tag_groups: list[str]
    urls: URLs
    platform: Optional[Platform]
    date_added: str
    twitter_username: str
    is_hidden: int
    date_launched: str
    self_reported_circulating_supply: float
    self_reported_market_cap: float
    infinite_supply: bool


class MetaOutput(Result):
    value: dict[str, Metadata] = Field(description="metas")


class ListingInput(BaseModel):
    start: Optional[int]
    limit: Optional[int]
    price_min: Optional[int]
    price_max: Optional[int]
    market_cap_min: Optional[int]
    market_cap_max: Optional[int]
    volume_24h_min: Optional[int]
    volume_24h_max: Optional[int]
    circulating_supply_min: Optional[int]
    circulating_supply_max: Optional[int]
    percent_change_24h_min: Optional[int]
    percent_change_24h_max: Optional[int]
    convert: Optional[str]
    convert_id: Optional[str]
    sort: Optional[str]
    sort_dir: Optional[str]
    cryptocurrency_type: Optional[str]
    tag: Optional[str]
    aux: Optional[str]


class Market(BaseModel):
    id: int
    name: str
    symbol: str
    slug: str
    num_market_pairs: int
    date_added: str
    tags: list[str]
    max_supply: int
    circulating_supply: int
    total_supply: int
    infinite_supply: bool
    platform: Optional[Platform]
    cmc_rank: int
    self_reported_circulating_supply: float
    self_reported_market_cap: float
    tvl_ratio: float
    last_updated: str
    quote: dict[str, Quote]


class ListingOutput(Result):
    value: list[Market] = Field(description="market data")


class CategoriesInput(BaseModel):
    start: Optional[int]
    limit: Optional[int]
    id: Optional[str]
    slug: Optional[str]
    symbol: Optional[str]


class Category(BaseModel):
    id: str
    name: str
    description: str
    num_tokens: int
    avg_price_change: float
    market_cap: float
    market_cap_change: float
    volume: float
    volume_change: float
    last_updated: str


class CategoriesOutput(Result):
    value: list[Category] = Field(description="categories")


class CategoryInput(BaseModel):
    id: str
    start: Optional[int]
    limit: Optional[int]
    convert: Optional[str]
    convert_id: Optional[str]


class Coin(BaseModel):
    id: int
    name: str
    symbol: str
    slug: str
    num_market_pairs: int
    date_added: str
    tags: list[str]
    max_supply: int
    circulating_supply: int
    total_supply: int
    is_active: int
    infinite_supply: bool
    cmc_rank: int
    is_fiat: int
    tvl_ratio: float
    last_updated: str
    quote: dict[str, Quote]


class CategorySingle(BaseModel):
    id: str
    name: str
    description: str
    num_tokens: int
    avg_price_change: float
    market_cap: float
    market_cap_change: float
    volume: float
    volume_change: float
    last_updated: str
    coins: list[Coin]


class CategoryOutput(Result):
    value: Optional[CategorySingle] = Field(description="category")


if __name__ == '__main__':
    schema = ToolSchema(
        # name="query_quotes_latest",
        # name="query_id_map",
        # name="query_metadata",
        # name="query_listings",
        # name="query_categories",
        name="query_category",
        description="This is a test",
        # args_schema=ParamSchema.from_model_type(QuoteInput),
        # result_schema=ParamSchema.from_model_type(QuoteOutput),
        # args_schema=ParamSchema.from_model_type(IdMapInput),
        # result_schema=ParamSchema.from_model_type(IdMapOutput),
        # args_schema=ParamSchema.from_model_type(MetaInput),
        # result_schema=ParamSchema.from_model_type(MetaOutput),
        # args_schema=ParamSchema.from_model_type(ListingInput),
        # result_schema=ParamSchema.from_model_type(ListingOutput),
        # args_schema=ParamSchema.from_model_type(CategoriesInput),
        # result_schema=ParamSchema.from_model_type(CategoriesOutput),
        args_schema=ParamSchema.from_model_type(CategoryInput),
        result_schema=ParamSchema.from_model_type(CategoryOutput),
    )

    print("\n===============Running Tool===============\n")

    # args = QuoteInput(id="2703")
    # args = IdMapInput(start=10, limit=2)
    # args = MetaInput(id="-1")
    # args = ListingInput(limit=5, convert="ETH,USD")
    # args = CategoriesInput(start=-1)
    args = CategoryInput(id="605e2ce9d41eae1066535f7c", limit=2)

    resp = schema.run_tool("../../../go-tools/output/crycur.so", **args.dict(by_alias=True, exclude_none=True))

    if resp is not None:
        print(json.dumps(resp, indent=2))
