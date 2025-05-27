package router

// import (
//     "github.com/generic-org/payment-gateway/pkg/interfaces/api/handler"
//     "github.com/generic-org/payment-gateway/pkg/interfaces/api/middleware"
//     "github.com/gin-gonic/gin" // Or your chosen router
// )

// func SetupRouter(paymentHandler *handler.PaymentAPIHandler, merchantHandler *handler.MerchantAPIHandler) *gin.Engine {
//     r := gin.Default()
//
//     // Global middleware
//     // r.Use(middleware.Logging())
//     // r.Use(middleware.ErrorHandler())
//
//     v1 := r.Group("/v1")
//     {
//         // Authenticated group for most payment operations
//         authGroup := v1.Group("/")
//         // authGroup.Use(middleware.Authenticate( /* provide necessary dependencies */ ))
//         {
//             payments := authGroup.Group("/payments")
//             // payments.POST("", paymentHandler.CreatePayment)
//             // payments.GET("/:paymentId", paymentHandler.GetPayment)
//             // payments.POST("/:paymentId/refunds", paymentHandler.RefundPayment)
//         }
//         // Potentially merchant specific endpoints
//         // merchants := authGroup.Group("/merchants")
//         // merchants.GET("/:merchantId", merchantHandler.GetMerchant)
//     }
//     return r
// }
