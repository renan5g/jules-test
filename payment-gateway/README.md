# Payment Gateway API

## Overview

This project is a scalable and robust Payment Gateway API built with Go. It aims to provide a clean, secure, and efficient way to process payments, manage merchants, and integrate with various payment providers.

## Features (Planned)

*   **Multiple Payment Methods:** Support for Credit Cards, Boleto, Pix, and potentially others.
*   **Merchant Management:** Onboarding and management of merchants.
*   **Payment Processing:** Authorization, capture, sale, refund, and void operations.
*   **Transaction Management:** Tracking of all payment and transaction statuses.
*   **Customer Management:** Optional storage of customer details for faster checkouts.
*   **Payment Method Tokenization:** Securely store and reuse payment methods.
*   **Fraud Detection:** Basic integration point for fraud detection services.
*   **Webhook Notifications:** Inform merchants about payment status changes.
*   **Provider Agnostic:** Designed to easily integrate with multiple payment providers (e.g., Stripe, PayPal).
*   **RESTful API:** Clean and well-documented API endpoints.
*   **Security:** Focus on secure coding practices, PCI DSS compliance considerations (though full compliance is out of scope for this example).
*   **Observability:** Structured logging (via `log/slog` adapter), placeholder for distributed tracing (OpenTelemetry), and planned metrics.
*   **Configuration Management:** Flexible configuration loading.

## Project Structure (Hexagonal Architecture)

The project follows a Hexagonal Architecture (Ports and Adapters) pattern to ensure a separation of concerns and maintainability.

```
payment-gateway/
├── api/                             # OpenAPI specifications
├── cmd/
│   └── api/                         # Main application entrypoint
│       └── main.go
├── configs/                         # Configuration files (e.g., config.yaml.example)
├── pkg/
│   ├── application/                 # Application core
│   │   ├── dto/                     # Data Transfer Objects for use cases
│   │   ├── port/                    # Ports (interfaces) for outbound services (e.g., Payment Processor, Notification)
│   │   ├── service/                 # Application services (interfaces, e.g., PaymentService)
│   │   └── usecase/                 # Use case implementations
│   ├── domain/                      # Domain core - business logic and entities
│   │   ├── entity/                  # Domain entities (e.g., Payment, Merchant)
│   │   ├── repository/              # Domain repository interfaces
│   │   └── service/                 # Domain services (e.g., FraudDetectionService)
│   ├── infrastructure/              # Adapters for external concerns
│   │   ├── config/                  # Configuration loading implementation
│   │   ├── logging/                 # Logging setup (e.g., slog adapter)
│   │   ├── tracing/                 # Tracing setup (e.g., OpenTelemetry placeholder)
│   │   ├── paymentprovider/         # Adapters for specific payment providers (Stripe, PayPal)
│   │   │   ├── stripe/
│   │   │   └── paypal/
│   │   ├── persistence/             # Database adapters (e.g., PostgreSQL)
│   │   │   └── postgres/
│   │   ├── notification/            # Notification service adapters (e.g., email, SMS)
│   └── interfaces/                  # Adapters for inbound requests
│       └── api/                     # HTTP API layer
│           ├── dto/                 # Data Transfer Objects for API requests/responses
│           ├── handler/             # HTTP request handlers
│           ├── middleware/          # HTTP middleware (auth, logging, etc.)
│           └── router/              # API router setup
├── scripts/                         # Utility scripts (build, test, deploy)
├── test/                            # End-to-end and integration tests
├── go.mod
├── go.sum
└── README.md
```

## Code Examples / Implemented Stubs

To better illustrate the interaction between components and the general flow of data, the following key methods have been implemented with more detailed (though still stubbed) logic:

*   **`pkg/application/usecase/payment_usecase.go`**:
    *   The `CreatePayment` method now demonstrates the orchestration of calls to the merchant repository, fraud detection service, payment repository (save initial state, update final state), and the payment processor. It includes example logging throughout the process.
*   **`pkg/infrastructure/persistence/postgres/payment_postgres_repository.go`**:
    *   The `Save` method includes an example of how a repository would interact with a logger and contains placeholders for actual database operations.

These examples should provide a clearer picture of how the different layers and components are intended to work together.

## Getting Started

### Prerequisites

