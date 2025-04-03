# Task Manager API

A Go-based task management API with user authentication and role-based access control.

## Features

- Task CRUD operations
- User authentication (register/login)
- Role-based access control (admin/user)
- MongoDB persistence
- JWT-based authentication

## Prerequisites

- Go 1.21 or later
- MongoDB 4.4 or later
- Make (optional, for using Makefile commands)

## Getting Started

1. Clone the repository:
```bash
git clone https://github.com/yourusername/task-manager.git
cd task-manager
```

2. Install dependencies:
```bash
make deps
```

3. Set up environment variables:
```bash
cp .env.example .env
# Edit .env with your configuration
```

4. Run the application:
```bash
go run Delivery/main.go
```

## Testing

### Running Tests

To run all tests:
```bash
make test
```

To run tests with coverage:
```bash
make coverage
```

To run the linter:
```bash
make lint
```

### Test Environment

Tests require a MongoDB instance running locally. The test database is automatically created and dropped during test execution.

## CI/CD

The project uses GitHub Actions for continuous integration. The workflow:
- Runs on every push to main and pull requests
- Executes all tests
- Runs the linter
- Requires all checks to pass before merging

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Run tests and linter
5. Submit a pull request

## License

MIT
