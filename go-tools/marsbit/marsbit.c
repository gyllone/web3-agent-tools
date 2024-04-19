#include <tools.h>
#include <marsbit.h>

IMPL_OPTIONAL(Int)

IMPL_OPTIONAL(String)

void release_News(News data) {
    free(data.news_type);
    free(data.title);
    free(data.content);
    free(data.synopsis);
    free(data.publish_time);
}

IMPL_LIST(News)

IMPL_RESULT(List_News)



void release_SearchedNews(SearchedNews data) {
    free(data.title);
    free(data.content);
    free(data.synopsis);
    free(data.publish_time);
}

IMPL_LIST(SearchedNews)

void release_SearchedNewsObj(SearchedNewsObj data) {
    release_List_SearchedNews(data.News);
    release_List_SearchedNews(data.Lives);
    release_List_SearchedNews(data.ExcellentNews);
}

IMPL_OPTIONAL(SearchedNewsObj)

IMPL_RESULT(Optional_SearchedNewsObj)