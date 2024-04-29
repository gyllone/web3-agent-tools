#ifndef TWITTER_H
#define TWITTER_H

#include <tools.h>

typedef struct {
    String id;
    String name;
    String username;
} User;

void release_User(User data);

DEFINE_LIST(User)

DEFINE_OPTIONAL(List_User)

DEFINE_RESULT(List_User)


typedef struct {
    String id;
    String name;
    String description;
} Domain;

void release_Domain(Domain data);

typedef struct {
    String id;
    String name;
    String description;
} Entity;

void release_Entity(Entity data);

typedef struct {
    Domain domain;
    Entity entity;
} ContextAnnotation;

void release_ContextAnnotation(ContextAnnotation data);

DEFINE_LIST(ContextAnnotation);

typedef struct Tweet {
    String id;
    String text;
    List_ContextAnnotation context_annotations;
    String created_at;
} Tweet;

void release_Tweet(Tweet data);

typedef struct {
    Tweet tweet;
    User author;
} TimelineInfo;

void release_TimelineInfo(TimelineInfo data);

DEFINE_LIST(TimelineInfo)

DEFINE_OPTIONAL(List_TimelineInfo)

DEFINE_RESULT(List_TimelineInfo)

#endif