#ifndef GLOBAL_METRICS_H
#define GLOBAL_METRICS_H

#include <tools.h>

DEFINE_OPTIONAL(String)

typedef struct {
    Float total_market_cap;
    Float total_volume_24h;
    Float total_volume_24h_reported;
    Float altcoin_volume_24h;
    Float altcoin_volume_24h_reported;
    Float altcoin_market_cap;
    Float defi_volume_24h;
    Float defi_volume_24h_reported;
    Float defi_24h_percentage_change;
    Float defi_market_cap;
    Float stablecoin_volume_24h;
    Float stablecoin_volume_24h_reported;
    Float stablecoin_24h_percentage_change;
    Float stablecoin_market_cap;
    Float derivatives_volume_24h;
    Float derivatives_volume_24h_reported;
    Float derivatives_24h_percentage_change;
    Float total_market_cap_yesterday;
    Float total_volume_24h_yesterday;
    Float total_market_cap_yesterday_percentage_change;
    Float total_volume_24h_yesterday_percentage_change;
    String last_updated;
} Quote;

void release_Quote(Quote data);

DEFINE_DICT(Quote)

typedef struct {
    Int active_cryptocurrencies;
    Int total_cryptocurrencies;
    Int active_market_pairs;
    Int active_exchanges;
    Int total_exchanges;
    Float eth_dominance;
    Float btc_dominance;
    Float eth_dominance_yesterday;
    Float btc_dominance_yesterday;
    Float eth_dominance_24h_percentage_change;
    Float btc_dominance_24h_percentage_change;
    Float defi_volume_24h;
    Float defi_volume_24h_reported;
    Float defi_market_cap;
    Float defi_24h_percentage_change;
    Float stablecoin_volume_24h;
    Float stablecoin_volume_24h_reported;
    Float stablecoin_market_cap;
    Float stablecoin_24h_percentage_change;
    Float derivatives_volume_24h;
    Float derivatives_volume_24h_reported;
    Float derivatives_24h_percentage_change;
    Dict_Quote quote;
    String last_updated;
} Metric;

void release_Metric(Metric data);

DEFINE_OPTIONAL(Metric)

DEFINE_RESULT(Optional_Metric)

#endif