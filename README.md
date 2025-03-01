# Inventory Management System (Hexagonal Architecture)

This project is a simple CRUD application built with Go using Hexagonal Architecture. It was created as a practice project to understand and implement the principles of Hexagonal Architecture in a clean and structured manner.

## Features
- CRUD operations for `Product` entity
- Separation of concerns using Hexagonal Architecture
- Uses `Product`, `ProductDTO`, and `ProductModel` to maintain clear boundaries
- Implements `GORM` as the repository layer
- Uses `Fiber` as the web framework for handling HTTP requests
- Provides an HTTP handler for product operations

## Project Structure
```
inventory/
┣ cmd/
┃ ┗ main.go                    # Application entry point
┣ internal/
┃ ┣ app/
┃ ┃ ┣ adapters/                 # Infrastructure and framework-specific implementations
┃ ┃ ┃ ┣ handlers/
┃ ┃ ┃ ┃ ┗ http.go              # HTTP handlers for Product entity (Primary Adapter)
┃ ┃ ┃ ┗ repositories/
┃ ┃ ┃   ┗ gorm.go              # GORM-based repository implementation (Secondary Adapter)
┃ ┃ ┣ application/
┃ ┃ ┃ ┗ services/
┃ ┃ ┃   ┗ product.go           # Business logic for Product entity (Application Core)
┃ ┃ ┗ domain/
┃ ┃   ┣ entities/
┃ ┃ ┃ ┃ ┗ product.go          # Product domain model
┃ ┃   ┣ ports/
┃ ┃ ┃ ┃ ┣ product_repository.go # Repository port (Secondary Port)
┃ ┃ ┃ ┃ ┗ product_service.go   # Service port (Primary Port)
┃ ┃   ┗ errors.go              # Error handling
┃ ┣ infrastructure/
┃ ┃ ┗ database/
┃ ┃   ┗ database.go            # Database connection and initialization
┃ ┗ utils/
┃   ┗ map.go                   # Utility functions
┣ .env                          # Environment variables
┣ .gitignore                    # Git ignore file
┣ docker-compose.yaml           # Docker Compose configuration
┣ env_example.txt               # Example environment variables file
┣ go.mod                         # Go module definition
┗ go.sum                         # Go module dependencies
```

## Technologies Used
- **Go** (Golang)
- **Fiber** (Web framework for HTTP handling)
- **GORM** (ORM for database operations)
- **Hexagonal Architecture** (Ports and Adapters pattern)
- **Docker** (for containerization)

## Architecture Flow
This project follows Hexagonal Architecture to separate concerns and make the application more maintainable and testable. The flow of interactions is structured as follows:

![Untitled Diagram drawio](https://github.com/user-attachments/assets/6c575882-06d8-4b7f-a56e-900b47408494)



## Getting Started
### Prerequisites
- Go (1.18+ recommended)
- Docker (optional, for running with containerization)

### Installation & Setup
1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/inventory.git
   cd inventory
   ```
2. Copy the environment example file:
   ```sh
   cp env_example.txt .env
   ```
3. Update `.env` file with database configuration.
4. Run the application:
   ```sh
   go run cmd/main.go
   ```

### Running with Docker
To run the application using Docker, use the following command:
```sh
docker-compose up --build
```

## API Endpoints
| Method | Endpoint       | Description       |
|--------|--------------|------------------|
| GET    | `/products`  | Get all products |
| GET    | `/products/{id}` | Get product by ID |
| POST   | `/products`  | Create a new product |
| PUT    | `/products/{id}` | Update a product |
| DELETE | `/products/{id}` | Delete a product |

## License
This project is open-source and available under the [MIT License](LICENSE).

## Author
Developed by **Kobeieii** for practicing Hexagonal Architecture in Go.

