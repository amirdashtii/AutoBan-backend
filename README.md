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

## Database Setup & Configuration

### PostgreSQL Database
The project uses PostgreSQL as its primary database. The database is containerized using Docker and can be easily set up using Docker Compose.

#### Default Configuration
```yaml
Database Name: autoban
User: autoban
Password: autoban
Port: 5432
Host: localhost (when running locally)
```

#### Redis Cache
Redis is used for session management and caching.
```yaml
Port: 6379
Password: autoban (default, configurable via REDIS_PASSWORD)
Persistence: Enabled (AOF mode)
```

### Environment Variables
Create a `.env` file in the root directory using `.env.example` as a template. The following variables are required for database configuration:

```env
# PostgreSQL Configuration
DB_HOST=your_db_host      # Default: localhost
DB_PORT=your_db_port      # Default: 5432
DB_USER=your_db_user      # Default: autoban
DB_PASSWORD=your_db_password  # Default: autoban
DB_NAME=your_db_name      # Default: autoban

# Redis Configuration
REDIS_ADDR=your_redis_addr    # Default: localhost:6379
REDIS_PASSWORD=your_redis_password  # Default: autoban
REDIS_DB=your_redis_db        # Default: 0
```

### Data Persistence
- PostgreSQL data is persisted in a Docker volume named `db_data`
- Redis data is persisted in a Docker volume named `redis_data`

### Development Setup
1. Start the databases:
```bash
docker-compose up -d
```
This will start both PostgreSQL and Redis with the default configuration.

2. Verify the databases are running:
```bash
docker-compose ps
```

3. Connect to PostgreSQL:
```bash
psql -h localhost -p 5432 -U autoban -d autoban
```

4. Connect to Redis:
```bash
redis-cli -h localhost -p 6379
```

### Database Schema
The database schema is managed through migrations (documentation coming soon). 