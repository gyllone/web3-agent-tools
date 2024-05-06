#include <tools.h>
#include <twitter.h>

void release_TimelineInfo(TimelineInfo data) {
    free(data.id);
    free(data.text);
    free(data.created_at);
}

IMPL_LIST(TimelineInfo)

IMPL_OPTIONAL(List_TimelineInfo)

IMPL_RESULT(List_TimelineInfo)

IMPL_DICT(Result_List_TimelineInfo)

IMPL_OPTIONAL(Dict_Result_List_TimelineInfo)

IMPL_RESULT(Dict_Result_List_TimelineInfo)