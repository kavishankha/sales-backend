# 🐙 OCTOPUS INSIGHTS LOGIC SERVICE

![Go](https://img.shields.io/badge/Go-1.20+-00ADD8?logo=go)
![Echo](https://img.shields.io/badge/Echo-Framework-3A86FF)
<!-- ![Build](https://github.com/kavishankha/octopus-insights-logic-service/actions/workflows/ci.yml/badge.svg)
![Last Commit](https://img.shields.io/github/last-commit/kavishankha/octopus-insights-logic-service.svg)
![PRs](https://img.shields.io/github/issues-pr/kavishankha/octopus-insights-logic-service.svg)
![Release](https://img.shields.io/github/v/release/kavishankha/octopus-insights-logic-service.svg) -->

A clean and scalable Go application built using the **Echo** framework.  
This project follows a modular structure to ensure maintainability, scalability, and reusability for microservices or monolithic systems.

---

## 🗂️ Project Structure

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
```

---

## 📁 Folder Breakdown

### 🧩 `cmd/`
> Main entry point of the application. Initializes configuration, sets up the database, and starts the Echo server.

### ⚙️ `config/`
> Contains app-level config files including environment loaders and database settings.

### 📌 `routes/`
> Contains application route definitions and handlers mapping URLs to controller functions.

### 🧾 `tests/`
> Contains unit and integration tests for validating the application's functionality.

### 🔐 `internal/` — Application Logic
A core part of the system. Split into functional sub-packages:

- **`handlers/`** - Route controllers for HTTP requests.
- **`middlewares/`** - Custom Echo middleware (auth, logging, etc.).
- **`repositories/`** - Database abstraction layer (CRUD operations).
- **`services/`** - Business logic.
- **`models/`** - Structs representing domain and DB models.
- **`integrations/`** - Interfaces and clients for third-party APIs.

### 📦 `pkg/` — Shared Utilities
Contains reusable packages used across the app:

- **`cloud/`** - Cloud serivces.
- **`logger/`** - Structured logging with custom formatters.
- **`database/`** - DB connection management.
- **`mock/`** - Mock data.
- **`constant/`** - Constant values.
- **`utils/`** - Helpers like hashing, JWT handling, etc.

---

## 📄 Key Files

- **`.env`**  
  > Environment variables like DB credentials, secrets, etc.

- **`go.mod` / `go.sum`**  
  > Go module files managing project dependencies and their versions.

---




## 📦 Dependencies

| Package                                                              | Version     | Purpose                                      |
|----------------------------------------------------------------------|-------------|----------------------------------------------|
| [`github.com/labstack/echo/v4`](https://echo.labstack.com)          | v4.13.3     | Fast, minimalist web framework               |
| [`github.com/joho/godotenv`](https://github.com/joho/godotenv)      | v1.5.1      | Loads `.env` files into Go apps              |
| [`github.com/go-playground/validator/v10`](https://github.com/go-playground/validator) | v10.26.0 | Input validation                             |
| [`github.com/machinebox/graphql`](https://github.com/machinebox/graphql) | v0.2.2  | Simple GraphQL client                        |
| [`github.com/redis/go-redis/v9`](https://github.com/redis/go-redis) | v9.7.3      | Redis client for Go                          |


---
## 🚧 Getting Started

```bash
# Clone the repository
git clone https://github.com/your-username/octopus-insights-logic-service.git

# Navigate to the folder
cd octopus-my_app-logic-service

# Run the app local
./run.sh
```

---

## 💬 Contributing
Kavishanka Kodithuwakku


---

## 📜 License



---
