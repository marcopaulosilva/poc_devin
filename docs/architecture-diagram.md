# POC_DEVIN Architecture Diagram

This document provides a visual representation of the POC_DEVIN project architecture, showing the relationships between components and the flow of data through the system.

## Technology Stack

The project uses the following technologies:

```mermaid
flowchart LR
    subgraph "Backend"
        A["<img src='https://golang.org/doc/gopher/frontpage.png' width='50' height='50' /><br/>Go"]
    end
    
    subgraph "Containerization"
        B["<img src='https://www.docker.com/wp-content/uploads/2022/03/Moby-logo.png' width='50' height='50' /><br/>Docker"]
        C["<img src='https://kubernetes.io/images/favicon.png' width='50' height='50' /><br/>Kubernetes"]
    end
    
    subgraph "External APIs"
        D["<img src='https://developer.riotgames.com/static/img/riot-api-landing.png' width='50' height='50' /><br/>Riot Games API"]
    end
    
    A --- B
    B --- C
    A --- D
```

## Clean Architecture Overview

The project follows clean architecture principles with distinct layers:

```mermaid
graph TD
    subgraph "Application Layer" 
        A["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>cmd/api/main.go"]
        B["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>cmd/movement-speed/main.go"]
    end
    
    subgraph "Domain Layer"
        C["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>entities/champion.go"]
        D["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>usecases/champion_usecase.go"]
    end
    
    subgraph "Interface Layer"
        E["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>http/champion_repository.go"]
        F["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>api/movement_speed_handler.go"]
        G["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>api/server.go"]
    end
    
    subgraph "Infrastructure Layer"
        H["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>client/http_client.go"]
        I["<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>logger/logger.go"]
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
        A["<img src='https://kubernetes.io/images/favicon.png' width='20' height='20' /><br/>API Deployment"]
        B["<img src='https://kubernetes.io/images/favicon.png' width='20' height='20' /><br/>API Service"]
        C["<img src='https://kubernetes.io/images/favicon.png' width='20' height='20' /><br/>Riot API Secret"]
    end
    
    subgraph "External"
        D["<img src='https://developer.riotgames.com/static/img/riot-api-landing.png' width='20' height='20' /><br/>Riot Games API"]
        E["<img src='https://www.svgrepo.com/show/286690/user.svg' width='20' height='20' /><br/>Client/User"]
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
    participant User as "<img src='https://www.svgrepo.com/show/286690/user.svg' width='20' height='20' /><br/>User"
    participant API as "<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>API Service"
    participant Handler as "<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>Movement Speed Handler"
    participant UseCase as "<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>Champion UseCase"
    participant Repo as "<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>Champion Repository"
    participant Client as "<img src='https://golang.org/doc/gopher/frontpage.png' width='20' height='20' /><br/>HTTP Client"
    participant RiotAPI as "<img src='https://developer.riotgames.com/static/img/riot-api-landing.png' width='20' height='20' /><br/>Riot API"
    
    User->>API: GET /api/champions/movement-speed
    API->>Handler: Handle Request
    Handler->>UseCase: GetAllChampions()
    UseCase->>Repo: GetAllChampions()
    Repo->>Client: GET Request
    Client->>RiotAPI: HTTP Request
    RiotAPI-->>Client: Champion Data
    Client-->>Repo: Response
    Repo-->>UseCase: Champions
    UseCase-->>Handler: Champions
    Handler->>Handler: Sort by Movement Speed
    Handler-->>API: JSON Response
    API-->>User: Sorted Champions
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
