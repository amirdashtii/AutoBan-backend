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
   git clone https://github.com/amirdashtii/AutoBan.git
   cd AutoBan
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

## Usage

- Access the admin panel at `http://127.0.0.1:8000/admin/` to manage vehicles, services, and user profiles.
- Use the RESTful API to interact with the system programmatically.

## API Endpoints

### Authentication

- `POST /auth/jwt/create/` - Log in a user and obtain a token
- `POST /auth/jwt/refresh/` - Refresh the JWT token
- `POST /auth/jwt/verify/` - Verify the JWT token

### User Profiles

- `GET /profiles/me/` - Retrieve the authenticated user's profile
- `PUT /profiles/me/` - Update the authenticated user's profile

### Vehicles

- `GET /vehicles/me/` - List all vehicles for the authenticated user

### Vehicle Types

- `GET /types/` - List all vehicle types
- `GET /types/{id}/` - Retrieve a specific vehicle type

### Brands

- `GET /types/{type_id}/brands/` - List all brands for a specific vehicle type
- `GET /types/{type_id}/brands/{id}/` - Retrieve a specific brand for a specific vehicle type

### Models

- `GET /types/{type_id}/brands/{brand_id}/models/` - List all models for a specific brand and vehicle type
- `GET /types/{type_id}/brands/{brand_id}/models/{id}/` - Retrieve a specific model for a specific brand and vehicle type

## License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/amirdashtii/AutoBan/blob/main/LICENSE) file for details.
