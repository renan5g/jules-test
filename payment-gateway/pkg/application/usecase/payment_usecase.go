package usecase

import (
	"context"
	"fmt" // For placeholder error/info messages
	"time" // For setting CreatedAt/UpdatedAt

	"github.com/generic-org/payment-gateway/pkg/application/dto"
	app_port "github.com/generic-org/payment-gateway/pkg/application/port" // Renamed for clarity
	"github.com/generic-org/payment-gateway/pkg/application/service"
	"github.com/generic-org/payment-gateway/pkg/domain/entity"
	"github.com/generic-org/payment-gateway/pkg/domain/repository"
	domain_service "github.com/generic-org/payment-gateway/pkg/domain/service"
)

type paymentUseCase struct {
	paymentRepo      repository.PaymentRepository
	merchantRepo     repository.MerchantRepository
	transactionRepo  repository.TransactionRepository // Added for completeness
	paymentProcessor app_port.PaymentProcessor
	fraudService     domain_service.FraudDetectionService
	logger           app_port.Logger // Added logger
	// notificationService app_port.NotificationService
}

// NewPaymentUseCase constructor updated to accept a logger and transactionRepo.
func NewPaymentUseCase(
	pr repository.PaymentRepository,
	mr repository.MerchantRepository,
	tr repository.TransactionRepository, // Added
	pp app_port.PaymentProcessor,
	fs domain_service.FraudDetectionService,
	logger app_port.Logger, // Added
) service.PaymentService {
	if logger == nil {
		fmt.Println("Warning: paymentUseCase initialized with nil logger")
	}
	return &paymentUseCase{
		paymentRepo:      pr,
		merchantRepo:     mr,
		transactionRepo:  tr,
		paymentProcessor: pp,
		fraudService:     fs,
		logger:           logger,
	}
}

func (uc *paymentUseCase) logInfo(ctx context.Context, msg string, args ...interface{}) {
	if uc.logger != nil {
		uc.logger.Info(ctx, msg, args...)
	} else {
		// Basic fallback if logger is not available
		fmt.Printf("[INFO] %s %v\n", msg, args)
	}
}

func (uc *paymentUseCase) logError(ctx context.Context, msg string, args ...interface{}) {
	if uc.logger != nil {
		uc.logger.Error(ctx, msg, args...)
	} else {
		// Basic fallback if logger is not available
		fmt.Printf("[ERROR] %s %v\n", msg, args)
	}
}