*   Go (version 1.20 or higher recommended)
*   Docker & Docker Compose (for database and other services)
*   Make (optional, for using Makefile commands)

### Installation & Running

1.  **Clone the repository:**
    ```bash
    git clone https://github.com/generic-org/payment-gateway.git
    cd payment-gateway
    ```

2.  **Configuration:**
    *   Copy `configs/config.example.yaml` to `configs/config.yaml`.
    *   Update `configs/config.yaml` with your database credentials, API keys, etc.

3.  **Build the application:**
    ```bash
    go build -o ./bin/api ./cmd/api/main.go
    ```
    (Or use `make build` if a Makefile is provided)

4.  **Run database migrations (if applicable):**
    *   This step will depend on the migration tool chosen (e.g., `golang-migrate/migrate`).
    ```bash
    # Example: migrate -database "postgres://user:pass@host:port/dbname?sslmode=disable" -path db/migrations up
    ```

5.  **Run the application:**
    ```bash
    ./bin/api
    ```
    (Or use `make run` if a Makefile is provided)

    The API should now be running on the configured port (e.g., `http://localhost:8080`).

### Running with Docker (Recommended for Development)

1.  **Ensure Docker and Docker Compose are installed.**
2.  **Build and run the services:**
    ```bash
    docker-compose up --build
    ```
    This will typically start the API server and a PostgreSQL database.

## API Documentation

API documentation will be available via Swagger/OpenAPI once implemented.
*   Access at `/swagger/index.html` (or similar path).

## Testing

*   **Unit Tests:**
    ```bash
    go test ./...
    ```
    (Or `make test-unit`)
*   **Integration Tests:**
    *   Ensure dependent services (like database) are running.
    ```bash
    go test -tags=integration ./...
    ```
    (Or `make test-integration`)

## Dependencies

*   **Gin-Gonic (or other router):** For HTTP routing.
*   **Viper:** For configuration management.
*   **SQLx (or GORM/pgx):** For database interaction.
*   **Testify:** For assertions in tests.
*   **GoMock:** For generating mocks.
*   **Shopspring/Decimal:** For precise monetary calculations.
*   **golang-migrate/migrate:** For database migrations.
*   **Validator v10:** For request validation.

(These will be added to `go.mod` as the project develops.)

## Contributing

Contributions are welcome! Please follow these steps:
1.  Fork the repository.
2.  Create a new branch (`git checkout -b feature/your-feature-name`).
3.  Make your changes.
4.  Write tests for your changes.
5.  Ensure all tests pass (`go test ./...`).
6.  Lint your code (`golangci-lint run` - if configured).
7.  Commit your changes (`git commit -am 'Add some feature'`).
8.  Push to the branch (`git push origin feature/your-feature-name`).
9.  Create a new Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details (To be added).

## 6. API Endpoint Definitions (Conceptual)

This section outlines the primary API endpoints for the payment gateway.

*   **Authentication:** All API requests are expected to be authenticated via an API Key passed in the `X-API-Key` header. The API key identifies the merchant.

*   **Error Responses:** Errors are returned in a consistent JSON format:
    ```json
    {
      "code": "internal_error_code", // e.g., "validation_error", "payment_failed"
      "message": "A human-readable error message",
      "details": [
        {
          "resource": "Payment", // Optional: The resource related to the error
          "field": "amount",     // Optional: The specific field that caused the error
          "code": "invalid_format",
          "message": "Amount must be a positive number."
        }
      ]
    }
    ```

---

### Payments

#### `POST /v1/payments`

*   **Description:** Creates and initiates a new payment.
*   **Request Body:** (`application/json`) - See `pkg/interfaces/api/dto/CreatePaymentRequest`
    ```json
    {
      "amount": "100.50", // Decimal as string or number, handled by shopspring/decimal
      "currency": "USD",   // ISO 4217
      "merchantId": "merch_xxxx", // Usually derived from API Key, but can be explicit
      "paymentMethod": {
        "type": "credit_card", // "credit_card", "pix", "boleto"
        "creditCard": {
          "cardNumber": "4111111111111111",
          "expiryMonth": "12",
          "expiryYear": "2025",
          "cvv": "123",
          "holderName": "John Doe"
        }
        // ... other payment method details for pix, boleto etc.
      },
      "customer": { // Optional
        "id": "cust_xxxx",
        "name": "Jane Doe",
        "email": "jane.doe@example.com",
        "document": "12345678900" // e.g. CPF for Brazil
      },
      "orderId": "ORD-MERCH-12345", // Merchant's internal order ID
      "description": "Payment for Order ORD-MERCH-12345"
    }
    ```
