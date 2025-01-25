# Receipt Processor

A simple web service that processes receipts and calculates reward points based on predefined rules. The service is built using Go (`gin` framework) and uses in-memory storage to manage receipts.

---

## Features
- **Process Receipt**: Submit a receipt, validate its content, and store it with a generated unique ID.
- **Get Points**: Retrieve the reward points for a receipt using its unique ID.
- **In-Memory Storage**: Data is stored temporarily and does not persist between application restarts.

---

## API Endpoints

### 1. Process Receipt
- **Path**: `/receipts/process`
- **Method**: `POST`
- **Description**: Submits a receipt for processing, validates the input, and returns a unique ID for the receipt.
- **Request Body**:
    ```json
    {
      "retailer": "string",
      "purchaseDate": "YYYY-MM-DD",
      "purchaseTime": "HH:mm",
      "total": "string",
      "items": [
        {
          "shortDescription": "string",
          "price": "string"
        }
      ]
    }
    ```
- **Response**:
  - **Success (200)**:
    ```json
    { "id": "RECabc123XYZ" }
    ```
  - **Error (400)**:
    ```json
    {
      "error": "The receipt is invalid.",
      "details": "One or more required fields are missing."
    }
    ```

---

### 2. Get Points
- **Path**: `/receipts/{id}/points`
- **Method**: `GET`
- **Description**: Retrieves the reward points for a specific receipt using its unique ID.
- **Path Parameters**:
  - `id`: The unique receipt ID (e.g., `RECabc123XYZ`).
- **Response**:
  - **Success (200)**:
    ```json
    { "points": 100 }
    ```
  - **Error (404)**:
    ```json
    {
      "error": "No receipt found for that ID.",
      "details": "The requested receipt ID does not exist."
    }
    ```

---

## Points Calculation Rules
1. **Retailer Name**:
   - 1 point for each alphanumeric character in the retailer name.
2. **Round Dollar Amount**:
   - 50 points if the total is a round dollar amount (e.g., `$35.00`).
3. **Multiples of 0.25**:
   - 25 points if the total is a multiple of `$0.25`.
4. **Items**:
   - 5 points for every two items on the receipt.
5. **Item Description**:
   - If the trimmed length of the item description is a multiple of 3, multiply the price by `0.2` and round up to the nearest integer. Add the result as points.
6. **Purchase Date**:
   - 6 points if the day of the purchase is odd.
7. **Purchase Time**:
   - 10 points if the purchase time is between `2:00 PM` and `4:00 PM`.

---

## Setup Instructions

### Requirements

- Go 1.19 or later

### Steps to Setup and Run

1. **Clone the repository**:
    ```bash
    git clone https://github.com/your-username/receipt-processor.git
    cd receipt-processor
    ```

2. **Install dependencies**:
    ```bash
    go mod tidy
    ```

3. **Run the application**:
    ```bash
    go run cmd/server.go
    ```

4. **Test the endpoints**:
    Use Postman or curl to access the endpoints.

---

## Running Tests and Generating Coverage

### Run Unit Tests

Run all test cases for the application:
```bash
go test ./internal/... ./utils/...
```bash

###Generate coverage Report

1. Generate a coverage file
```bash
go test ./... -coverprofile=coverage.out
```bash

2. Generate an HTML report for a detailed visual report
```bash
go tool cover -html=coverage.out
```bash

---

## Project Structure

```plaintext
.
├── README.md
├── cmd
│   └── server.go           # Entry point of the application
├── go.mod                  # Go module definition
├── go.sum                  # Dependency management
├── internal
│   ├── handlers            # HTTP handler functions
│   │   ├── points.go       # Get points handler
│   │   ├── points_test.go  # Unit tests for points handler
│   │   ├── process.go      # Process receipt handler
│   │   └── process_test.go # Unit tests for process handler
│   ├── models              # Data models (Receipt, Item)
│   │   └── receipt.go
│   ├── routes              # API routes
│   │   └── receiptRoutes.go
│   ├── services            # Business logic for receipt validation and points calculation
│   │   └── receiptService.go
│   └── storage             # In-memory storage implementation
│       └── storage.go
└── utils                   # Utility functions
    ├── errorHandler.go     # Error response utilities
    ├── idGenerator.go      # Unique ID generator
    ├── idGenerator_test.go # Tests for ID generator
    ├── pointRules.go       # Logic for calculating points
    └── pointRules_test.go  # Tests for point calculation logic

---

##Future Improvements

1. Add persistent storage (e.g., a database) to store receipts permanently.
2. Enhance input validation with stricter rules.
4. Implement authentication and authorization for API endpoints, that is the user who has uploaded a receipt should only be able to get the points for that receipt.





