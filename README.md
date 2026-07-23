# 📊 Satudata API

**Satudata** adalah backend API untuk portal data terbuka yang menyediakan akses ke dataset statistik, artikel analitis, infografis, dan standar data. Dibangun dengan Go dan dirancang untuk performa tinggi dengan caching Redis, object storage MinIO, dan messaging NATS.

---

## 🏗️ Tech Stack

| Komponen | Teknologi |
|----------|-----------|
| **Language** | Go 1.23+ |
| **Framework** | Gin |
| **Database** | PostgreSQL 16 |
| **Cache** | Redis 7 |
| **Object Storage** | MinIO (S3-compatible) |
| **Messaging** | NATS |
| **Auth** | JWT (bcrypt) |
| **Logger** | Zap |

---

## 📁 Project Structure

```
satudata/
├── cmd/
│   ├── api/          # Entry point API server
│   └── seed/         # Database seeder
├── internal/
│   ├── config/       # Configuration management
│   ├── domain/       # Domain models (Release, Standard, User, AuditLog)
│   ├── handler/      # HTTP handlers
│   │   └── dto/      # Request/Response DTOs
│   ├── middleware/   # Auth, CORS, Rate Limiter, Logger, Metrics
│   ├── messaging/    # NATS client & event handlers
│   ├── repository/   # Data access layer (with Redis cache)
│   ├── router/       # Route configuration
│   └── service/      # Business logic layer
├── migrations/       # SQL migration files
├── pkg/
│   ├── cache/        # Redis client
│   ├── database/     # PostgreSQL/GORM connection
│   ├── logger/       # Zap logger
│   └── storage/      # MinIO client
├── Dockerfile
├── docker-compose.yml
├── Makefile
└── .env.example
```

---

## 🚀 Quick Start

### Prerequisites

- Go 1.23+
- PostgreSQL 16
- Redis 7
- MinIO (atau S3-compatible storage)
- NATS (opsional)

### 1. Clone & Setup

```bash
git clone <repo-url>
cd satudata
cp .env.example .env
```

### 2. Start Dependencies (Docker)

```bash
docker compose up -d postgres redis minio
```

### 3. Run Migrations

```bash
# Menggunakan psql langsung
cat migrations/*.sql | psql -h localhost -U satudata -d satudata
```

### 4. Seed Data (Opsional)

```bash
make seed
```

### 5. Run Server

```bash
make dev
# atau
go run ./cmd/api/main.go
```

Server akan berjalan di `http://localhost:8080`

---

## 🔌 API Endpoints

### Public Endpoints (No Auth)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check semua services |
| `GET` | `/api/v1/public/health` | API health check |
| `GET` | `/api/v1/public/releases` | Daftar releases (dengan filter) |
| `GET` | `/api/v1/public/releases/stats` | Statistik releases |
| `GET` | `/api/v1/public/releases/:id` | Detail release by ID |
| `GET` | `/api/v1/public/releases/slug/:slug` | Detail release by slug |
| `GET` | `/api/v1/public/standards` | Daftar standar data |

### Auth Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `POST` | `/api/v1/auth/login` | Login |
| `POST` | `/api/v1/auth/register` | Register |
| `POST` | `/api/v1/auth/refresh` | Refresh token |
| `POST` | `/api/v1/auth/logout` | Logout |

### Protected Endpoints (Auth Required)

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/api/v1/protected/profile` | Profile user |
| `POST` | `/api/v1/protected/releases` | Buat release baru |
| `PUT` | `/api/v1/protected/releases/:id` | Update release |
| `POST` | `/api/v1/protected/standards` | Buat standard baru |
| `PUT` | `/api/v1/protected/standards/:id` | Update standard |

### Admin Endpoints (Admin Only)

| Method | Path | Description |
|--------|------|-------------|
| `DELETE` | `/api/v1/admin/releases/:id` | Hapus release |
| `GET` | `/api/v1/admin/users` | Daftar users |
| `PUT` | `/api/v1/admin/users/:id/role` | Update role user |
| `GET` | `/api/v1/admin/audit-logs` | Audit logs |

---

## 🔐 Authentication

Gunakan JWT Bearer token:

```bash
# Login
curl -X POST http://localhost:8080/api/v1/auth/login \
  -H "Content-Type: application/json" \
  -d '{"email":"admin@satudata.go.id","password":"admin123"}'

# Gunakan token
curl -X GET http://localhost:8080/api/v1/protected/profile \
  -H "Authorization: Bearer <token>"
```

### Seed Credentials

| Role | Email | Password |
|------|-------|----------|
| Admin | admin@satudata.go.id | admin123 |
| Editor | editor@satudata.go.id | editor123 |

---

## 🧪 Testing

```bash
# Run all tests
make test

# Run with coverage
make test-coverage
```

---

## 🐳 Docker Deployment

```bash
# Build & run all services
make docker-up

# Build image saja
make docker-build

# Stop services
make docker-down
```

---

## ⚙️ Environment Variables

See `.env.example` for all available configurations.

| Variable | Default | Description |
|----------|---------|-------------|
| `SERVER_PORT` | 8080 | HTTP server port |
| `DB_HOST` | localhost | PostgreSQL host |
| `DB_PORT` | 5432 | PostgreSQL port |
| `REDIS_HOST` | localhost | Redis host |
| `MINIO_ENDPOINT` | localhost:9000 | MinIO endpoint |
| `NATS_URL` | nats://localhost:4222 | NATS server URL |
| `JWT_SECRET` | supersecretkey | JWT signing secret |
| `APP_ENV` | development | Environment |

---

## 📊 Architecture Overview

```
┌──────────┐     ┌──────────┐     ┌──────────┐
│  Client  │────▶│  Gin API │────▶│  Service │
└──────────┘     └──────────┘     └──────────┘
                       │                │
                       ▼                ▼
                 ┌──────────┐     ┌──────────┐
                 │Middleware│     │  NATS    │
                 │ - Auth   │     │ (Events) │
                 │ - CORS   │     └──────────┘
                 │ - Logger │
                 │ - Rate   │
                 └──────────┘
                       │
         ┌─────────────┼─────────────┐
         ▼             ▼             ▼
   ┌──────────┐  ┌──────────┐  ┌──────────┐
   │PostgreSQL│  │  Redis   │  │  MinIO   │
   └──────────┘  └──────────┘  └──────────┘
```

---

## 📝 License

MIT
