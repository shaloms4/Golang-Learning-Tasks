# Task Manager API Documentation

## MongoDB Setup

1. **MongoDB URI Configuration**:  
   Set up your MongoDB URI in a `.env` file at the root of your project:
   ```bash
   MONGO_URI=mongodb+srv://your-username:your-password@cluster0.wmhc1.mongodb.net/taskmanager?retryWrites=true&w=majority
   ```

2. **Install MongoDB Driver**:  
   Ensure the MongoDB Go driver is installed:
   ```bash
   go get go.mongodb.org/mongo-driver/mongo
   ```

3. **MongoDB Setup Documentation**:  
   For more information on setting up MongoDB, visit the official MongoDB setup documentation:  
   [MongoDB Setup Guide](https://www.mongodb.com/docs/guides/)

4. **Load Environment Variables**:  
   The `.env` file is loaded using `godotenv` to retrieve the `MONGO_URI`.

---

## Base URL

```
http://localhost:8080
```

---

## How to Try It

### Step 1: Clone the Repository

Clone the repository to your local machine.

```bash
git clone https://github.com/shaloms4/Golang-Learning-Tasks
```

### Step 2: Install Dependencies

Navigate to the project directory and install the necessary dependencies.

```bash
cd Golang-Learning-Tasks/task_5
go mod tidy
```

### Step 3: Set Up MongoDB

1. **Create a MongoDB Cluster**:  
   If you don’t have MongoDB set up yet, you can create a free MongoDB cluster by following this guide:
   [Create a Free MongoDB Cluster](https://www.mongodb.com/cloud/atlas/register)

2. **Create a `.env` file** in the root of your project.
3. **Add the MongoDB URI** to the `.env` file:
   ```bash
   MONGO_URI=mongodb+srv://your-username:your-password@cluster0.wmhc1.mongodb.net/taskmanager?retryWrites=true&w=majority
   ```

### Step 4: Run the Application

Run the application to start the server:

```bash
go run main.go
```

This will start the server at `http://localhost:8080`.

### Step 5: Test the API

You can now use Postman (or any API testing tool) to interact with the API.

---

## Endpoints

### 1. **Get All Tasks**
- **Endpoint**: `GET /tasks`
- **Response**:
    ```json
    [
        {
            "_id": "ObjectId",
            "title": "Sample Task 1",
            "description": "Task description",
            "due_date": "2025-03-30T12:00:00Z",
            "status": "not started"
        }
    ]
    ```

### 2. **Get Task by ID**
- **Endpoint**: `GET /tasks/:id`
- **Response**:
    ```json
    {
        "_id": "ObjectId",
        "title": "Task 1",
        "description": "Task description",
        "due_date": "2025-03-30T12:00:00Z",
        "status": "not started"
    }
    ```
- **Error**: `404 Task not found`

### 3. **Add New Task**
- **Endpoint**: `POST /tasks`
- **Request**:
    ```json
    {
        "title": "Task 1",
        "description": "Task description",
        "due_date": "2025-03-30T12:00:00Z",
        "status": "not started"
    }
    ```
- **Response**:
    ```json
    {
        "message": "Task created",
        "task_id": "ObjectId"
    }
    ```

### 4. **Update Task**
- **Endpoint**: `PUT /tasks/:id`
- **Request**:
    ```json
    {
        "title": "Updated Task Title"
    }
    ```
- **Response**:
    ```json
    {
        "message": "Task updated",
        "task": {
            "_id": "ObjectId",
            "title": "Updated Task Title",
            "status": "not started"
        }
    }
    ```
- **Error**: `404 Task not found`

### 5. **Delete Task**
- **Endpoint**: `DELETE /tasks/:id`
- **Response**:
    ```json
    {
        "message": "Task removed"
    }
    ```
- **Error**: `404 Task not found`

---

## Task Model

```json
{
    "_id": "ObjectId", // Auto-generated by MongoDB
    "title": "string",  // Required
    "description": "string", // Optional
    "due_date": "string", // ISO8601 format, Optional
    "status": "string"  // Optional
}
```

---

## Error Responses

- **400 Bad Request**: Missing or invalid data.
    ```json
    { "error": "Title is required" }
    ```
  
- **404 Not Found**: Resource not found.
    ```json
    { "error": "Task not found" }
    ```

- **500 Internal Server Error**: Server-side error.
    ```json
    { "error": "Something went wrong" }
    ```

### Resources

- [MongoDB Documentation](https://www.mongodb.com/docs/)
- [MongoDB Atlas Setup](https://www.mongodb.com/cloud/atlas)
- [MongoDB Go Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)
