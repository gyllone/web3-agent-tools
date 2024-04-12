#ifndef NEWS_H
#define NEWS_H

#include <tools.h>

DEFINE_OPTIONAL(String);
DEFINE_OPTIONAL(Int);

typedef struct {
    String title;
    String content;
    String timestamp;
} NewsItem;
void release_NewsItem(NewsItem item);

DEFINE_LIST(NewsItem)
DEFINE_RESULT(List_NewsItem)

#endif