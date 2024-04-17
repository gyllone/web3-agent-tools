#ifndef KEY_H
#define KEY_H

#include <tools.h>

typedef struct {
    Int credit_limit_monthly;
    String credit_limit_monthly_reset;
    String credit_limit_monthly_reset_UTC;
    Int rate_limit_minute;
} Plan;

void release_Plan(Plan data);

typedef struct {
    Int requests_made;
    Int requests_left;
} CurrentMinute;

void release_CurrentMinute(CurrentMinute data);

typedef struct {
    Int credits_used;
} CurrentDay;

void release_CurrentDay(CurrentDay data);

typedef struct {
    Int credits_used;
    Int credits_left;
} CurrentMonth;

void release_CurrentMonth(CurrentMonth data);

typedef struct {
    CurrentMinute current_minute;
    CurrentDay current_day;
    CurrentMonth current_month;
} Usage;

void release_Usage(Usage data);

typedef struct {
    Plan plan;
    Usage usage;
} Info;

void release_Info(Info data);

DEFINE_OPTIONAL(Info)

DEFINE_RESULT(Optional_Info)

#endif