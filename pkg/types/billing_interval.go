package types

type BillingInterval string

const (
	BillingIntervalMonthly BillingInterval = "MONTHLY"
	BillingIntervalYearly  BillingInterval = "YEARLY"
)
