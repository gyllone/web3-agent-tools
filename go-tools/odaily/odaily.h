#ifndef ODAILY_H
#define ODAILY_H

#include <tools.h>

DEFINE_OPTIONAL(Int)

DEFINE_OPTIONAL(Bool)

typedef struct {
    Int id;
    Int is_top;
    String title;
    String description;
    String cover;
    String news_url;
    String extraction_tags;
    String updated_at;
} News;

void release_News(News data);

DEFINE_LIST(News)

DEFINE_OPTIONAL(List_News)

DEFINE_RESULT(List_News)



DEFINE_OPTIONAL(String)

typedef struct {
    Int id;
    String title;
    String summary;
    String updated_at;
} Post;

void release_Post (Post data);

DEFINE_LIST(Post)

DEFINE_OPTIONAL(List_Post)

DEFINE_RESULT(List_Post)

#endif