*   **Success Response:** (`201 Created` or `202 Accepted` if processing is asynchronous) - See `pkg/interfaces/api/dto/PaymentResponse`
    ```json
    {
      "id": "pay_xxxxxxxxxxxx",
      "status": "PENDING", // Or "AUTHORIZED", "SUCCEEDED" depending on flow
      "amount": "100.50",
      "currency": "USD",
      "merchantId": "merch_xxxx",
      "orderId": "ORD-MERCH-12345",
      "createdAt": "2023-10-27T10:00:00Z",
      "updatedAt": "2023-10-27T10:00:00Z"
      // ... other relevant details
    }
    ```
*   **Possible Statuses in Response:** `PENDING`, `AUTHORIZED` (for two-step auth/capture), `SUCCEEDED` (for direct sale or if provider confirms immediately), `FAILED`.

---

#### `GET /v1/payments/{paymentId}`

*   **Description:** Retrieves the details and status of a specific payment.
*   **Path Parameters:**
    *   `paymentId` (string, required): The unique identifier of the payment.
*   **Success Response:** (`200 OK`) - See `pkg/interfaces/api/dto/PaymentResponse`
    ```json
    {
      "id": "pay_xxxxxxxxxxxx",
      "status": "SUCCEEDED",
      "amount": "100.50",
      "currency": "USD",
      // ... all other fields from PaymentResponse
      "transactions": [ // Optional: could include transaction history
        {
          "id": "txn_yyyyyyyyyyyy",
          "type": "SALE", // Or "AUTHORIZATION", "CAPTURE"
          "status": "SUCCEEDED",
          "amount": "100.50",
          "providerResponse": "Approved", // Or some code
          "createdAt": "2023-10-27T10:00:05Z"
        }
      ]
    }
    ```
*   **Error Responses:** `404 Not Found` if payment ID is invalid.

---

### Refunds

#### `POST /v1/payments/{paymentId}/refunds`

*   **Description:** Initiates a full or partial refund for a previously successful payment.
*   **Path Parameters:**
    *   `paymentId` (string, required): The unique identifier of the payment to be refunded.
*   **Request Body:** (`application/json`) - See `pkg/interfaces/api/dto/RefundPaymentRequest` (API DTO for refund)
    ```json
    {
      "amount": "50.00", // Amount to refund. If omitted, full refund.
      "reason": "Customer request" // Optional
    }
    ```
*   **Success Response:** (`201 Created` or `202 Accepted`) - See `pkg/interfaces/api/dto/RefundResponse`
    ```json
    {
      "id": "ref_zzzzzzzzzzzz", // Refund transaction ID
      "paymentId": "pay_xxxxxxxxxxxx",
      "status": "PENDING", // Or "SUCCEEDED" if processed synchronously
      "amountRefunded": "50.00",
      "currency": "USD",
      "createdAt": "2023-10-27T11:00:00Z"
    }
    ```
*   **Error Responses:** `400 Bad Request` (e.g., payment not refundable, invalid amount), `404 Not Found`.

---

#### `GET /v1/refunds/{refundId}`

*   **Description:** Retrieves the details and status of a specific refund.
*   **Path Parameters:**
    *   `refundId` (string, required): The unique identifier of the refund.
*   **Success Response:** (`200 OK`) - See `pkg/interfaces/api/dto/RefundResponse`
    ```json
    {
      "id": "ref_zzzzzzzzzzzz",
      "paymentId": "pay_xxxxxxxxxxxx",
      "status": "SUCCEEDED",
      "amountRefunded": "50.00",
      "currency": "USD",
      "createdAt": "2023-10-27T11:00:00Z"
      // ... other relevant details
    }
    ```
*   **Error Responses:** `404 Not Found`.

## 7. Detailed Payment Flow Example (Credit Card Sale)

This describes the sequence of events for a typical credit card payment that is processed as a "Sale" (Authorize and Capture in one step).

