# Testing Documentation

## Overview

The Task Manager API includes a comprehensive test suite covering different layers of the application:
- Domain Models
- Repositories
- Use Cases
- Controllers

## Test Structure

### Domain Tests
Location: `Domain/domain_test.go`
- Tests model validation
- Tests response structure serialization
- No external dependencies

### Repository Tests
Location: `Repositories/*_repository_test.go`
- Tests MongoDB integration
- Requires local MongoDB instance
- Tests CRUD operations
- Tests error handling

### Use Case Tests
Location: `Usecases/*_usecases_test.go`
- Tests business logic
- Uses mock repositories
- Tests error handling
- Tests timeout handling

### Controller Tests
Location: `Delivery/controllers/controller_test.go`
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
# Copy example environment file
cp .env.example .env

# Edit .env with test configuration
MONGODB_URL=mongodb://localhost:27017
DB_NAME=task_manager_test
JWT_SECRET=test-secret-key
```

### Running All Tests
```bash
# Using Makefile
make test

# Directly
go test ./... -v
```

### Running Specific Test Suites
```bash
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
# Generate coverage report
make coverage

# View coverage in browser
go tool cover -html=coverage.out
```

## Test Coverage Metrics

Current coverage status:
- Domain: 100%
- Repositories: ~95%
- Use Cases: ~90%
- Controllers: ~85%

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
# Clean test cache
make clean

# Run specific failing test
go test -run TestName ./...
```

3. Coverage Issues
```bash
# Regenerate coverage
make coverage

# Check for untested code
go tool cover -func=coverage.out
```

## Future Improvements

1. Test Coverage
   - Increase controller test coverage
   - Add integration tests
   - Add performance tests

2. Test Infrastructure
   - Add test containers for MongoDB
   - Implement test data factories
   - Add benchmark tests

3. Documentation
   - Add test case documentation
   - Document mock implementations
   - Add test coverage reports 