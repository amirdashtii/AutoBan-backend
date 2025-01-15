# AutoBan

AutoBan is a Django-based application for managing vehicles, services, and user profiles. This project provides a comprehensive system for vehicle management, including vehicle types, brands, models, and services such as oil changes. It also includes user authentication and profile management.

## Features

- User authentication and profile management
- Vehicle management (types, brands, models, and vehicles)
- Service management (including oil changes)
- Admin panel for managing all entities
- RESTful API for interacting with the system

## Installation

### Using Docker for Database

1. Clone the repository:

   ```bash
   git clone https://github.com/amirdashtii/AutoBan-backend.git
   cd AutoBan-backend
   ```

2. Create a .env file based on the provided .env.example:

   ```bash
   cp .env.example .env
   ```

3. Start the database container:

   ```bash
   docker-compose up -d
   ```

4. Create and activate a virtual environment:

   ```bash
   python3 -m venv venv
   source venv/bin/activate
   ```

5. Install the dependencies:

   ```bash
   pip install -r requirements.txt
   ```

6. Apply the migrations:

   ```bash
   python manage.py migrate
   ```

7. Create a superuser:

   ```bash
   python manage.py createsuperuser
   ```

8. Run the development server:

   ```bash
   python manage.py runserver
   ```

## Importing Initial Data

To import initial data for vehicle types, brands, and models from CSV files, use the following command:

   ```bash
   python manage.py import_data --types data/types.csv --brands data/brands.csv --models data/models.csv
   ```

## Usage

- Access the admin panel at `http://127.0.0.1:8000/admin/` to manage vehicles, services, and user profiles.
- Use the RESTful API to interact with the system programmatically.

## API Endpoints

### user

- `GET /api/auth/users/` - List all users (admin only)
- `POST /api/auth/users/` - Create a new user
- `GET /api/auth/users/me/` - Retrieve the authenticated user's profile
- `PUT /api/auth/users/me/` - Update the authenticated user's profile
- `DELETE /api/auth/users/me/` - Delete the authenticated user's profile
- `GET /api/auth/users/{id}/` - Retrieve a specific user (admin only)
- `PUT /api/auth/users/{id}/` - Update a specific user (admin only)
- `DELETE /api/auth/users/{id}/` - Delete a specific user (admin only)

### Authentication

- `POST /api/auth/jwt/create/` - Log in a user and obtain a token
- `POST /api/auth/jwt/refresh/` - Refresh the JWT token
- `POST /api/auth/jwt/verify/` - Verify the JWT token

### User Profiles

- `GET /api/profiles/me/` - Retrieve the authenticated user's profile
- `PUT /api/profiles/me/` - Update the authenticated user's profile

### Vehicles

- `GET /api/vehicles/me/` - List all vehicles for the authenticated user

### Vehicle Types

- `GET /api/types/` - List all vehicle types
- `GET /api/types/{id}/` - Retrieve a specific vehicle type

### Brands

- `GET /api/types/{type_id}/brands/` - List all brands for a specific vehicle type
- `GET /api/types/{type_id}/brands/{id}/` - Retrieve a specific brand for a specific vehicle type

### Models

- `GET /api/types/{type_id}/brands/{brand_id}/models/` - List all models for a specific brand and vehicle type
- `GET /api/types/{type_id}/brands/{brand_id}/models/{id}/` - Retrieve a specific model for a specific brand and vehicle type

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/amirdashtii/AutoBan-backend/blob/main/LICENSE) file for details.
