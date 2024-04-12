#include <news.h>

IMPL_OPTIONAL(String);
IMPL_OPTIONAL(Int);

void release_NewsItem(NewsItem item) {
    free(item.title);
    free(item.content);
    free(item.timestamp);
}

IMPL_LIST(NewsItem)
IMPL_RESULT(List_NewsItem)