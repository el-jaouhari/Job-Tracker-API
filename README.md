# Job Tracker API

A RESTful API built with Go and Gin for tracking job applications. This API allows you to manage job applications, track their status, and organize your job search process.

## Features

- ✅ Create, read, update, and delete job applications
- ✅ Track application status (applied, interviewing, rejected, offer)
- ✅ PostgreSQL database for persistent storage
- ✅ Docker support for easy deployment
- ✅ RESTful API design

## Tech Stack

- **Language**: Go 1.24
- **Web Framework**: Gin
- **Database**: PostgreSQL 16.2
- **ORM**: GORM
- **Containerization**: Docker & Docker Compose

## Prerequisites

- Go 1.24 or higher
- Docker and Docker Compose (for containerized deployment)
- PostgreSQL (for local development without Docker)

## Project Structure

```
Job-Tracker-API/
├── cmd/
│   └── service/
│       └── main.go          # Application entry point
├── db/
│   └── schema.sql           # Database schema
├── internal/
│   ├── httpx/
│   │   └── httpx.go         # HTTP handlers and routes
│   ├── repository/
│   │   └── jobs.go          # Data access layer
│   └── service/
│       └── jobs.go          # Business logic layer
├── docker-compose.yml       # Docker Compose configuration
├── Dockerfile              # Docker image definition
├── go.mod                  # Go module dependencies
├── go.sum                  # Go module checksums
└── makefile                # Make commands
```

## Getting Started

### Local Development

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd Job-Tracker-API
   ```

2. **Set up the database**
   
   Start PostgreSQL using Docker Compose:
   ```bash
   make db-up
   ```
   
   Or manually:
   ```bash
   docker compose up -d db
   ```

3. **Install dependencies**
   ```bash
   go mod download
   ```

4. **Set environment variables (optional)**
   ```bash
   export DATABASE_URL="postgres://myuser:mypassword@localhost:5432/job_tracker?sslmode=disable"
   ```
   
   If not set, the application defaults to `postgres://localhost:5432/job_tracker`

5. **Run the application**
   ```bash
   make run
   ```
   
   Or directly:
   ```bash
   go run cmd/service/main.go
   ```

   The API will be available at `http://localhost:8080`

### Docker Deployment

1. **Build and start all services**
   ```bash
   docker compose up --build
   ```

2. **Run in detached mode**
   ```bash
   docker compose up -d --build
   ```

3. **View logs**
   ```bash
   docker compose logs -f app
   ```

4. **Stop services**
   ```bash
   docker compose down
   ```

The application will be available at `http://localhost:8080` and the database at `localhost:5432`.

## API Endpoints

### Get All Jobs
```http
GET /jobs
```

**Response:**
```json
[
  {
    "title": "Software Engineer",
    "company": "Tech Corp",
    "url": "https://example.com/job",
    "location": "San Francisco, CA",
    "type": "Full-time",
    "application_status": "applied"
  }
]
```

### Create a Job
```http
POST /jobs
Content-Type: application/json

{
  "title": "Software Engineer",
  "company": "Tech Corp",
  "url": "https://example.com/job",
  "location": "San Francisco, CA",
  "type": "Full-time",
  "application_status": "applied"
}
```

**Response:**
```json
{
  "message": "Job created successfully"
}
```

### Update Job Status
```http
PUT /jobs/:id?status=interviewing
```

**Status values**: `applied`, `interviewing`, `rejected`, `offer`

**Response:**
```json
{
  "message": "Job updated successfully"
}
```

### Delete a Job
```http
DELETE /jobs/:id
```

**Response:**
```json
{
  "message": "Job deleted successfully"
}
```

## Makefile Commands

- `make run` - Run the application locally
- `make build` - Build the application binary
- `make db-up` - Start the database container
- `make db-down` - Stop the database container
- `make db-restart` - Restart the database container
- `make db-logs` - View database container logs

## Database Schema

The `jobs` table has the following structure:

- `id` - Primary key (SERIAL)
- `title` - Job title (VARCHAR)
- `company` - Company name (VARCHAR)
- `url` - Job posting URL (VARCHAR)
- `location` - Job location (VARCHAR)
- `type` - Job type (VARCHAR)
- `application_status` - Status enum: `applied`, `interviewing`, `rejected`, `offer`
- `created_at` - Timestamp
- `updated_at` - Timestamp

## Environment Variables

| Variable | Description | Default |
|----------|-------------|---------|
| `DATABASE_URL` | PostgreSQL connection string | `postgres://localhost:5432/job_tracker` |

## Development

### Database Connection

The application uses GORM with PostgreSQL. The connection string format is:
```
postgres://user:password@host:port/database?sslmode=disable
```

### Running Tests

```bash
go test ./...
```

## License

[Add your license here]

## Contributing

[Add contribution guidelines here]

