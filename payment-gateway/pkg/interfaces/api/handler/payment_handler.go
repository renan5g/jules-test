package handler

import (
	"github.com/generic-org/payment-gateway/pkg/application/service" 
	// "net/http" // Uncomment if using net/http directly
	// "github.com/gin-gonic/gin" // Or your chosen router
)

type PaymentAPIHandler struct {
	paymentService service.PaymentService
}

func NewPaymentAPIHandler(ps service.PaymentService) *PaymentAPIHandler {
	return &PaymentAPIHandler{paymentService: ps}
}

// Example stubs, assuming a framework like Gin:
// func (h *PaymentAPIHandler) CreatePayment(c *gin.Context) { /* TODO: implement */ }
// func (h *PaymentAPIHandler) GetPayment(c *gin.Context) { /* TODO: implement */ }
// func (h *PaymentAPIHandler) RefundPayment(c *gin.Context) { /* TODO: implement */ }
