package models

import "time"

type Request struct {
	URL         string        `json :"url"`
	CustomShort string        `json : "short"`
	Expiry      time.Duration `json :"expiry"`
}

type Response struct {
	URL             string        `json :"url"`
	CustomShort     string        `json : "short"`
	Expiry          time.Duration `json :"expiry"`
	xRateRemaining  int           `json:"rate_limit`
	xRateLimitReset time.Duration `json:"rate_limit_reset"`
}