1.  **Client Request:** Merchant's system sends a `POST /v1/payments` request to the Payment Gateway's API.
    *   Payload includes amount, currency, credit card details, customer info, etc.
    *   `X-API-Key` header is present for authentication.

2.  **API Layer (`pkg/interfaces/api`)**:
    *   **Router (`router.go`):** Directs the request to the appropriate handler (e.g., `PaymentAPIHandler.CreatePayment`).
    *   **Middleware (`middleware/auth.go`):** (If configured for this route or globally) Authenticates the request using `X-API-Key`. Fetches merchant details from `MerchantRepository` (via an application service or directly if simpler for auth). If invalid, returns `401 Unauthorized`.
    *   **Handler (`handler/payment_handler.go` - `CreatePayment` method):**
        *   Binds the JSON request body to `interfaces/api/dto.CreatePaymentRequest`.
        *   Performs validation on the request DTO (e.g., using `validator` tags). If invalid, returns `400 Bad Request` with `ErrorResponse`.
        *   Maps `interfaces/api/dto.CreatePaymentRequest` to `application/dto.CreatePaymentInput`.

3.  **Application Layer (`pkg/application`)**:
    *   **Payment Use Case (`usecase/payment_usecase.go` - `CreatePayment` method):**
        *   Receives `application/dto.CreatePaymentInput`.
        *   Validates merchant existence and status using `MerchantRepository.FindByID` (or `FindByAPIKey` if API key was passed down).
        *   Creates a new `domain/entity.Payment` object with an initial status (e.g., `PENDING`).
        *   Creates a `domain/entity.Transaction` object associated with the payment, also `PENDING`.
        *   Calls `FraudDetectionService.AssessRisk` with the payment details. If high risk and configured to block, it might update payment status to `FAILED` and return.
        *   Persists the initial `Payment` and `Transaction` entities using `PaymentRepository.Save` (which might also save associated transactions within the same DB transaction).
        *   Prepares a `application/port.ProcessorPaymentRequest` DTO. This involves mapping data from `CreatePaymentInput` and the `Payment` entity to what the `PaymentProcessor` interface expects. Sensitive details like full card number are only held briefly and sent to the processor.
        *   Calls `PaymentProcessor.Process(ctx, processorRequest)`. The `paymentProcessor` is an interface instance (e.g., `StripeAdapter`, `PayPalAdapter`) injected into the use case.

4.  **Infrastructure Layer (`pkg/infrastructure`)**:
    *   **Payment Provider Adapter (e.g., `paymentprovider/stripe/stripe_adapter.go` - `Process` method):**
        *   Receives `application/port.ProcessorPaymentRequest`.
        *   Maps this request to the specific format required by the external payment provider's SDK/API (e.g., Stripe API call).
        *   Makes the actual HTTP call to the external payment provider (e.g., Stripe).
        *   Receives the response from the provider.
        *   Maps the provider's response back to `application/port.ProcessorPaymentResponse` (including status, transaction ID from provider, any errors).

5.  **Application Layer (Return Path)**:
    *   **Payment Use Case (`usecase/payment_usecase.go`):**
        *   Receives `ProcessorPaymentResponse` from the adapter.
        *   Updates the `Payment` and `Transaction` entities based on the processor's response (e.g., status to `SUCCEEDED` or `FAILED`, provider transaction ID, etc.).
        *   Saves the updated `Payment` and `Transaction` entities using `PaymentRepository.Update`.
        *   If payment was successful, may trigger a notification via `NotificationService.SendNotification` (e.g., email receipt to customer).
        *   Maps the updated `Payment` entity (and relevant transaction info) to `application/dto.PaymentOutput`.
        *   Returns `PaymentOutput` (or an error) to the API Handler.

6.  **API Layer (Return Path)**:
    *   **Handler (`handler/payment_handler.go`):**
        *   Receives `application/dto.PaymentOutput` (or error).
        *   If an error occurred, maps it to `interfaces/api/dto.APIErrorResponse` and returns the appropriate HTTP status code (e.g., 400, 500).
        *   If successful, maps `application/dto.PaymentOutput` to `interfaces/api/dto.PaymentResponse`.
        *   Sends the JSON response back to the client with HTTP status `201 Created` or `200 OK`.

7.  **Client:** Receives the HTTP response.
