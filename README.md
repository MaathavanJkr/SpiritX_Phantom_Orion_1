# User Authentication API

This is a RESTful API for user authentication and management.

## Prerequisites

- Docker
- Docker Compose

## Getting Started

1. Clone the repository:
   ```bash
   git clone <repository-url>
   ```

2. Navigate to the project directory:
   ```bash
   cd <project-directory>
   ```

3. Start the application using Docker Compose:
   ```bash
   docker-compose up
   ```

4. The API will be available at `http://localhost:8080`

## API Endpoints

### Authentication
- `POST /login` - User login
- `POST /register` - User registration

### User Management
- `POST /users` - Create a new user
- `GET /users` - Get all users
- `GET /users/:id` - Get user by ID

## Environment Variables

The following environment variables can be configured in the `.env` file:

- `DB_HOST` - Database host
- `DB_PORT` - Database port
- `DB_USER` - Database username
- `DB_PASSWORD` - Database password
- `DB_NAME` - Database name
- `JWT_SECRET` - Secret key for JWT token generation

## Testing the API

You can use tools like Postman or cURL to test the API endpoints.

Example registration request:
