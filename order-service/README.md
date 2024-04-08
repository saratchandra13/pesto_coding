# Order Service

This is a basic order service implemented in Go. It uses the Gin framework for the web server and Gorm for database interactions.

## Prerequisites

- Go 1.16 or later
- MySQL

## Running the Service

1. Clone the repository:
   git clone https://github.com/lazycoder1995/order_service.git

2. Navigate to the project directory:
    cd order_service

3. Run the service:
    go run main.go
4. The service will start and listen on port 8080.

## API Endpoints

### Create Order

Creates a new order.

**Request:**

```bash
curl -X POST -H "Content-Type: application/json" -d '{"userID": 1, "productIDs": [1, 2, 3]}' http://localhost:8080/orders
```

### Change order status

Changes the status of an order.

**Request:**

```bash
curl -X PUT -H "Content-Type: application/json" -d '{"status": "shipped"}' http://localhost:8080/orders/1/status
```
