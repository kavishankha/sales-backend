# 🐙 ABT Backend Service

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Echo](https://img.shields.io/badge/Echo-Framework-3A86FF)

A clean and scalable **Go** application built using the **Echo** web framework. This project follows a modular structure to ensure maintainability, scalability, and reusability for microservices or monolithic systems.

---

## 🗂️ Project Structure

The project adheres to a standard, modular Go layout:

```bash
backend-service/
├── cmd/
├── config/
├── routes/
├── tests/
├── internal/
│   ├── handlers/
│   ├── integrations/
│   ├── middlewares/
│   ├── repositories/
│   ├── services/
│   └── models/
├── pkg/
│   ├── cloud/
│   ├── logger/
│   ├── database/
│   ├── mock/
│   ├── constant/
│   └── utils/
├── .env
├── go.mod
├── go.sum
└── README.md
````

-----

## 📁 Folder Breakdown

### 🧩 `cmd/`

> The **Main Entry Point** of the application. It's responsible for initializing configuration, setting up the database connection, and starting the Echo HTTP server.

### ⚙️ `config/`

> Contains application-level configuration files, including environment loaders and database connection settings.

### 📌 `routes/`

> Defines all application routes, mapping URLs to the corresponding handler functions in the `internal/handlers` package.

### 🧾 `tests/`

> Stores **unit and integration tests** to validate the application's functionality and business logic.

### 🔐 `internal/` — Application Logic

This is the core of the system, containing business-specific logic, split into functional sub-packages:

- **`handlers/`** - **Route controllers** that handle incoming HTTP requests and responses.
- **`middlewares/`** - Custom **Echo middleware** (e.g., authentication, request logging, rate limiting).
- **`repositories/`** - The **Database Abstraction Layer** responsible for direct data access (CRUD operations).
- **`services/`** - Implements the core **Business Logic** and orchestrates interactions between handlers and repositories.
- **`models/`** - Structs representing the **Domain and Database models** (data structures).
- **`integrations/`** - Defines interfaces and clients for interacting with **external or third-party APIs**.

### 📦 `pkg/` — Shared Utilities

Contains reusable, generic packages that can be safely used across the application and potentially in other projects:

- **`cloud/`** - Clients and abstractions for cloud services.
- **`logger/`** - Structured logging implementation with custom formatters.
- **`database/`** - Generic database connection management and setup utilities.
- **`mock/`** - Helper files for generating mock data for testing or development.
- **`constant/`** - Application-wide constant values (e.g., limits, status codes).
- **`utils/`** - General-purpose helpers (e.g., hashing, JWT handling, error handling).

-----

## 📄 Key Files

- **`.env`**
  > Contains sensitive environment variables necessary for the application, such as DB credentials, external API secrets, etc.
- **`go.mod` / `go.sum`**
  > Standard Go module files that manage project dependencies and their exact versions.

-----

## 📦 Dependencies

| Package | Version | Purpose |
| :--- | :--- | :--- |
| [`github.com/labstack/echo/v4`](https://echo.labstack.com) | v4.13.4 | Fast, minimalist web framework |
| [`github.com/joho/godotenv`](https://github.com/joho/godotenv) | v1.5.1 | Loads environment variables from `.env` files |
| [`github.com/go-playground/validator/v10`](https://github.com/go-playground/validator) | v10.27.0 | Input validation library |
| [`github.com/jackc/pgx/v5`](https://github.com/jackc/pgx) | v5.7.6 | PostgreSQL driver (modern, fast) |
| `github.com/lib/pq` | v1.10.9 | Legacy PostgreSQL driver |
| `github.com/bradfitz/gomemcache` | v0.0.0-20250403215159-8d39553ac7cf | Memcached client |

-----

## 🚧 Getting Started

### Prerequisites

1.  **Go** must be installed (version 1.20+ is recommended).
2.  A running **PostgreSQL** database instance (local or remote).

### Installation and Setup

```bash
# Clone the repository
git clone [https://github.com/kavishankha/sales-backend.git](https://github.com/kavishankha/sales-backend.git)

# Navigate to the project directory
cd backend-service

# Create and configure the .env file
# Copy the contents below into a new file named .env
# and update the values with your actual database credentials.
```

**.env Example:**

```ini
DB_HOST=localhost
DB_PORT=<db_port>
DB_USER=<db_user_name>
DB_PASS=<db_password>
DB_NAME=<db_name>
CSV_PATH=<csv_file_path> # Path to any required CSV file
```

### Running the Application Locally

```bash
# Run the application using the included script
./run.sh
```
```bash
# Wait until MIGRATION and ETL RUN
2025/10/07 01:55:35 Starting pprof on :6060
2025/10/07 01:55:35 SQL DB connected successfully!
2025/10/07 01:55:35 Tables migration complete!
2025/10/07 01:55:35 Materialized views created successfully!
2025/10/07 01:55:35 PGX connection established for high-performance operations
2025/10/07 01:56:05 CSV loaded into staging successfully.
2025/10/07 02:01:36 Data normalized successfully.
2025/10/07 02:02:18 Materialized views refreshed successfully!

```
-----

## 🧪 Testing

### Running Unit Tests

Run all tests across all packages with verbose output:

```bash
go test -v ./...
```

Run all tests and generate a coverage report:

```bash
go test -v -cover ./...
```

-----

## 📊 Profiling

The application exposes the standard Go pprof handlers for diagnostics when run locally.

### Accessing Profiling Data

Access the pprof index in your browser:

**Browser:** `http://localhost:6060/debug/pprof/`

### Collecting a CPU Profile

Use the `go tool pprof` command to collect a 30-second CPU profile and start an interactive session:

```bash
go tool pprof http://localhost:6060/debug/pprof/profile?seconds=30
```

### Optional Profiles

You can profile other aspects using the following URLs:

* **Heap:** `http://localhost:6060/debug/pprof/heap`
* **Goroutines:** `http://localhost:6060/debug/pprof/goroutine`
* **Block:** `http://localhost:6060/debug/pprof/block`
* **Mutex:** `http://localhost:6060/debug/pprof/mutex`

-----

## 💬 Contributing

This project is maintained by **Kavishanka Kodithuwakku**.

-----

## 📜 License

This project is licensed under the **[Placeholder License, e.g., MIT License]**. See the `LICENSE` file for details.

```
```