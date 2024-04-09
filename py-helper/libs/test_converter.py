from typing import List, Dict, Optional
from pydantic import BaseModel, Field

from libs.schema import Schema


class Baz(BaseModel):
    foo2: bool = Field(description="Foo2")
    bar2: Dict[str, int] = Field(description="Bar2")
    baz2: Optional[str] = Field(description="Baz2")


class Bar(BaseModel):
    foo1: Baz = Field(description="Foo1")
    bar1: Dict[str, Baz] = Field(description="Bar1")


class Foo(BaseModel):
    foo: Optional[float] = Field(description="Foo")
    bar: List[List[Bar]] = Field(description="Bar")
    baz: Dict[str, List[Baz]] = Field(description="Baz")


if __name__ == "__main__":
    raw_schema = Foo.schema()
    schema = Schema.parse_obj(raw_schema)
    print(schema.json(indent=2, exclude_none=True))

    # converter = ValueConverter(schema)
    # argtypes = converter.get_arg_types()
    # # print(argtypes)
    # structure = converter.get_structure_type()
    # # print(structure)
    #
    # py_values = Foo(
    #     foo=1.0,
    #     bar=[
    #         [Bar(
    #             foo1=Baz(foo2=True, bar2={"x": 1, "y": 2}, baz2="foo"),
    #             bar1={"x": Baz(foo2=True, bar2={"x": 1, "y": 2}, baz2="foo")},
    #         )],
    #         [Bar(
    #             foo1=Baz(foo2=False, bar2={"x": 3, "y": 4}, baz2=None),
    #             bar1={"y": Baz(foo2=False, bar2={"x": 3, "y": 4}, baz2=None)},
    #         )],
    #     ],
    #     baz={
    #         "x": [Baz(foo2=True, bar2={"x": 1, "y": 2}, baz2="foo")],
    #         "y": [Baz(foo2=False, bar2={"x": 3, "y": 4}, baz2=None)],
    #     }
    # )
    # print(py_values.json())
    # argvalues = converter.py_object_to_c_values(py_values.dict())
    #
    # c_structure = structure(
    #     foo=argvalues[0],
    #     bar=argvalues[1],
    #     baz=argvalues[2],
    # )
    # py_object = converter.c_struct_to_py_object(c_structure)
    # print(json.dumps(py_object))
