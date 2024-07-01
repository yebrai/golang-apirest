# Simple Go REST API

This is a simple REST API built with Go. It provides basic CRUD operations for managing Todo items.

## Getting Started

### Prerequisites

- Go 1.16 or higher
- cURL or any API client (e.g., Postman)
- PowerShell (for Windows users)

### Project Structure

├── main.go
├── handlers
│ └── todo.go
└── go.mod

bash
Copiar código

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yebrai/golang-apirest.git
    cd golang-apirest
    ```

2. Build and run the application:

    ```sh
    go run main.go
    ```

### API Endpoints

#### Create Todo

- **URL**: `/todo/create`
- **Method**: `POST`
- **Body**:

    ```json
    {
        "title": "Learn Go",
        "status": "Pending"
    }
    ```

- **Example with cURL**:

    ```sh
    curl -X POST http://localhost:8080/todo/create -d '{"title":"Learn Go","status":"Pending"}' -H "Content-Type: application/json"
    ```

- **Example with PowerShell**:

    ```powershell
    $headers = @{
        "Content-Type" = "application/json"
    }
    $body = @{
        "title" = "Learn Go"
        "status" = "Pending"
    } | ConvertTo-Json
    Invoke-RestMethod -Uri http://localhost:8080/todo/create -Method Post -Body $body -Headers $headers
    ```

#### Edit Todo

- **URL**: `/todo/edit`
- **Method**: `PUT`
- **Body**:

    ```json
    {
        "id": "1",
        "title": "Learn Go",
        "status": "Completed"
    }
    ```

- **Example with cURL**:

    ```sh
    curl -X PUT http://localhost:8080/todo/edit -d '{"id":"1","title":"Learn Go","status":"Completed"}' -H "Content-Type: application/json"
    ```

- **Example with PowerShell**:

    ```powershell
    $headers = @{
        "Content-Type" = "application/json"
    }
    $body = @{
        "id" = "1"
        "title" = "Learn Go"
        "status" = "Completed"
    } | ConvertTo-Json
    Invoke-RestMethod -Uri http://localhost:8080/todo/edit -Method Put -Body $body -Headers $headers
    ```
