#include <tools.h>
#include <cryptocurrency.h>

//IMPL_OPTIONAL(String)
//
//IMPL_OPTIONAL(Int)
//
//IMPL_OPTIONAL(Bool)

IMPL_LIST(Float)

IMPL_LIST(List_Float)

IMPL_DICT(String)

IMPL_LIST(Dict_String)

void release_QuoteResult(QuoteResult result) {
    release_Bool(result.is_fail);
    release_String(result.error_message);
    if (result.is_fail== false) {
        release_List_List_Float(result.data);
    }
}

void release_IdMapResult(IdMapResult result) {
    release_Bool(result.is_fail);
    release_String(result.error_message);
    if (result.is_fail == false) {
        release_List_Dict_String(result.data);
    }
}

void release_MetadataResult(MetadataResult result) {
    release_Bool(result.is_fail);
    release_String(result.error_message);
    if (result.is_fail == false) {
        release_List_Dict_String(result.data);
    }
}