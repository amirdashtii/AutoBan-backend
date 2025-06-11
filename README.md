# AutoBan 🚗

Vehicle Maintenance Management System

## About
AutoBan is a vehicle maintenance management system built with Go and PostgreSQL. It allows users to manage their vehicle maintenance records and service schedules efficiently.

## Technologies
- Go
- PostgreSQL
- Docker & Docker Compose
- JWT for Authentication

## Prerequisites
- Go 1.x or higher
- Docker and Docker Compose
- An `.env` file (you can copy from `.env.example`)

## Installation & Setup

1. Clone the repository:
```bash
git clone https://github.com/yourusername/AutoBan.git
cd AutoBan
```

2. Copy the environment file:
```bash
cp .env.example .env
```

3. Run the database with Docker:
```bash
docker-compose up -d
```

4. Run the application:
```bash
go run cmd/app/main.go
```

## API Structure

### Authentication
- `POST /api/v1/auth/register` - Register new user
- `POST /api/v1/auth/login` - User login
- `POST /api/v1/auth/logout` - User logout (requires authentication)
- `POST /api/v1/auth/refresh-token` - Refresh access token
- `GET /api/v1/auth/sessions` - Get user's active sessions (requires authentication)
- `POST /api/v1/auth/logout-all` - Logout from all devices (requires authentication)

### Users
- `GET /api/v1/users/me` - Get user profile
- `PUT /api/v1/users/me` - Update user profile
- `PUT /api/v1/users/me/change-password` - Change user password
- `DELETE /api/v1/users/me` - Delete user account

### Admin
- `GET /api/v1/users` - List all users
- `GET /api/v1/users/{id}` - Get specific user details
- `PUT /api/v1/users/{id}` - Update user information
- `POST /api/v1/users/{id}/role` - Change user role
- `POST /api/v1/users/{id}/status` - Change user status
- `POST /api/v1/users/{id}/change-password` - Change user password
- `DELETE /api/v1/users/{id}` - Delete user

## Upcoming Features
- Periodic service management
- Maintenance and repair records
- Service reminders
- Cost reporting
- Multi-vehicle management
- Mobile application

## Contributing
We welcome contributions to this project. Please submit a Pull Request to contribute.

## License
This project is licensed under the MIT License. See the LICENSE file for more details. 