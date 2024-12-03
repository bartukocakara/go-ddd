# 📦 Layer Descriptions
## cmd/
Contains the main entry point for the application. The main.go initializes and runs the server.

## internal/
Houses the core business logic, organized into distinct domains:

## user/

handler.go: Handles HTTP requests related to users.
service.go: Contains business logic for user operations.
repository.go: Handles data access for the user domain.
user.go: Defines the User model and its attributes.
product/

Similar structure as the user domain, focusing on product-related functionalities.
## pkg/
Contains shared utilities, helpers, and reusable components.

## configs/
Stores configuration files (e.g., environment variables, database connections).

🛠️ Setup Instructions
Prerequisites
Go 1.18 or higher
Docker (optional for database setup)