# Consumer Application with PostgreSQL Architecture

This document provides a visual representation of the consumer application architecture with PostgreSQL integration, showing the relationships between components and the flow of data through the system.

## Multi-Cluster Architecture Overview

The project uses two Kubernetes clusters with the following components:

```mermaid
graph TD
    subgraph "API Cluster"
        A[Movement Speed API Deployment]
        B[API Service]
        C[Riot API Secret]
    end
    
    subgraph "Consumer Cluster"
        D[Champion Consumer Deployment]
        E[PostgreSQL Deployment]
        F[PostgreSQL Service]
    end
    
    subgraph "External"
        G[Riot Games API]
    end
    
    A --> C
    A --> G
    B --> A
    D --> B
    D --> F
    F --> E
```

## Clean Architecture Overview

The consumer application follows clean architecture principles with distinct layers:

```mermaid
graph TD
    subgraph "Application Layer"
        A[cmd/consumer/main.go]
    end
    
    subgraph "Domain Layer"
        B[entities/champion_record.go]
        C[repositories/champion_repository.go]
    end
    
    subgraph "Interface Layer"
        D[api/movement_speed_client.go]
        E[db/postgres_champion_repository.go]
    end
    
    subgraph "Infrastructure Layer"
        F[client/http_client.go]
        G[db/postgres.go]
        H[logger/logger.go]
    end
    
    A --> B
    A --> C
    A --> D
    A --> E
    A --> G
    A --> H
    D --> F
    D --> B
    E --> B
    E --> C
    E --> G
    E --> H
```

## Data Flow

The following diagram illustrates the flow of data through the system:

```mermaid
sequenceDiagram
    participant Consumer
    participant API Client
    participant Movement Speed API
    participant Riot API
    participant Champion Repository
    participant PostgreSQL
    
    Consumer->>API Client: GetChampionsByMovementSpeed()
    API Client->>Movement Speed API: GET /api/champions/movement-speed
    Movement Speed API->>Riot API: Fetch Champion Data
    Riot API-->>Movement Speed API: Champion Data
    Movement Speed API-->>API Client: Sorted Champions JSON
    API Client-->>Consumer: Champion Records
    Consumer->>Champion Repository: SaveChampions()
    Champion Repository->>PostgreSQL: INSERT/UPDATE Champions
    PostgreSQL-->>Champion Repository: Success
    Champion Repository-->>Consumer: Success
```

## PostgreSQL Database Integration

The following diagram shows the PostgreSQL database integration for the consumer application:

```mermaid
graph TD
    subgraph "Kubernetes Environment"
        A[Consumer Application]
        B[PostgreSQL Container]
        
        subgraph "PostgreSQL"
            C[PostgreSQL Engine]
            D[Champions Database]
            E[Champions Table]
        end
    end
    
    A -->|"SQL Queries"| B
    B --> C
    C --> D
    D --> E
```

## Component Descriptions

### Application Layer
- **cmd/consumer/main.go**: Entry point that initializes components and runs the synchronization process

### Domain Layer
- **entities/champion_record.go**: Data structures representing champion records in the database
- **repositories/champion_repository.go**: Interface for database operations

### Interface Layer
- **api/movement_speed_client.go**: Client for retrieving champion data from the movement speed API
- **db/postgres_champion_repository.go**: PostgreSQL implementation of the champion repository

### Infrastructure Layer
- **client/http_client.go**: HTTP client implementation for API requests
- **db/postgres.go**: PostgreSQL connection and initialization
- **logger/logger.go**: Logging functionality with different severity levels

### Kubernetes Components
- **kubernetes/consumer/deployment.yaml**: Deployment configuration for the consumer application
- **kubernetes/consumer/postgres.yaml**: Deployment and service configuration for PostgreSQL
