package repository

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type TransactionRepository interface {
	Save(ctx context.Context, transaction *entity.Transaction) error
	FindByID(ctx context.Context, id string) (*entity.Transaction, error)
	// Add other methods like FindByPaymentID, etc.
}
