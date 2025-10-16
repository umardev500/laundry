package types

import "strings"

type BillingInterval string

const (
	BillingIntervalMonthly BillingInterval = "MONTHLY"
	BillingIntervalYearly  BillingInterval = "YEARLY"
)

func (b BillingInterval) Normalize() BillingInterval {
	return BillingInterval(strings.ToUpper(string(b)))
}
