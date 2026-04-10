# 📇 Contacts Agenda API (Go + Gin + SQLite)

Simple REST API to manage your contacts via JWT authentication.

## Techs

- Go (Golang)
- Gin (framework HTTP)
- SQLite
- JWT (autenticação)
- Bcrypt (hash de senha)

## Project Structure

```
contact-api
├───cmd
│   └───api/main.go # Entry point
├───internal
│   ├───database/ # Database connection
│   ├───handler/ # Handlers HTTP
│   ├───middleware/ # Middlewares (JWT)
│   └───model/ # Structures (DTOs)
```

## Running

### 1. Clone repo

```bash
git clone https://github.com/rogeriods/contacts-api.git
cd contacts-api
```

### 2. Install Packages

```bash
go mod tidy
```

### 3. Run Project

```bash
go run cmd/api/main.go
```

On resources folder we have 'api.http' to test API.
