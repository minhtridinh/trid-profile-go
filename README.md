TriD Profile API

A RESTful API for managing user profiles, built with Go, Gin, and PostgreSQL.

## Prerequisites

- Go 1.22+
- PostgreSQL
- JetBrains GoLand or IntelliJ IDEA with Go plugin

## Setup

1. Clone the repository:

   ```bash
   git clone github.com/minhtridinh/trid-profile-go
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file based on `.env.example` and configure your environment variables.

4. Generate a secure JWT secret key for authentication:

   ```bash
   # Make the script executable
   chmod +x scripts/generate_jwt_secret.sh
   
   # Run the script to generate and update JWT_SECRET in .env
   ./scripts/generate_jwt_secret.sh
   ```

5. Run the application:

   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

- GET `/api/v1/profiles/:id` - Get a user profile by ID
- POST `/api/v1/profiles` - Create a new profile
- PUT `/api/v1/profiles/:id` - Update a profile (requires JWT)
- POST `/api/v1/users/register` - Register a new user
- POST `/api/v1/users/login` - Login and get JWT

### Authentication Endpoints

- POST `/api/v1/auth/register` - Register a new user
- POST `/api/v1/auth/login` - Login and get JWT token

### User Endpoints (Protected - requires JWT authentication)

- GET `/api/v1/users/me` - Get current user's information
- PUT `/api/v1/users/me` - Update current user's information

### Admin Endpoints (Protected - requires admin role)

- GET `/api/v1/admin/users` - Get all users
- GET `/api/v1/admin/users/:id` - Get a user by ID
- DELETE `/api/v1/admin/users/:id` - Delete a user

## Database

- Uses PostgreSQL with GORM for ORM.

- Run migrations to set up the database schema:

  ```bash
  scripts/migrate.sh
  ```

## Commands

Here are the available utility commands in the project:

1. **Generate JWT Secret**

   ```bash
   # Generate a secure random JWT_SECRET and update it in .env file
   ./scripts/generate_jwt_secret.sh
   ```

2. **Run Database Migrations**

   ```bash
   # Apply database migrations
   ./scripts/migrate.sh
   ```

## License

MIT

