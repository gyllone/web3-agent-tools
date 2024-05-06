#ifndef TWITTER_H
#define TWITTER_H

#include <tools.h>

typedef struct Tweet {
    String id;
    String text;
    String created_at;
} TimelineInfo;

void release_TimelineInfo(TimelineInfo data);

DEFINE_LIST(TimelineInfo)

DEFINE_OPTIONAL(List_TimelineInfo)

DEFINE_RESULT(List_TimelineInfo)

DEFINE_DICT(Result_List_TimelineInfo)

DEFINE_OPTIONAL(Dict_Result_List_TimelineInfo)

DEFINE_RESULT(Dict_Result_List_TimelineInfo)

#endif