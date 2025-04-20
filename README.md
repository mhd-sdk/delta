# Delta App with WebAuthn Authentication

This is a modern web application that uses WebAuthn for secure, passwordless authentication.

## Features

- WebAuthn passwordless authentication
- Automatic reauthentication every 2 hours
- PostgreSQL database for storing user credentials
- React frontend with Tanstack Router
- Go backend with Fiber framework

## Requirements

- Go 1.21+
- Node.js 20+
- PostgreSQL 16+
- A browser that supports WebAuthn (most modern browsers)
- A WebAuthn-compatible authenticator (biometric sensor, security key, or device with platform authenticator)

## Running the Application

### Using Docker Compose (Recommended)

```bash
# Clone the repository
git clone https://github.com/yourusername/delta.git
cd delta

# Run the application
docker-compose up
```

### Manual Setup

#### Backend

```bash
# Navigate to the backend directory
cd backend

# Copy the example environment file
cp .env.example .env

# Edit the environment variables as needed
nano .env

# Run the backend
go run main.go
```

#### Frontend

```bash
# Navigate to the frontend directory
cd frontend

# Install dependencies
yarn install

# Run the frontend
yarn dev
```

## Accessing the Application

The application will be available at:

- Frontend: http://localhost:5173
- Backend API: http://localhost:8080

## Authentication Flow

1. Register with a username and email
2. Your device will prompt you to create a WebAuthn credential
3. Once registered, you can login with your username
4. You'll be prompted to authenticate with your registered WebAuthn credential
5. Every 2 hours, the application will request reauthentication

## Technology Stack

- **Frontend**: React, Typescript, Tanstack Router, Zustand, Tailwind CSS
- **Backend**: Go, Fiber framework, WebAuthn, JWT
- **Database**: PostgreSQL
- **Authentication**: WebAuthn (FIDO2)
