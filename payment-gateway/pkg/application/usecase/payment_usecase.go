package usecase

import (
	"context"
	"github.com/generic-org/payment-gateway/pkg/application/dto"
	"github.com/generic-org/payment-gateway/pkg/application/service"
	"github.com/generic-org/payment-gateway/pkg/domain/repository"
	app_port "github.com/generic-org/payment-gateway/pkg/application/port"
	domain_service "github.com/generic-org/payment-gateway/pkg/domain/service"
)

type paymentUseCase struct {
	paymentRepo     repository.PaymentRepository
	merchantRepo    repository.MerchantRepository
	// transactionRepo repository.TransactionRepository
	paymentProcessor app_port.PaymentProcessor
	fraudService    domain_service.FraudDetectionService
	// notificationService app_port.NotificationService
}

func NewPaymentUseCase(
	pr repository.PaymentRepository,
	mr repository.MerchantRepository,
	pp app_port.PaymentProcessor,
	fs domain_service.FraudDetectionService,
) service.PaymentService {
	return &paymentUseCase{
		paymentRepo:  pr,
		merchantRepo: mr,
		paymentProcessor: pp,
		fraudService: fs,
	}
}

// Implement service.PaymentService interface methods here (stubs for now)
func (uc *paymentUseCase) CreatePayment(ctx context.Context, input dto.CreatePaymentInput) (*dto.PaymentOutput, error) {
	// TODO: Implementation
	return nil, nil
}

func (uc *paymentUseCase) GetPayment(ctx context.Context, paymentID string) (*dto.PaymentOutput, error) {
	// TODO: Implementation
	return nil, nil
}
 
func (uc *paymentUseCase) RefundPayment(ctx context.Context, input dto.RefundPaymentInput) (*dto.RefundOutput, error) {
	// TODO: Implementation
	return nil, nil
}
