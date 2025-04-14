# POC_DEVIN Architecture Diagram

This document provides a visual representation of the POC_DEVIN project architecture, showing the relationships between components and the flow of data through the system.

## Clean Architecture Overview

The project follows clean architecture principles with distinct layers:

```mermaid
graph TD
    subgraph "Application Layer"
        A[cmd/api/main.go]
        B[cmd/movement-speed/main.go]
    end
    
    subgraph "Domain Layer"
        C[entities/champion.go]
        D[usecases/champion_usecase.go]
    end
    
    subgraph "Interface Layer"
        E[http/champion_repository.go]
        F[api/movement_speed_handler.go]
        G[api/server.go]
    end
    
    subgraph "Infrastructure Layer"
        H[client/http_client.go]
        I[logger/logger.go]
    end
    
    A --> D
    A --> G
    A --> F
    B --> D
    D --> C
    D --> E
    E --> H
    F --> D
    A --> I
    B --> I
    E --> I
    F --> I
    G --> I
```

## Kubernetes Deployment Architecture

The application is deployed to Kubernetes with the following components:

```mermaid
graph TD
    subgraph "Kubernetes Cluster"
        A[API Deployment]
        B[API Service]
        C[Riot API Secret]
    end
    
    subgraph "External"
        D[Riot Games API]
        E[Client/User]
    end
    
    A --> C
    A --> D
    B --> A
    E --> B
```

## Data Flow

The following diagram illustrates the flow of data through the system:

```mermaid
sequenceDiagram
    participant User
    participant API Service
    participant Movement Speed Handler
    participant Champion UseCase
    participant Champion Repository
    participant HTTP Client
    participant Riot API
    
    User->>API Service: GET /api/champions/movement-speed
    API Service->>Movement Speed Handler: Handle Request
    Movement Speed Handler->>Champion UseCase: GetAllChampions()
    Champion UseCase->>Champion Repository: GetAllChampions()
    Champion Repository->>HTTP Client: GET Request
    HTTP Client->>Riot API: HTTP Request
    Riot API-->>HTTP Client: Champion Data
    HTTP Client-->>Champion Repository: Response
    Champion Repository-->>Champion UseCase: Champions
    Champion UseCase-->>Movement Speed Handler: Champions
    Movement Speed Handler->>Movement Speed Handler: Sort by Movement Speed
    Movement Speed Handler-->>API Service: JSON Response
    API Service-->>User: Sorted Champions
```

## Component Descriptions

### Application Layer
- **cmd/api/main.go**: REST API entry point that initializes components and serves HTTP endpoints
- **cmd/movement-speed/main.go**: CLI application that displays champions sorted by movement speed

### Domain Layer
- **entities/champion.go**: Data structures representing League of Legends champions
- **usecases/champion_usecase.go**: Business logic for retrieving and processing champion data

### Interface Layer
- **http/champion_repository.go**: Adapter for retrieving champion data from Riot API
- **api/movement_speed_handler.go**: HTTP handler for the movement speed endpoint
- **api/server.go**: HTTP server implementation with graceful shutdown

### Infrastructure Layer
- **client/http_client.go**: HTTP client implementation for external API requests
- **logger/logger.go**: Logging functionality with different severity levels

### Kubernetes Components
- **kubernetes/api-deployment.yaml**: Deployment configuration for the API
- **kubernetes/api-service.yaml**: Service configuration to expose the API
- **kubernetes/riot-api-secret.yaml**: Secret for storing the Riot API key
