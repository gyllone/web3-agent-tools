#include <tools.h>
#include <twitter.h>

void release_User(User data) {
    free(data.id);
    free(data.name);
    free(data.username);
}

IMPL_LIST(User)

IMPL_OPTIONAL(List_User)

IMPL_RESULT(List_User)

void release_Domain(Domain data) {
    free(data.id);
    free(data.name);
    free(data.description);
}

void release_Entity(Entity data) {
    free(data.id);
    free(data.name);
    free(data.description);
}

void release_ContextAnnotation(ContextAnnotation data) {
    release_Domain(data.domain);
    release_Entity(data.entity);
}

IMPL_LIST(ContextAnnotation)

void release_Tweet(Tweet data) {
    free(data.id);
    free(data.text);
    release_List_ContextAnnotation(data.context_annotations);
    free(data.created_at);
}

void release_TimelineInfo(TimelineInfo data) {
    release_Tweet(data.tweet);
    release_User(data.author);
}

IMPL_LIST(TimelineInfo)

IMPL_OPTIONAL(List_TimelineInfo)

IMPL_RESULT(List_TimelineInfo)