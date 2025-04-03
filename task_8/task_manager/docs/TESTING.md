# Testing Documentation

## Overview

The Task Manager API includes a comprehensive test suite covering different layers of the application:
- Domain Models
- Repositories
- Use Cases
- Controllers

## Test Structure

### Domain Tests
Location: `task_manager/Domain/domain_test.go`
- Tests model validation
- Tests response structure serialization
- No external dependencies

### Repository Tests
Location: `task_manager/Repositories/*_repository_test.go`
- Tests MongoDB integration
- Requires local MongoDB instance
- Tests CRUD operations
- Tests error handling

### Use Case Tests
Location: `task_manager/Usecases/*_usecases_test.go`
- Tests business logic
- Uses mock repositories
- Tests error handling
- Tests timeout handling

### Controller Tests
Location: `task_manager/Delivery/controllers/controller_test.go`
- Tests HTTP endpoints
- Tests request/response handling
- Tests middleware integration
- Tests error responses

## Running Tests

### Prerequisites
- Go 1.21 or later
- MongoDB 4.4 or later running locally
- Environment variables set (see below)

### Environment Setup
```bash
# Navigate to task_manager directory
cd task_8/task_manager

# Copy example environment file
cp .env.example .env

# Edit .env with test configuration
MONGODB_URL=mongodb://localhost:27017
DB_NAME=task_manager_test
JWT_SECRET=test-secret-key
```

### Running All Tests
```bash
# Navigate to task_manager directory
cd task_8/task_manager

# Using Makefile
make test

# Directly
go test ./... -v
```

### Running Specific Test Suites
```bash
# Navigate to task_manager directory
cd task_8/task_manager

# Domain tests only
go test ./Domain/... -v

# Repository tests only
go test ./Repositories/... -v

# Use case tests only
go test ./Usecases/... -v

# Controller tests only
go test ./Delivery/controllers/... -v
```

### Test Coverage
```bash
# Navigate to task_manager directory
cd task_8/task_manager

# Generate coverage report
make coverage

# View coverage in browser
go tool cover -html=coverage.out
```

## Known Issues and Limitations

### Repository Tests
1. MongoDB Connection
   - Tests require a local MongoDB instance
   - Connection issues may cause test failures
   - Solution: Ensure MongoDB is running before tests

2. Database Cleanup
   - Tests drop the database after completion
   - Interrupted tests may leave test data
   - Solution: Manual cleanup may be required

### Controller Tests
1. Mock Setup
   - Some tests fail due to incorrect mock expectations
   - Solution: Review and update mock implementations

2. Authentication
   - JWT token validation tests need improvement
   - Solution: Add more test cases for token validation

## Best Practices

1. Writing New Tests
   - Follow existing test patterns
   - Use descriptive test names
   - Include both success and failure cases
   - Mock external dependencies

2. Test Data
   - Use consistent test data
   - Clean up after tests
   - Use unique identifiers

3. Error Handling
   - Test error cases explicitly
   - Verify error messages
   - Test error status codes

## CI/CD Integration

Tests are automatically run in the CI pipeline:
- On every push to main
- On pull requests
- Using GitHub Actions

### CI Environment
- Ubuntu latest
- Go 1.21
- MongoDB service
- Environment variables set

## Troubleshooting

### Common Issues

1. MongoDB Connection
```bash
# Check MongoDB status
mongosh --eval "db.serverStatus()"

# Restart MongoDB service
sudo service mongod restart
```

2. Test Failures
```bash
# Navigate to task_manager directory
cd task_8/task_manager

# Clean test cache
make clean

# Run specific failing test
go test -run TestName ./...
```