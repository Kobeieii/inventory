# Inventory Management System

This is a CRUD (Create, Read, Update, Delete) project built using Go with Hexagonal Architecture. It provides an inventory management system that follows clean architecture principles to maintain separation of concerns.

## Project Structure

```
inventory/
┣ cmd/
┃ ┗ main.go                 # Entry point of the application
┣ internal/
┃ ┣ app/
┃ ┃ ┣ adapters/
┃ ┃ ┃ ┣ handlers/
┃ ┃ ┃ ┃ ┗ http.go          # HTTP handlers for API endpoints
┃ ┃ ┃ ┗ repositories/
┃ ┃ ┃   ┗ gorm.go         # GORM-based repository implementation
┃ ┃ ┣ application/
┃ ┃ ┃ ┗ services/
┃ ┃ ┃   ┗ product.go      # Business logic for product operations
┃ ┃ ┗ domain/
┃ ┃   ┣ entities/
┃ ┃ ┃ ┃ ┗ product.go     # Domain entity definition
┃ ┃   ┣ ports/
┃ ┃ ┃ ┃ ┣ product_repository.go # Interface for product repository
┃ ┃ ┃ ┃ ┗ product_service.go    # Interface for product service
┃ ┃   ┗ errors.go        # Custom error handling
┃ ┣ infrastructure/
┃ ┃ ┗ database/
┃ ┃   ┗ database.go      # Database configuration and initialization
┃ ┗ utils/
┃   ┗ map.go            # Utility functions
┣ .env                 # Environment variables
┣ .gitignore           # Git ignore rules
┣ docker-compose.yaml  # Docker configuration for dependencies
┣ env_example.txt      # Example environment variables
┣ go.mod               # Go module dependencies
┗ go.sum               # Go module checksum
```

## Features
- **Hexagonal Architecture**: Follows clean architecture principles.
- **CRUD Operations**: Supports Create, Read, Update, and Delete operations for products.
- **GORM Integration**: Uses GORM as the ORM for database interactions.
- **REST API**: Provides HTTP endpoints for managing inventory.
- **Docker Support**: Includes Docker Compose for database setup.

## Installation & Setup

### Prerequisites
- Go 1.18+
- Docker & Docker Compose
- PostgreSQL or any supported database

### Steps to Run
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd inventory
   ```

2. Copy the example environment variables:
   ```sh
   cp env_example.txt .env
   ```

3. Start the database using Docker:
   ```sh
   docker-compose up -d
   ```

4. Run the application:
   ```sh
   go run cmd/main.go
   ```

## API Endpoints
| Method | Endpoint        | Description        |
|--------|---------------|--------------------|
| GET    | /products      | Get all products  |
| GET    | /products/{id} | Get product by ID |
| POST   | /products      | Create a product  |
| PUT    | /products/{id} | Update a product  |
| DELETE | /products/{id} | Delete a product  |

## Contributing
Feel free to fork this repository and submit pull requests. Contributions are welcome!

## License
This project is open-source and available under the [MIT License](LICENSE).

