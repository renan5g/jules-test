package entity

import "time"

type Merchant struct {
	ID          string
	Name        string
	Document    string // e.g., CNPJ, EIN
	APIKey      string // For authenticating merchant requests
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
