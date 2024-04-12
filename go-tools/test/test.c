#include <test.h>

IMPL_LIST(Float)

IMPL_DICT(Int)

IMPL_LIST(Dict_Int)

IMPL_OPTIONAL(Param)

void release_Param(Param param) {
    release_Bool(param.foo);
    release_List_Float(param.bar);
    release_Dict_Int(param.baz);
}

void release_Output(Output output) {
    release_Bool(output.status);
    release_String(output.err);
    release_Optional_Param(output.result);
}