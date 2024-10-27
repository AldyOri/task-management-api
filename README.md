# Overview

This is my personal project of me trying to learn how to build a RESTful API, and how to implement JWT for authentication in Go

# Task Management API

A RESTful API for managing tasks, built using Go with Echo, GORM, JWT for authentication, and PostgreSQL as the database. This project also includes auto-generated Swagger documentation.

## Technologies Used

- **Go** - The main language used for building the application.
- **Echo** - A fast, minimalist web framework for creating REST APIs.
- **GORM** - An ORM library for Go, used to interact with the PostgreSQL database.
- **JWT (JSON Web Token)** - For secure user authentication.
- **PostgreSQL** - The database used to store user and task data.
- **Swaggo/Swagger** - For generating API documentation.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your system.
- [PostgreSQL](https://www.postgresql.org/download/) installed and a database created.

### Setting up the .env File

Create a `.env` file in the root of your project to store environment variables. Replace the placeholders with your actual configuration:

```plaintext
DB_HOST=<db-host>
DB_USER=<db-user>
DB_PASSWORD=<db-password>
DB_NAME=<db-name>
DB_PORT=<db-port>
JWT_SECRET=<jwt-secret>
```

### Installation and Setup

1. **Clone the repository**:
   ```bash
   git clone <repository-url>
   cd <repository-name>
   ```

2. **Install dependencies**:
   ```bash
   go mod download
   ```

3. **Run the Application**:
   ```bash
   go run main.go
   ```

The API should now be running at `http://localhost:8000`.

### API Documentation

The Swagger documentation is available at:

- **Swagger UI**: `http://localhost:8000/swagger/index.html`

The documentation includes all API endpoints and allows testing directly in the browser. The JWT token should be set in the `Authorization` header using the "Bearer" scheme to test authenticated endpoints.

### Endpoints

Here is an overview of the main endpoints:

#### Auth Routes
- **POST** `/api/auth/register` - Register a new user.
- **POST** `/api/auth/login` - Log in and receive a JWT token.
- **POST** `/api/auth/me` - Retrieve the current user’s information (requires JWT).

#### Task Routes (Protected)
- **POST** `/api/tasks` - Create a new task.
- **GET** `/api/tasks` - Retrieve all tasks for the authenticated user.
  - `completed` (bool): Filter tasks by completion status (true or false).
- **GET** `/api/tasks/:id` - Retrieve a specific task by its ID.
- **PATCH** `/api/tasks/:id` - Update a specific task by its ID.
- **DELETE** `/api/tasks/:id` - Delete a specific task by its ID.

#### Image Routes
- **POST** `/api/tasks/:task_id/images` - Upload an image for a specific task. (requires JWT).
- **GET**  `/api/images/:id` - Retrieve a specific image by its ID.
- **DELETE**  `/api/images/:id` - Delete a specific image by its ID.

### Project Structure

- `controllers/` - Contains route handler functions for each endpoint.
- `models/` - Defines data models for GORM and structures for request/response formats.
- `routes/` - Routes for API endpoint
- `middleware/` - JWT authentication and image middleware.
- `utils/` - Helper functions for extracting user ID from the JWT token and extracting task ID from route params.
- `config/` - Database connection setup and environment variable management.
- `docs/` - Documentation for the API.

### Environment Variables

| Variable      | Description                             |
|---------------|-----------------------------------------|
| `DB_HOST`     | Database host                           |
| `DB_USER`     | Database username                       |
| `DB_PASSWORD` | Database password                       |
| `DB_NAME`     | Database name                           |
| `DB_PORT`     | Database port (default is 5432)         |
| `JWT_SECRET`  | Secret key for signing JWT tokens       |

### Important Notes

- **JWT Authentication**: The `Authorization` header should include the token as `Bearer <token>` for protected routes.
- **Swagger Documentation**: You can access Swagger to explore and test endpoints. Make sure the app is running.

### Acknowledgments

- [Echo Framework](https://echo.labstack.com/)
- [GORM](https://gorm.io/)
- [Swaggo](https://github.com/swaggo/swag)

