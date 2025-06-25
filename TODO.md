# TODO List for AutoBan Project

## Auth Use Case
1. [x] Implement login functionality
2. [x] Implement register functionality
3. [x] Implement logout functionality
4. [x] Implement refresh token functionality

## User Use Case
5. [x] Implement get user functionality
6. [x] Implement update user functionality
7. [x] Implement change password functionality
8. [x] Implement delete user functionality

## Admin Use Case
9. [x] Implement list users functionality
10. [x] Implement view other users functionality
11. [x] Implement additional admin operations

## General
11. [ ] Write unit tests for user use case
12. [ ] Write unit tests for auth use case
13. [ ] Write unit tests for admin use case
14. [ ] Review and refactor code for best practices

## Database Setup
15. [x] Set up PostgreSQL using Docker Compose
16. [x] Create a single instance of PostgreSQL for development
17. [x] Ensure the database setup is easy to test and mock
18. [ ] Write integration tests for database interactions
19. [x] Document the database setup and usage in the README file

## API Endpoints
20. [x] Create endpoints for user registration
21. [x] Create endpoints for user login
22. [x] Create endpoints for user management 
23. [x] Create endpoints for admin management 

## Vehicle Management
1. [x] Implement vehicle type management
   - [x] Add vehicle types API
   - [x] List vehicle types API
2. [x] Implement vehicle brand management
   - [x] Add vehicle brands API
   - [x] List vehicle brands API
   - [x] Filter brands by type API
3. [x] Implement vehicle model management
   - [x] Add vehicle models API
   - [x] List vehicle models API
   - [x] Filter models by brand API
4. [x] Implement vehicle generation management
   - [x] Add vehicle generations API
   - [x] List vehicle generations API
   - [x] Filter generations by model API
5. [x] Implement user vehicle management
   - [x] Add user vehicle API
   - [x] List user vehicles API
   - [x] Update user vehicle API
   - [x] Delete user vehicle API
6. [ ] Add initial data for common vehicle types, brands, and models
   - [ ] Add Iranian manufacturers and models
   - [ ] Add common foreign manufacturers and models