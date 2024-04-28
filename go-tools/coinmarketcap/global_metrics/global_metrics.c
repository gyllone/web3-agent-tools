#include <tools.h>
#include <global_metrics.h>

IMPL_OPTIONAL(String)

void release_Quote(Quote data) {
    free(data.last_updated);
}

IMPL_DICT(Quote)

void release_Metric(Metric data) {
    release_Dict_Quote(data.quote);
    free(data.last_updated);
}

IMPL_OPTIONAL(Metric)

IMPL_RESULT(Metric)