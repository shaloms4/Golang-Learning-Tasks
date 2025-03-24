# Task Manager API Documentation

## Base URL
```
http://localhost:8080
```

## Endpoints Overview

| Method | Endpoint                | Description                            |
|--------|-------------------------|----------------------------------------|
| GET    | /tasks                  | Get a list of all tasks.              |
| GET    | /tasks/{id}             | Get details of a specific task.       |
| POST   | /tasks                  | Create a new task.                    |
| PUT    | /tasks/{id}             | Update an existing task.              |
| DELETE | /tasks/{id}             | Remove a task by ID.                  |

---

### 1. Get All Tasks
**Request:**
```
GET /tasks
```
**Response:**
```json
[
    {
        "id": "1",
        "title": "Task 1",
        "description": "First task",
        "due_date": "2025-03-25T00:00:00Z",
        "status": "Pending"
    }
]
```

### 2. Get a Single Task
**Request:**
```
GET /tasks/{id}
```
**Response:** (If found)
```json
{
    "id": "1",
    "title": "Task 1",
    "description": "First task",
    "due_date": "2025-03-25T00:00:00Z",
    "status": "Pending"
}
```
**Response:** (If not found)
```json
{
    "error": "Task not found"
}
```

### 3. Create a New Task
**Request:**
```
POST /tasks
```
**Request Body:**
```json
{
    "id": "2",
    "title": "New Task",
    "description": "This is a new task",
    "due_date": "2025-03-25T00:00:00Z",
    "status": "Not started"
}
```
**Response:**
```json
{
    "message": "Task created"
}
```

### 4. Update an Existing Task
**Request:**
```
PUT /tasks/{id}
```
**Request Body:**
```json
{
    "title": "Updated Task Title",
    "description": "Updated description",
    "due_date": "2025-03-25T00:00:00Z",
    "status": "In Progress"
}
```
**Response:**
```json
{
    "message": "Task updated",
    "task": {
        "id": "1",
        "title": "Updated Task Title",
        "description": "Updated description",
        "due_date": "2025-03-25T00:00:00Z",
        "status": "In Progress"
    }
}
```

### 5. Delete a Task
**Request:**
```
DELETE /tasks/{id}
```
**Response:**
```json
{
    "message": "Task removed"
}
```
**Response:** (If task not found)
```json
{
    "error": "Task not found"
}
```

## Error Handling
- **400 Bad Request**: If the request body is incorrect or missing required fields.
- **404 Not Found**: If a task with the specified ID does not exist.
- **500 Internal Server Error**: If an unexpected error occurs on the server.

