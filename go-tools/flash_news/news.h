#ifndef NEWS_H
#define NEWS_H

#include <tools.h>

typedef struct {
    String title;
    String content;
    Int    unixtime;
} NewsItem;

DEFINE_LIST(NewsItem);

DEFINE_OPTIONAL(String);

DEFINE_OPTIONAL(Int);

typedef struct {
    Bool success;
    List_NewsItem items;
} NewsResult;

typedef struct {
    Bool success;
    String json_string;
} NewsResultJson;

void release_NewsItem(NewsItem item);
void release_NewsResult(NewsResult result);
void release_NewsResultJson(NewsResultJson result);

#endif