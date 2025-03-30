# Task Management System With Clean Architecture

## Overview

This Task Management API follows **Clean Architecture** principles, ensuring separation of concerns and maintainability. It uses **MongoDB** for data persistence, **JWT** for authentication, and **bcrypt** for password hashing. The API provides user registration, login, task CRUD operations, and user role management.

---

## Key Concepts: **Clean Architecture**

The project structure adheres to **Clean Architecture**, which separates the application into distinct layers:

- **Domain Layer**: Contains core business logic and entities.
- **Usecase Layer**: Implements business rules for tasks and users.
- **Repository Layer**: Handles database interactions (MongoDB).
- **Infrastructure Layer**: Deals with external systems like JWT and bcrypt.
- **Controller Layer**: Contains HTTP handlers that interact with the usecases.
- **Router Layer**: Manages the routing of requests to the appropriate controllers.

## Example Usage

### 1. **User Registration**  
- **Endpoint**: `POST /register`
- **Description**: Registers a new user with a hashed password.

#### Request Example:
```bash
POST /register
Content-Type: application/json

{
  "username": "john_doe",
  "password": "securepassword"
}
```

#### Response Example:
```json
{
  "message": "User registered successfully",
}
```

---

### 2. **User Login**  
- **Endpoint**: `POST /login`
- **Description**: Logs in a user and returns a JWT token.

#### Request Example:
```bash
POST /login
Content-Type: application/json

{
  "username": "john_doe",
  "password": "securepassword"
}
```

#### Response Example:
```json
{
  "token": "your_jwt_token"
}
```

---

### 3. **Create a Task**  
- **Endpoint**: `POST /tasks`
- **Description**: Creates a new task.

#### Request Example:
```bash
POST /tasks
Authorization: Bearer your_jwt_token_here
Content-Type: application/json

{
  "title": "Complete Project",
  "description": "Finish the task management API",
  "status": "in-progress",
  "due_date": "2025-04-30T00:00:00Z"
}
```

#### Response Example:
```json
{
  "message": "Task created successfully",
}
```

---

### 4. **Fetch All Tasks**  
- **Endpoint**: `GET /tasks`
- **Description**: Retrieves all tasks.

#### Request Example:
```bash
GET /tasks
Authorization: Bearer your_jwt_token_here
```

#### Response Example:
```json
[
  {
    "id": "60dbf0c91a789d1bc8e5862f",
    "title": "Complete Project",
    "description": "Finish the task management API",
    "status": "in-progress",
    "due_date": "2025-04-30T00:00:00Z"
  },
  {
    "id": "60dbf0c91a789d1bc8e58630",
    "title": "Write Documentation",
    "description": "Document the API for users",
    "status": "not-started",
    "due_date": "2025-04-10T00:00:00Z"
  }
]
```

---

### 5. **Update a Task**  
- **Endpoint**: `PUT /tasks/{id}`
- **Description**: Updates a task by ID.

#### Request Example:
```bash
PUT /tasks/60dbf0c91a789d1bc8e5862f
Authorization: Bearer your_jwt_token_here
Content-Type: application/json

{
  "status": "completed"
}
```

#### Response Example:
```json
{
  "message": "Task updated successfully"
}
```

---

### 6. **Delete a Task**  
- **Endpoint**: `DELETE /tasks/{id}`
- **Description**: Deletes a task by ID.

#### Request Example:
```bash
DELETE /tasks/60dbf0c91a789d1bc8e5862f
Authorization: Bearer your_jwt_token_here
```

#### Response Example:
```json
{
  "message": "Task deleted successfully"
}
```

---

## Testing the API

1. **Create a User**: Use the `POST /register` endpoint to add a user.
2. **Login**: Use the `POST /login` endpoint to authenticate and receive a JWT token.
3. **Create Tasks**: Use the `POST /tasks` endpoint to add tasks.
4. **Get Tasks**: Use the `GET /tasks` or `GET /tasks/{id}` endpoint to retrieve tasks.
5. **Update/Delete Tasks**: Use the `PUT /tasks/{id}` and `DELETE /tasks/{id}` endpoints to modify or remove tasks.
