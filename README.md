# Go Clean Architecture Example

This project demonstrates a Go application built using clean architecture principles. It makes HTTP requests to an external API and displays the results in the console.

## Project Structure

The project follows clean architecture with the following layers:

- **Domain Layer**: Contains business logic, entities, and use cases
  - `entities`: Core business models
  - `usecases`: Application business rules

- **Interface Layer**: Adapters that convert data between the domain and external services
  - `http`: HTTP-specific implementations of repositories

- **Infrastructure Layer**: External frameworks and tools
  - `client`: HTTP client implementation
  - `logger`: Logging utilities

## Features

- Clean architecture implementation
- HTTP client for making external API requests
- Formatted console output with colored logging
- Example application demonstrating user data retrieval

## Running the Application

```bash
go run cmd/app/main.go
```

This will fetch user data from JSONPlaceholder API and display it in the console.
