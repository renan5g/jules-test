package entity

import (
	"time"
	// "github.com/shopspring/decimal"
)

type TransactionType string

const (
	TransactionTypeAuthorization TransactionType = "AUTHORIZATION"
	TransactionTypeCapture       TransactionType = "CAPTURE"
	TransactionTypeSale          TransactionType = "SALE" // For direct sale without separate auth/capture
	TransactionTypeRefund        TransactionType = "REFUND"
	TransactionTypeVoid          TransactionType = "VOID"
)

type TransactionStatus string

const (
	TransactionStatusSucceeded TransactionStatus = "SUCCEEDED"
	TransactionStatusFailed    TransactionStatus = "FAILED"
	TransactionStatusPending   TransactionStatus = "PENDING"
)

type Transaction struct {
	ID                 string
	PaymentID          string
	Type               TransactionType
	Amount             float64 // Placeholder; use decimal.Decimal
	Status             TransactionStatus
	ProviderResponse   string // Raw response, or more structured
	ProviderFee        float64 // Placeholder; use decimal.Decimal
	GatewayFee         float64 // Placeholder; use decimal.Decimal
	CreatedAt          time.Time
}
