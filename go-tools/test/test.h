#ifndef TEST_H
#define TEST_H

#include <tools.h>

DEFINE_LIST(Float)

DEFINE_DICT(Int)

DEFINE_OPTIONAL(String)

DEFINE_LIST(Dict_Int)

typedef struct {
	Bool foo;
	List_Float bar;
	Dict_Int baz;
} Param;

void release_Param(Param param);

DEFINE_OPTIONAL(Param)

typedef struct {
	Bool status;
	String err;
	Optional_Param result;
} Output;

void release_Output(Output output);

#endif