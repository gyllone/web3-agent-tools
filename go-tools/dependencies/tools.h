#ifndef TOOLS_H
#define TOOLS_H

#include <stdlib.h>
#include <stdbool.h>

typedef char* String;
typedef long long int Int;
typedef bool Bool;
typedef double Float;

static void release_String(String str) {
    free(str);
}

static void release_Int(Int i) {}

static void release_Bool(Bool b) {}

static void release_Float(Float f) {}

#define DEFINE_OPTIONAL(type) \
typedef struct { \
    Bool is_some; \
    type value; \
} Optional_##type; \
extern Optional_##type some_##type(type value); \
extern Optional_##type none_##type(); \
extern void release_Optional_##type(Optional_##type opt);

#define IMPL_OPTIONAL(type) \
Optional_##type some_##type(type value) { \
    Optional_##type opt; \
    opt.is_some = true; \
    opt.value = value; \
    return opt; \
} \
Optional_##type none_##type() { \
    Optional_##type opt; \
    opt.is_some = false; \
    return opt; \
} \
void release_Optional_##type(Optional_##type opt) { \
    if (opt.is_some) { \
        release_##type(opt.value); \
    } \
}

#define DEFINE_LIST(type) \
typedef struct { \
    size_t len; \
    type* values; \
} List_##type; \
extern List_##type new_List_##type(size_t len); \
extern void release_List_##type(List_##type list);

#define IMPL_LIST(type) \
List_##type new_List_##type(size_t len) { \
    List_##type list; \
    list.len = len; \
    if (len > 0) { \
        list.values = (type*)malloc(len * sizeof(type)); \
    } else { \
        list.values = NULL; \
    } \
    return list; \
} \
void release_List_##type(List_##type list) { \
    for (size_t i = 0; i < list.len; i++) { \
        release_##type(list.values[i]); \
    } \
    free(list.values); \
}

#define DEFINE_DICT(type) \
typedef struct { \
    size_t len; \
    String* keys; \
    type* values; \
} Dict_##type; \
\
extern Dict_##type new_Dict_##type(size_t len); \
extern void release_Dict_##type(Dict_##type dict);

#define IMPL_DICT(type) \
Dict_##type new_Dict_##type(size_t len) { \
    Dict_##type dict; \
    dict.len = len; \
    if (len > 0) { \
        dict.keys = (String*)malloc(len * sizeof(String)); \
        dict.values = (type*)malloc(len * sizeof(type)); \
    } else { \
        dict.keys = NULL; \
        dict.values = NULL; \
    } \
    return dict; \
} \
\
void release_Dict_##type(Dict_##type dict) { \
    for (size_t i = 0; i < dict.len; i++) { \
        free(dict.keys[i]); \
        release_##type(dict.values[i]); \
    } \
    free(dict.keys); \
    free(dict.values); \
}

#define DEFINE_RESULT(type) \
typedef struct { \
    Bool status; \
    String error; \
    type value; \
} Result_##type; \
extern Result_##type ok_##type(type value); \
extern Result_##type err_##type(String error); \
extern void release_Result_##type(Result_##type result);

#define IMPL_RESULT(type) \
Result_##type ok_##type(type value) { \
    Result_##type result; \
    result.status = true; \
    result.error = ""; \
    result.value = value; \
    return result; \
} \
Result_##type err_##type(String error) { \
    Result_##type result; \
    result.status = false; \
    result.error = error; \
    return result; \
} \
void release_Result_##type(Result_##type result) { \
    if (result.status) { \
        release_##type(result.value); \
    } else { \
        free(result.error); \
    } \
}

#endif