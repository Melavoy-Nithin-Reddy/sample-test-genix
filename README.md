# Sample User Service

A simple Go microservice for managing users, designed to demonstrate test-genix capabilities for automated test generation.

## Overview

This microservice provides a RESTful API for user management with the following features:
- Create, read, update, and delete users
- User validation and business logic
- In-memory data storage
- Health check endpoint
- User statistics

## Architecture

The service follows a clean architecture pattern with clear separation of concerns:

```
sample-user-service/
├── main.go                    # Application entry point
├── models/                    # Data models and DTOs
│   └── user.go
├── repository/                # Data access layer
│   └── user_repository.go
├── service/                   # Business logic layer
│   └── user_service.go
└── handlers/                  # HTTP handlers
    └── user_handler.go
```

## API Endpoints

### Health Check
- `GET /health` - Service health status

### User Management
- `POST /api/v1/users` - Create a new user
- `GET /api/v1/users` - Get all users
- `GET /api/v1/users/{id}` - Get user by ID
- `PUT /api/v1/users/{id}` - Update user
- `DELETE /api/v1/users/{id}` - Delete user
- `GET /api/v1/users/stats` - Get user statistics

## Getting Started

### Prerequisites
- Go 1.21 or higher

### Installation

1. Clone or navigate to the service directory:
```bash
cd sample-user-service
```

2. Install dependencies:
```bash
go mod download
```

3. Run the service:
```bash
go run main.go
```

The service will start on port 8080 by default. You can change this by setting the `PORT` environment variable:
```bash
PORT=3000 go run main.go
```

## Example Usage

### Create a User
```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "username": "johndoe",
    "email": "john@example.com",
    "full_name": "John Doe"
  }'
```

### Get All Users
```bash
curl http://localhost:8080/api/v1/users
```

### Get User by ID
```bash
curl http://localhost:8080/api/v1/users/user_1234567
```

### Update User
```bash
curl -X PUT http://localhost:8080/api/v1/users/user_1234567 \
  -H "Content-Type: application/json" \
  -d '{
    "email": "newemail@example.com",
    "full_name": "John Updated Doe"
  }'
```

### Delete User
```bash
curl -X DELETE http://localhost:8080/api/v1/users/user_1234567
```

### Get User Statistics
```bash
curl http://localhost:8080/api/v1/users/stats
```

## Using with test-genix

This microservice is designed to work with test-genix for automated test generation. Here's how to use it:

### 1. Make Changes to the Service
For example, add a new feature or modify existing functionality:
```bash
# Edit any file in the service
vim service/user_service.go
```

### 2. Generate Tests with test-genix

Navigate to the test-genix directory and run:
```bash
cd ../test-genix

# Generate tests for the sample-user-service
go run cmd/orchestrator/main.go \
  --config config/test-orchestrator-config.yaml \
  --target ../sample-user-service
```

### 3. Review Generated Tests
test-genix will analyze your changes and generate appropriate tests in the service directory.

## Testing

Run the generated tests:
```bash
go test ./...
```

Run tests with coverage:
```bash
go test -cover ./...
```

Run tests with verbose output:
```bash
go test -v ./...
```

## Features for Test Generation

This service includes various patterns that test-genix can detect and generate tests for:

1. **Input Validation** - Models have validation methods
2. **Error Handling** - Repository and service layers return specific errors
3. **Business Logic** - Service layer contains business rules
4. **CRUD Operations** - Complete create, read, update, delete functionality
5. **Concurrency** - Repository uses mutex for thread-safe operations
6. **HTTP Handlers** - REST API endpoints with proper status codes
7. **Statistics** - Aggregation logic for user stats

## Development

### Adding New Features

When adding new features, test-genix can automatically generate tests:

1. Add your new functionality
2. Commit your changes
3. Run test-genix to generate tests
4. Review and refine the generated tests

### Example: Adding a Search Feature

```go
// In service/user_service.go
func (s *UserService) SearchUsers(query string) ([]*models.User, error) {
    // Implementation
}
```

test-genix will detect this new method and generate appropriate tests covering:
- Happy path scenarios
- Edge cases (empty query, no results)
- Error conditions

## License

This is a sample project for demonstration purposes.