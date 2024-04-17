#include <tools.h>
#include <key.h>

void release_Plan(Plan data) {
    free(data.credit_limit_monthly_reset);
    free(data.credit_limit_monthly_reset_UTC);
}

void release_CurrentMinute(CurrentMinute data) {}

void release_CurrentDay(CurrentDay data) {}

void release_CurrentMonth(CurrentMonth data) {}

void release_Usage(Usage data) {
    release_CurrentMinute(data.current_minute);
    release_CurrentDay(data.current_day);
    release_CurrentMonth(data.current_month);
}

void release_Info(Info data) {
    release_Plan(data.plan);
    release_Usage(data.usage);
}

IMPL_OPTIONAL(Info)

IMPL_RESULT(Optional_Info)
