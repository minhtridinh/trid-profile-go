TriD Profile API

A RESTful API for managing user profiles, built with Go, Gin, and PostgreSQL.

## Prerequisites

- Go 1.22+
- PostgreSQL
- JetBrains GoLand or IntelliJ IDEA with Go plugin

## Setup

1. Clone the repository:

   ```bash
   git clone github.com/dinhminhtri/triD-profile
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Create a `.env` file based on `.env.example` and configure your environment variables.

4. Run the application:

   ```bash
   go run cmd/api/main.go
   ```

## API Endpoints

- GET `/api/v1/profiles/:id` - Get a user profile by ID
- POST `/api/v1/profiles` - Create a new profile
- PUT `/api/v1/profiles/:id` - Update a profile (requires JWT)
- POST `/api/v1/users/register` - Register a new user
- POST `/api/v1/users/login` - Login and get JWT

## Database

- Uses PostgreSQL with GORM for ORM.

- Run migrations to set up the database schema:

  ```bash
  scripts/migrate.sh
  ```

## License

MIT