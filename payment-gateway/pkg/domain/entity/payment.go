package entity

import (
	"time"
	// Consider using a proper decimal library, e.g., "github.com/shopspring/decimal"
)

type PaymentStatus string

const (
	PaymentStatusPending    PaymentStatus = "PENDING"
	PaymentStatusAuthorized PaymentStatus = "AUTHORIZED"
	PaymentStatusSucceeded  PaymentStatus = "SUCCEEDED"
	PaymentStatusFailed     PaymentStatus = "FAILED"
	PaymentStatusRefunded   PaymentStatus = "REFUNDED"
)

type Payment struct {
	ID                string
	Amount            float64 // Placeholder; use decimal.Decimal for real applications
	Currency          string
	Status            PaymentStatus
	MerchantID        string
	CustomerID        string // Optional
	PaymentMethodID   string
	Transactions      []Transaction
	CreatedAt         time.Time
	UpdatedAt         time.Time
}
