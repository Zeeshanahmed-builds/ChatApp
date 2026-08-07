# Chat Application Backend

A scalable real-time chat application backend built with **Go (Golang)** using the **Gin Framework**, **PostgreSQL**, **MQTT (EMQX)**, **Docker**, **Kubernetes**, and **GitHub Actions CI/CD**.

---

# Features

- User Registration
- User Login
- JWT Authentication
- Protected APIs
- User Management
- Private Messaging
- Message Persistence
- MQTT Real-time Communication
- PostgreSQL Database
- Database Migrations
- Docker Support
- Kubernetes Deployment
- GitHub Actions CI/CD

---

# Technology Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend Language |
| Gin | HTTP Framework |
| PostgreSQL | Database |
| GORM | ORM |
| MQTT (EMQX) | Real-time Messaging |
| JWT | Authentication |
| Docker | Containerization |
| Kubernetes | Container Orchestration |
| GitHub Actions | CI/CD |

---

# Project Structure

```
.
├── db
│   ├── database.go
│   └── migrations
│
├── handler
│   ├── auth.go
│   ├── message.go
│   └── users.go
│
├── middleware
│   └── auth.go
│
├── models
│   ├── gorm.go
│   └── request.go
│
├── mqtt
│   ├── mqtt.go
│   └── subscribe.go
│
├── repo
│   ├── message_repo
│   └── users_repo
│
├── routes
│
├── service
│   ├── message_service
│   └── users_service
│
├── utils
│
├── k8s
│
├── Dockerfile
├── docker-compose.yaml
├── main.go
└── README.md
```

---

# Folder Explanation

## db

Contains database configuration and SQL migration files.

```
db/
```

Responsible for

- PostgreSQL Connection
- Database Initialization
- SQL Migrations

---

## handler

Contains HTTP handlers.

Responsibilities:

- Handle Requests
- Validate Request
- Call Services
- Return Response

Example:

```
POST /signup

↓

handler/auth.go

↓

User Service

↓

Repository

↓

Database
```

---

## middleware

Contains middleware.

Current middleware:

- JWT Authentication

Responsibilities:

- Verify Token
- Protect Routes
- Extract User ID

---

## models

Contains application models.

Includes

- Database Models
- Request Models
- Response Models

---

## mqtt

Contains MQTT configuration.

Responsible for

- MQTT Connection
- Publisher
- Subscriber

Uses EMQX Broker.

---

## repo

Repository Layer.

Responsible for

- Database Queries
- CRUD Operations
- SQL/GORM Logic

Repositories:

- Users Repository
- Message Repository

---

## service

Business Logic Layer.

Responsible for

- Authentication Logic
- User Logic
- Message Logic

This layer communicates between Handlers and Repositories.

---

## routes

Contains all API routes.

Responsible for

- Registering APIs
- Applying Middleware

---

## utils

Utility functions.

Examples:

- Generate JWT
- Verify JWT

---

# Kubernetes Structure

```
k8s
│
├── api
│
├── postgres
│
├── emqx
│
├── ingress.yaml
│
├── namespace.yaml
│
└── secret.yaml
```

---

## namespace.yaml

Creates the Kubernetes namespace.

```
chat-app
```

---

## secret.yaml

Stores sensitive data.

Examples

- Database Password
- JWT Secret
- MQTT Credentials

---

## api

Contains

### deployment.yaml

Deploys Backend Application.

### service.yaml

Creates Backend Service.

---

## postgres

Contains PostgreSQL resources.

### deployment.yaml

Deploy PostgreSQL.

### service.yaml

Expose PostgreSQL inside the cluster.

### pvc.yaml

Persistent Volume Claim for database storage.

---

## emqx

Contains EMQX Broker resources.

### deployment.yaml

Deploy MQTT Broker.

### service.yaml

Expose MQTT Broker.

---

## ingress.yaml

Exposes backend using Kubernetes Ingress.

---

# Docker

## Build Image

```bash
docker build -t chat-app .
```

---

## Run Container

```bash
docker run -p 8080:8080 chat-app
```

---

# Docker Compose

Run all services.

```bash
docker compose up -d
```

---

# Database Migration

Run migration tool.

Example

```bash
make migrate-up
```

Rollback

```bash
make migrate-down
```

---

# Run Application

```bash
go run main.go
```

---

# Kubernetes Deployment

Create namespace

```bash
kubectl apply -f k8s/namespace.yaml
```

Deploy secrets

```bash
kubectl apply -f k8s/secret.yaml
```

Deploy PostgreSQL

```bash
kubectl apply -f k8s/postgres/
```

Deploy EMQX

```bash
kubectl apply -f k8s/emqx/
```

Deploy Backend

```bash
kubectl apply -f k8s/api/
```

Deploy Ingress

```bash
kubectl apply -f k8s/ingress.yaml
```

---

# Verify Resources

Pods

```bash
kubectl get pods -n chat-app
```

Services

```bash
kubectl get svc -n chat-app
```

Ingress

```bash
kubectl get ingress -n chat-app
```

---

# CI Pipeline

GitHub Actions automatically performs:

- Checkout Repository
- Install Go
- Download Dependencies
- Verify Dependencies
- Check Formatting
- Run Tests
- Build Project

---

# CD Pipeline

GitHub Actions automatically:

- Checkout Repository
- Login to Docker Hub
- Build Docker Image
- Push Docker Image

Image:

```
zeeshan1678/chat-app:latest
```

---

# API Endpoints

## Authentication

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /signup | Register User |
| POST | /login | Login |

---

## Users

| Method | Endpoint | Description |
|---------|----------|-------------|
| GET | /users | Get Users |

---

## Messages

| Method | Endpoint | Description |
|---------|----------|-------------|
| POST | /message | Send Message |
| POST | /message/history | Get Conversation History |

---

# Architecture

```
                Client
                   │
             HTTP Request
                   │
                   ▼
              Gin Router
                   │
                   ▼
             Authentication
             (JWT Middleware)
                   │
                   ▼
               Handlers
                   │
                   ▼
               Services
                   │
                   ▼
            Repository Layer
                   │
          ┌────────┴─────────┐
          ▼                  ▼
     PostgreSQL          MQTT Broker
        (GORM)             (EMQX)
```

---
