#ifndef MARSBIT_H
#define MARSBIT_H

#include <tools.h>

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(String)

typedef struct {
    String news_type;
    String title;
    String content;
    String synopsis;
    String publish_time;
} News;

void release_News(News data);

DEFINE_LIST(News)

DEFINE_OPTIONAL(List_News)

DEFINE_RESULT(List_News)




typedef struct {
    String title;
    String content;
    String synopsis;
    String publish_time;
} SearchedNews;

void release_SearchedNews(SearchedNews data);

DEFINE_LIST(SearchedNews)

typedef struct {
    List_SearchedNews News;
    List_SearchedNews Lives;
    List_SearchedNews ExcellentNews;
}SearchedNewsObj;

void release_SearchedNewsObj(SearchedNewsObj data);

DEFINE_OPTIONAL(SearchedNewsObj)

DEFINE_RESULT(SearchedNewsObj)
#endif