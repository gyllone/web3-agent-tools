from typing import Literal, Optional, Any, Type
from pydantic import BaseModel, Field, validator


class Property(BaseModel):
    description: Optional[str] = None
    default: Optional[Any] = None
    ref: Optional[str] = Field(None, alias="$ref")
    allOf: Optional[list["Property"]] = None
    type: Optional[Literal["string", "number", "integer", "boolean", "array", "object"]] = None
    items: Optional["Property"] = None
    additionalProperties: Optional["Property"] = None

    def get_ref_type(self) -> Optional[str]:
        return self.ref.split("/")[-1] if self.ref else None

    @validator("allOf")
    def validate_all_of(cls, value: Optional[list["Property"]]) -> Optional[list["Property"]]:
        if value is not None:
            if len(value) != 1:
                raise ValueError("allOf should have exactly one element")
        return value

    def check_ref(self, type_name: str):
        ref_name = self.get_ref_type()
        if ref_name == type_name:
            raise ValueError(f"Type {type_name} references itself")

        if self.allOf:
            self.allOf[0].check_ref(type_name)
        if self.items:
            self.items.check_ref(type_name)
        if self.additionalProperties:
            self.additionalProperties.check_ref(type_name)


class Object(BaseModel):
    properties: Optional[dict[str, Property]] = None
    required: Optional[list[str]] = None


class Schema(Object):
    definitions: Optional[dict[str, Object]] = None

    @classmethod
    def from_model_type(cls, model: Type[BaseModel]) -> "Schema":
        schema = model.schema()
        return cls.parse_obj(schema)

    @validator("definitions")
    def validate_definitions(cls, values: Optional[dict[str, Object]]) -> Optional[dict[str, Object]]:
        if values:
            for tp_name, obj in values.items():
                for prop_name, prop in obj.properties.items():
                    prop.check_ref(tp_name)
        return values
