# AutoBan 🚗

Vehicle Maintenance Management System

## About
AutoBan is a vehicle maintenance management system built with Go and PostgreSQL. It allows users to manage their vehicle maintenance records, service schedules, and vehicle catalog efficiently.

## Technologies
- Go
- PostgreSQL
- Docker & Docker Compose
- JWT for Authentication
- Redis for session and cache

## Prerequisites
- Go 1.x or higher
- Docker and Docker Compose
- An `.env` file (you can copy from `.env.example`)

## Quick Start

1. **Clone the repository:**
   ```bash
   git clone https://github.com/yourusername/AutoBan.git
   cd AutoBan
   ```

2. **Copy the environment file:**
   ```bash
   cp .env.example .env
   ```
   Edit `.env` with your database and Redis settings.

3. **Run the services with Docker:**
   ```bash
   docker-compose up -d
   ```

4. **Run the application:**
   ```bash
   go run cmd/app/main.go
   ```

---

## API Endpoints

### Authentication
- `POST   /api/v1/auth/register` - Register a new user
- `POST   /api/v1/auth/login` - User login
- `POST   /api/v1/auth/refresh-token` - Refresh access token
- `POST   /api/v1/auth/send-verifycode` - Send verify code
- `POST   /api/v1/auth/active` - active user
- `POST   /api/v1/auth/logout` - Logout (requires token)
- `POST   /api/v1/auth/logout-all` - Logout from all devices (requires token)
- `GET    /api/v1/auth/sessions` - List active sessions (requires token)

### User Profile
- `GET    /api/v1/users/me` - Get user profile (requires token)
- `PUT    /api/v1/users/me` - Update profile (requires token)
- `PUT    /api/v1/users/me/change-password` - Change password (requires token)
- `DELETE /api/v1/users/me` - Delete account (requires token)

### Vehicle Catalog (Public)
- `GET    /api/v1/vehicles/hierarchy` - Get full vehicle hierarchy (types, brands, models, generations)
- `GET    /api/v1/vehicles/types` - List vehicle types
- `GET    /api/v1/vehicles/types/{type_id}` - Get vehicle type details
- `GET    /api/v1/vehicles/types/{type_id}/brands` - List brands for a type
- `GET    /api/v1/vehicles/types/{type_id}/brands/{brand_id}` - Get brand details
- `GET    /api/v1/vehicles/types/{type_id}/brands/{brand_id}/models` - List models for a brand
- `GET    /api/v1/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}` - Get model details
- `GET    /api/v1/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations` - List generations for a model
- `GET    /api/v1/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id}` - Get generation details

### User Vehicles (Requires Token)
- `POST   /api/v1/user/vehicles` - Add a vehicle to user
- `GET    /api/v1/user/vehicles` - List user vehicles
- `GET    /api/v1/user/vehicles/{vehicle_id}` - Get user vehicle details
- `PUT    /api/v1/user/vehicles/{vehicle_id}` - Update user vehicle
- `DELETE /api/v1/user/vehicles/{vehicle_id}` - Delete user vehicle

### Service Visits (Requires Token)
- `GET    /api/v1/user/vehicles/{vehicle_id}/service-visits` - List service visits
- `POST   /api/v1/user/vehicles/{vehicle_id}/service-visits` - Add a service visit
- `GET    /api/v1/user/vehicles/{vehicle_id}/service-visits/last` - Get last service visit
- `GET    /api/v1/user/vehicles/{vehicle_id}/service-visits/{visit_id}` - Get service visit details
- `PUT    /api/v1/user/vehicles/{vehicle_id}/service-visits/{visit_id}` - Update service visit
- `DELETE /api/v1/user/vehicles/{vehicle_id}/service-visits/{visit_id}` - Delete service visit

#### Oil Changes & Oil Filters
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-changes` - List oil changes
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-changes/last` - Last oil change
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-changes/{oil_change_id}` - Oil change details
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-filters` - List oil filter changes
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-filters/last` - Last oil filter change
- `GET    /api/v1/user/vehicles/{vehicle_id}/oil-filters/{oil_filter_id}` - Oil filter change details

### Admin - User Management (Requires Admin Token)
- `GET    /api/v1/admin/users` - List users
- `GET    /api/v1/admin/users/{id}` - Get user details
- `PUT    /api/v1/admin/users/{id}` - Update user
- `POST   /api/v1/admin/users/{id}/role` - Change user role
- `POST   /api/v1/admin/users/{id}/status` - Change user status
- `POST   /api/v1/admin/users/{id}/change-password` - Change user password
- `DELETE /api/v1/admin/users/{id}` - Delete user

### Admin - Vehicle Catalog Management (Requires Admin Token)
- `POST   /api/v1/admin/vehicles/types` - Create vehicle type
- `PUT    /api/v1/admin/vehicles/types/{type_id}` - Update vehicle type
- `DELETE /api/v1/admin/vehicles/types/{type_id}` - Delete vehicle type
- `POST   /api/v1/admin/vehicles/types/{type_id}/brands` - Create brand
- `PUT    /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}` - Update brand
- `DELETE /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}` - Delete brand
- `POST   /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models` - Create model
- `PUT    /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}` - Update model
- `DELETE /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}` - Delete model
- `POST   /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations` - Create generation
- `PUT    /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id}` - Update generation
- `DELETE /api/v1/admin/vehicles/types/{type_id}/brands/{brand_id}/models/{model_id}/generations/{generation_id}` - Delete generation

---

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

---

## Technical Notes
- JWT authentication, session management and caching with Redis
- Full vehicle hierarchy is cached in Redis for fast access
- Uses GORM and PostgreSQL for database
- Clean, maintainable architecture
- API documentation with Swagger (if enabled)

---

## Contributing
Contributions are welcome! Please submit a Pull Request.

## License
This project is licensed under the MIT License. See the LICENSE file for more details. 