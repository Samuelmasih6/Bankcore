# Bankcore

A production-inspired backend banking system built in **Go**, following clean architecture and industry best practices. The project demonstrates how modern financial services manage accounts, money transfers, transactions, authentication, containerization, cloud-native deployment, and orchestration.

This project is designed to gain hands-on experience with backend engineering concepts commonly used in fintech and large-scale distributed systems.

---

## Features

- User account management
- Secure money transfers between accounts
- ACID-compliant database transactions
- RESTful API built with Gin
- PostgreSQL database
- SQLC for type-safe SQL generation
- Database migrations with golang-migrate
- Dockerized application
- Docker Compose for local development
- JWT-based authentication
- Password hashing with bcrypt
- Token-based authorization middleware
- Unit and integration testing
- Mock generation for testing
- GitHub Actions CI pipeline
- AWS deployment preparation
- Kubernetes deployment and orchestration

---

## Tech Stack

### Backend
- Go
- Gin
- PostgreSQL
- pgx/v5
- SQLC
- golang-migrate

### Authentication
- JWT
- bcrypt

### Testing
- Testify
- Mockgen

### DevOps
- Docker
- Docker Compose
- GitHub Actions
- Kubernetes

### Cloud
- AWS EC2
- Amazon ECR
- Amazon RDS

---

## Project Structure

```
Bankcore/
├── api/                 # HTTP handlers and routes
├── db/
│   ├── migration/       # Database migrations
│   ├── mock/            # Generated mocks
│   └── sqlc/            # Generated SQL queries
├── token/               # JWT implementation
├── util/                # Helper utilities
├── worker/              # Background workers (future)
├── docs/                # API documentation
├── scripts/             # Deployment scripts
├── Dockerfile
├── docker-compose.yml
├── Makefile
├── go.mod
└── main.go
```

---

## Architecture

```
                Client
                   │
             HTTP Request
                   │
               Gin Router
                   │
            Authentication
                   │
             API Handlers
                   │
             Business Logic
                   │
               SQL Store
                   │
          SQLC Generated Queries
                   │
              PostgreSQL
```

---

## Implemented Modules

### Database
- PostgreSQL
- Schema migrations
- Type-safe SQL generation using SQLC
- Connection pooling with pgx

### Banking
- Create accounts
- Retrieve accounts
- List accounts
- Update balances
- Money transfers
- Transaction history
- Atomic transfer transactions

### Authentication
- User registration
- Password hashing
- User login
- JWT access tokens
- Protected routes

### Testing
- CRUD tests
- Transaction tests
- API tests
- Mock generation
- Integration tests

### DevOps
- Docker containerization
- Docker Compose
- GitHub Actions CI
- Multi-stage Docker builds
- AWS container registry (ECR)
- Kubernetes deployment

---

## Running Locally

### Clone Repository

```bash
git clone https://github.com/Samuelmasih6/Bankcore.git
cd Bankcore
```

### Start PostgreSQL

```bash
docker compose up -d
```

### Run Migrations

```bash
make migrateup
```

### Start Server

```bash
make server
```

---

## Running Tests

```bash
make test
```

Generate test coverage

```bash
make cover
```

---

## Docker

Build image

```bash
docker build -t bankcore .
```

Run container

```bash
docker run -p 8080:8080 bankcore
```

---

## Kubernetes

The project includes Kubernetes manifests for deploying:

- Backend API
- PostgreSQL
- Services
- Deployments
- Configurations
- Secrets

Deploy using:

```bash
kubectl apply -f k8s/
```

---

## CI/CD

GitHub Actions automatically:

- Builds the application
- Runs unit tests
- Executes integration tests
- Builds Docker images
- Pushes images to Amazon ECR
- Prepares Kubernetes deployment

---

## API Endpoints

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | `/users` | Register user |
| POST | `/users/login` | Login |
| POST | `/accounts` | Create account |
| GET | `/accounts/:id` | Get account |
| GET | `/accounts` | List accounts |
| POST | `/transfers` | Transfer money |

---

## Future Improvements

- Refresh tokens
- Email verification
- Password reset
- Redis caching
- Rate limiting
- Prometheus monitoring
- Grafana dashboards
- Distributed tracing
- gRPC services
- Microservices architecture

---

## Learning Outcomes

Through this project, I gained practical experience with:

- Backend development in Go
- Database design
- SQL optimization
- Transaction management
- REST API development
- Authentication and authorization
- Testing strategies
- Docker containerization
- CI/CD pipelines
- AWS deployment
- Kubernetes orchestration
- Production-ready backend architecture

---

## Author

**Samuel Masih**

GitHub: https://github.com/Samuelmasih6

---

## Acknowledgements

This project is inspired by real-world backend engineering practices and modern cloud-native application development.
