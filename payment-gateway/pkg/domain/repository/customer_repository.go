package repository

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
)

type CustomerRepository interface {
	Save(ctx context.Context, customer *entity.Customer) error
	FindByID(ctx context.Context, id string) (*entity.Customer, error)
	FindByEmail(ctx context.Context, email string) (*entity.Customer, error)
}