// CreatePayment implements the core logic for creating and processing a payment.
func (uc *paymentUseCase) CreatePayment(ctx context.Context, input dto.CreatePaymentInput) (*dto.PaymentOutput, error) {
	uc.logInfo(ctx, "CreatePayment started", "merchantID", input.MerchantID, "amount", input.Amount)

	// 1. Validate Merchant (simplified example)
	merchant, err := uc.merchantRepo.FindByID(ctx, input.MerchantID)
	if err != nil {
		uc.logError(ctx, "Failed to find merchant", "merchantID", input.MerchantID, "error", err)
		return nil, fmt.Errorf("merchant not found: %w", err)
	}
	if merchant == nil { // Double check after error
		uc.logError(ctx, "Merchant is nil after FindByID check", "merchantID", input.MerchantID)
		return nil, fmt.Errorf("merchant not found with ID: %s", input.MerchantID)
	}
	// TODO: Add more merchant validation (e.g., is active?)

	// 2. Create Payment Entity
	payment := &entity.Payment{
		ID:              fmt.Sprintf("pay_%v", time.Now().UnixNano()), // Simple ID generation
		Amount:          input.Amount,
		Currency:        input.Currency,
		MerchantID:      input.MerchantID,
		Status:          entity.PaymentStatusPending,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		PaymentMethodID: "pm_" + input.PaymentMethod.Type, // Simplified
		// CustomerID: input.Customer.ID, // If customer details are provided
	}
    
    // Add a pending transaction
    initialTransaction := entity.Transaction{
        ID:        fmt.Sprintf("txn_%v", time.Now().UnixNano()),
        PaymentID: payment.ID,
        Type:      entity.TransactionTypeSale, // Assuming direct sale for now
        Amount:    payment.Amount,
        Status:    entity.TransactionStatusPending,
        CreatedAt: time.Now(),
    }
    payment.Transactions = append(payment.Transactions, initialTransaction)


	// 3. Assess Fraud Risk
	uc.logInfo(ctx, "Assessing fraud risk", "paymentID", payment.ID)
	fraudAssessment, err := uc.fraudService.AssessRisk(ctx, payment)
	if err != nil {
		uc.logError(ctx, "Fraud assessment failed", "paymentID", payment.ID, "error", err)
		// Depending on policy, might proceed or fail here. For now, just log.
	} else {
		uc.logInfo(ctx, "Fraud assessment result", "paymentID", payment.ID, "riskLevel", fraudAssessment.RiskLevel, "score", fraudAssessment.Score)
		if fraudAssessment.RiskLevel == domain_service.RiskHigh {
			payment.Status = entity.PaymentStatusFailed
			// Also update transaction status
            if len(payment.Transactions) > 0 {
                payment.Transactions[0].Status = entity.TransactionStatusFailed
                payment.Transactions[0].ProviderResponse = "Blocked due to high fraud risk"
            }
			if err := uc.paymentRepo.Save(ctx, payment); err != nil { // Save the failed payment attempt
				uc.logError(ctx, "Failed to save high-risk payment", "paymentID", payment.ID, "error", err)
			}
			return nil, fmt.Errorf("payment blocked due to high fraud risk (score: %d)", fraudAssessment.Score)
		}
	}

	// 4. Save Initial Payment Entity
	uc.logInfo(ctx, "Saving initial payment", "paymentID", payment.ID)
	if err := uc.paymentRepo.Save(ctx, payment); err != nil {
		uc.logError(ctx, "Failed to save initial payment", "paymentID", payment.ID, "error", err)
		return nil, fmt.Errorf("could not save payment: %w", err)
	}

	// 5. Prepare Payment Processor Request
	// This mapping would be more complex in a real scenario
	processorRequest := app_port.ProcessorPaymentRequest{
		Amount:          int64(input.Amount * 100), // Convert to cents
		Currency:        input.Currency,
		OrderID:         payment.ID, // Use our payment ID as OrderID for processor
		Description:     input.Description,
		PaymentMethodToken: fmt.Sprintf("token_for_%s", input.PaymentMethod.Type), // Highly simplified
		// CustomerInfo: map[string]string{"email": input.Customer.Email}, // If customer exists
	}
	uc.logInfo(ctx, "Processing payment with external provider", "paymentID", payment.ID)

	// 6. Call Payment Processor
	processorResponse, err := uc.paymentProcessor.Process(ctx, processorRequest)
	if err != nil {
		uc.logError(ctx, "Payment processing failed by provider", "paymentID", payment.ID, "error", err)
		payment.Status = entity.PaymentStatusFailed
        if len(payment.Transactions) > 0 {
            payment.Transactions[0].Status = entity.TransactionStatusFailed
            payment.Transactions[0].ProviderResponse = err.Error()
        }
		// Attempt to update the payment status to FAILED in DB
		if updateErr := uc.paymentRepo.Update(ctx, payment); updateErr != nil {
			uc.logError(ctx, "Failed to update payment to FAILED status after provider error", "paymentID", payment.ID, "updateError", updateErr)
		}
		return nil, fmt.Errorf("payment provider error: %w", err)
	}
	uc.logInfo(ctx, "Payment processed by provider", "paymentID", payment.ID, "providerStatus", processorResponse.Status, "providerTxnID", processorResponse.TransactionID)

	// 7. Update Payment Entity based on Processor Response
    transaction := &payment.Transactions[0] // Get the initial transaction
	transaction.ProviderResponse = string(processorResponse.RawResponse) // Or a summary
    transaction.GatewayFee = 0.10 // Example fee
    transaction.ProviderFee = 0.25 // Example fee

	switch processorResponse.Status {
	case "succeeded":
		payment.Status = entity.PaymentStatusSucceeded
        transaction.Status = entity.TransactionStatusSucceeded
	case "pending":
		payment.Status = entity.PaymentStatusPending // Or a more specific pending status
        transaction.Status = entity.TransactionStatusPending
	case "failed":
		payment.Status = entity.PaymentStatusFailed
        transaction.Status = entity.TransactionStatusFailed
	default:
		payment.Status = entity.PaymentStatusFailed // Or some unknown/error status
        transaction.Status = entity.TransactionStatusFailed
		uc.logError(ctx, "Unhandled payment processor status", "paymentID", payment.ID, "status", processorResponse.Status)
	}
	payment.UpdatedAt = time.Now()

	// 8. Save Updated Payment Entity
	uc.logInfo(ctx, "Saving final payment state", "paymentID", payment.ID, "status", payment.Status)
	if err := uc.paymentRepo.Update(ctx, payment); err != nil { // Use Update to reflect changes
		uc.logError(ctx, "Failed to update final payment state", "paymentID", payment.ID, "error", err)
		// This is tricky: payment was processed but DB update failed. Needs reconciliation strategy.
		return nil, fmt.Errorf("could not update payment after processing: %w", err)
	}

	// 9. Map to Output DTO
	output := &dto.PaymentOutput{
		ID:         payment.ID,
		MerchantID: payment.MerchantID,
		Amount:     payment.Amount,
		Currency:   payment.Currency,
		Status:     string(payment.Status),
		OrderID:    input.OrderID, // Use original OrderID from input for merchant's reference
		CreatedAt:  payment.CreatedAt,
		UpdatedAt:  payment.UpdatedAt,
	}
	uc.logInfo(ctx, "CreatePayment completed successfully", "paymentID", payment.ID, "status", output.Status)
	return output, nil
}

func (uc *paymentUseCase) GetPayment(ctx context.Context, paymentID string) (*dto.PaymentOutput, error) {
	uc.logInfo(ctx, "GetPayment started", "paymentID", paymentID)
	// TODO: Implementation
	// 1. Fetch payment from uc.paymentRepo.FindByID(ctx, paymentID)
	// 2. Map entity.Payment to dto.PaymentOutput
	// 3. Return output or error
	return nil, fmt.Errorf("GetPayment not implemented")
}

func (uc *paymentUseCase) RefundPayment(ctx context.Context, input dto.RefundPaymentInput) (*dto.RefundOutput, error) {
	uc.logInfo(ctx, "RefundPayment started", "paymentID", input.PaymentID, "amount", input.Amount)
	// TODO: Implementation
	// 1. Fetch original payment: uc.paymentRepo.FindByID(ctx, input.PaymentID)
	// 2. Validate if refundable (e.g., status is SUCCEEDED, not already fully refunded)
	// 3. Create a new entity.Transaction for refund (status PENDING)
	// 4. Call uc.paymentProcessor.Refund(ctx, processorRefundRequest)
	// 5. Update transaction and payment status based on response.
	// 6. Save updated entities.
	// 7. Map to dto.RefundOutput.
	return nil, fmt.Errorf("RefundPayment not implemented")
}
