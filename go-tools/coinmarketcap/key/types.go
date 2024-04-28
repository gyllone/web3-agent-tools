package main

import "time"

type StatusResp struct {
	ErrorMessage string `json:"error_message"`
}

type Plan struct {
	CreditLimitMonthly               int64     `json:"credit_limit_monthly"`
	CreditLimitMonthlyReset          string    `json:"credit_limit_monthly_reset"`
	CreditLimitMonthlyResetTimestamp time.Time `json:"credit_limit_monthly_reset_timestamp"`
	RateLimitMinute                  int64     `json:"rate_limit_minute"`
}

type Usage struct {
	CurrentMinute struct {
		RequestsMade int64 `json:"requests_made"`
		RequestsLeft int64 `json:"requests_left"`
	} `json:"current_minute"`
	CurrentDay struct {
		CreditsUsed int64 `json:"credits_used"`
	} `json:"current_day"`
	CurrentMonth struct {
		CreditsUsed int64 `json:"credits_used"`
		CreditsLeft int64 `json:"credits_left"`
	} `json:"current_month"`
}

type Info struct {
	Plan  Plan  `json:"plan"`
	Usage Usage `json:"usage"`
}

type InfoResp struct {
	Data   Info       `json:"data"`
	Status StatusResp `json:"status"`
}
