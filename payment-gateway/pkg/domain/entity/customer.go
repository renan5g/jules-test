package entity

import "time"

type Customer struct {
	ID        string
	Name      string
	Email     string
	Document  string // e.g., CPF, SSN
	CreatedAt time.Time
	UpdatedAt time.Time
}
