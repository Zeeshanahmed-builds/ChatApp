# Chat Application Backend

A real-time chat application backend built with **Go (Golang)**, **Gin**, **GORM**, **PostgreSQL**, and **EMQX MQTT**. The project supports user authentication, private messaging, group chats, and message history.

---

## Features

- User Registration
- User Login with JWT Authentication
- Protected API Routes
- Private One-to-One Messaging
- Group Creation
- Add/Remove Group Members
- Group Messaging
- Message History
- MQTT Integration using EMQX
- PostgreSQL Database
- Repository-Service-Handler Architecture

---

## Tech Stack

| Technology | Purpose |
|------------|---------|
| Go | Backend Language |
| Gin | HTTP Framework |
| PostgreSQL | Database |
| GORM | ORM |
| MQTT (EMQX) | Real-time Messaging |
| JWT | Authentication |

---

## Project Structure

```
Backend/
│
├── db/
│   ├── migrations/
│   └── db.go
│
├── handler/
│
├── middleware/
│
├── models/
│
├── mqtt/
│
├── repo/
│   ├── users_repo/
│   ├── message_repo/
│   └── group_repo/
│
├── routes/
│
├── service/
│   ├── users_service/
│   ├── message_service/
│   └── group_service/
│
├── utils/
│
├── main.go
├── go.mod
└── .env
```

---

# Architecture

```
Client
   │
   ▼
Routes
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ▼
Repositories
   │
   ▼
PostgreSQL
```

Real-time messages are published through **EMQX MQTT**.

---

# Installation

## 1. Clone Repository

```bash
git clone <repository-url>

cd Backend
```

---

## 2. Install Dependencies

```bash
go mod tidy
```

---

## 3. Create Environment File

Create a `.env` file.

Example:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=chatapp

JWT_SECRET=your_secret_key

MQTT_BROKER=tcp://localhost:1883
MQTT_CLIENT_ID=chat-backend
```

---

## 4. Run PostgreSQL

Make sure PostgreSQL is running and create the database.

Example:

```sql
CREATE DATABASE chatapp;
```

---

## 5. Run Database Migrations

Run all migrations inside:

```
db/migrations/
```

---

## 6. Start EMQX

Using Docker:

```bash
docker run -d \
--name emqx \
-p 1883:1883 \
-p 8083:8083 \
-p 8084:8084 \
-p 8883:8883 \
-p 18083:18083 \
emqx/emqx
```

Dashboard:

```
http://localhost:18083
```

Default credentials:

```
Username: admin
Password: public
```

---

## 7. Run the Server

```bash
go run main.go
```

Server:

```
http://localhost:8080
```

---

# API Endpoints

## Authentication

### Register

```
POST /signup
```

### Login

```
POST /login
```

Returns a JWT token.

---

## Private Messages

### Send Message

```
POST /message
```

Authentication Required

---

### Get Conversation History

```
POST /message/history
```

Authentication Required

---

## Groups

### Create Group

```
POST /group
```

---

### Get User Groups

```
GET /group/my-groups
```

---

### Get Group Details

```
GET /group/:id
```

---

### Get Group Members

```
GET /group/:id/members
```

---

### Add Members

```
POST /group/members
```

---

### Remove Member

```
POST /group/members/remove
```

---

### Send Group Message

```
POST /group/message
```

---

### Group Message History

```
POST /group/message/history
```

---

# MQTT

Broker:

```
EMQX
```

The application connects to the MQTT broker on startup.

Messages are published and subscribed through MQTT to support real-time communication.

---

# Authentication

Protected routes require:

```
Authorization: Bearer <JWT_TOKEN>
```

---

# Dependencies

- Gin
- GORM
- PostgreSQL Driver
- Eclipse Paho MQTT
- JWT
- godotenv
- bcrypt

---

# Running the Project

```bash
go mod tidy

go run main.go
```

---

# Future Improvements

- Read Receipts
- Online/Offline Presence
- Typing Indicators
- File Sharing
- Message Reactions
- WebSocket Support
- Push Notifications

---

# License

This project is intended for educational and learning purposes.