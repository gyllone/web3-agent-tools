#include <tools.h>
#include <odaily.h>

IMPL_OPTIONAL(Int)

IMPL_OPTIONAL(Bool)

void release_News(News data) {
    free(data.title);
    free(data.description);
    free(data.cover);
    free(data.news_url);
    free(data.extraction_tags);
    free(data.updated_at);
}

IMPL_LIST(News)

IMPL_RESULT(List_News)



IMPL_OPTIONAL(String)

void release_Post (Post data) {
    free(data.title);
    free(data.summary);
    free(data.updated_at);
}

IMPL_LIST(Post)

IMPL_RESULT(List_Post)