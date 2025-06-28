# Currency Exchange Matching Algorithm

This project implements a currency exchange matching system using Go, RabbitMQ, and PostgreSQL. It aims to efficiently match buyers and sellers of different currencies, support partial exchanges, and persist orders and transaction data.

## Project Goals

- Efficiently match currency buy and sell requests.
- Support partial currency exchanges.
- Persist transactions and remaining orders securely.

## Technologies Used

- **Go:** Programming language for concurrency and performance.
- **RabbitMQ:** Messaging broker for queuing and asynchronous communication.
- **PostgreSQL:** Database for reliable data persistence.
- **Docker:** Containerization for easy local and deployment environments.

## Architecture Overview

```
├── Producer
│   └── Sends buy/sell orders to RabbitMQ queues
├── Consumer
│   └── Listens to RabbitMQ queues
│       └── Matching Algorithm
│           ├── Match exact amounts
│           └── Partial matching support
└── Storage
    ├── PostgreSQL
    │   ├── Pending orders
    │   └── Completed transactions
    └── Logs
```

## Data Models

### Order

```go
type Order struct {
  ID          string
  UserID      string
  Type        string // "buy" or "sell"
  CurrencyIn  string // Currency being provided
  CurrencyOut string // Currency desired
  AmountIn    float64
  AmountOut   float64
  Timestamp   time.Time
}
```

### Matched Order

```go
type MatchedOrder struct {
  BuyOrderID  string
  SellOrderID string
  Amount      float64
  Timestamp   time.Time
}
```

## Matching Logic

- Orders matched based on currencies requested/provided.
- Partial fulfillment if full matching isn't possible.
- Remaining unmatched amounts are re-queued for future matches.

## RabbitMQ Queues

- `currency.buy`: Receives buy requests.
- `currency.sell`: Receives sell requests.
- `currency.matched`: Outputs matched transactions.

## API Endpoints (Future Implementation)

- `POST /orders`: Submit new buy/sell orders.
- `GET /orders/{id}`: Retrieve specific order details.
- `GET /transactions`: List matched transactions.

## Development Setup

### Requirements

- Docker and Docker Compose
- Go 1.22 or later

### Running Locally

```sh
# Clone repository
git clone <repo-url>
cd <repo-name>

# Start RabbitMQ and PostgreSQL
make up

# Run Go service locally
go run ./cmd/main.go
```

### Testing

```sh
# Run unit tests
go test ./...
```

## Potential Improvements

- Authentication and user management.
- Regulatory compliance considerations.
- Scalability through microservices architecture.

---

Happy coding!
