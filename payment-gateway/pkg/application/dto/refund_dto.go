package dto
import "time"
// import "github.com/shopspring/decimal"

type RefundPaymentInput struct {
	PaymentID     string
	Amount        float64 // Placeholder; use decimal.Decimal. Can be full or partial.
	Reason        string // Optional
	MerchantID    string // For validation
}

type RefundOutput struct {
	RefundID        string
	PaymentID       string
	Status          string // e.g., "SUCCEEDED", "PENDING", "FAILED"
	AmountRefunded  float64 // Placeholder
	Currency        string
	ProcessedAt     time.Time
}
