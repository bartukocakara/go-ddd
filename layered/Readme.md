# Layered Folder Structure Template for Go Projects
## cmd/
Contains the application's entry point (main.go). This is where the server is initialized and started.

## internal/
The core application logic resides here, divided into different layers:

### handlers/:
Handle HTTP requests and responses. They act as controllers, passing requests to services.

### services/:
Contain business logic and application rules. Services interact with repositories to fetch and manipulate data.

### repositories/:
Responsible for data access, encapsulating database queries and interactions.

### models/:
Define data structures (entities) used throughout the application.

## pkg/
Contains reusable utility functions, shared libraries, or helper code.

## configs/
Stores configuration files for the application (e.g., environment variables, database configuration).