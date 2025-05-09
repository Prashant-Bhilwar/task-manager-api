# Task Manager API

A simple backend API for managing user tasks, built with **Golang**, **Gin**, and **PostgreSQL**.  
Includes full **JWT-based authentication** and **CRUD operations** for tasks.

---

## Features

- User registration with hashed passwords
- Login with JWT token generation
- JWT-based route protection
- Task management for each user:
  - Create task
  - List all tasks
  - Update task
  - Delete task

---

## Tech Stack

- [Golang](https://golang.org/)
- [Gin](https://github.com/gin-gonic/gin)
- [PostgreSQL](https://www.postgresql.org/)
- [GORM](https://gorm.io/)
- JWT Authentication
- .env-based configuration

---

## Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/task-manager.git
cd task-manager
```

### 2. Install dependencies

```bash
go mod tidy
```

### 3. Create `.env` file

Create a `.env` file in the root directory and add:

```
DB_HOST=localhost
DB_USER=postgres
DB_PASSWORD=yourpassword
DB_NAME=taskdb
DB_PORT=5432
JWT_SECRET=your_jwt_secret
```

### 4. Run the application

```bash
go run main.go
```

---

## API Endpoints

### **Auth Routes**

- `POST /register`  
  **Body:**  
  ```json
  { "username": "prashant", "password": "secret123" }
  ```

- `POST /login`  
  **Response:**  
  ```json
  { "token": "<JWT Token>" }
  ```

---

### **Protected Task Routes (require JWT)**

> Add this header to all requests:
```
Authorization: Bearer <token>
```

- `POST /api/tasks`  
  Create a task  
  **Body:**  
  ```json
  { "title": "Learn Go" }
  ```

- `GET /api/tasks`  
  Get all tasks for the logged-in user

- `PUT /api/tasks/:id`  
  Update a task  
  **Body:**  
  ```json
  { "title": "Updated title", "completed": true }
  ```

- `DELETE /api/tasks/:id`  
  Delete a task

---

## Project Structure

```
task-manager/
├── main.go
├── .env
├── go.mod
├── config/          # DB connection
├── models/          # User and Task models
├── controllers/     # Auth and task handlers
├── routes/          # Route definitions
├── middleware/      # JWT auth middleware
├── utils/           # JWT token utility
```

---

## License

This project is open-source and available under the MIT License.