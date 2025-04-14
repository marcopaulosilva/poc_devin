# Go Clean Architecture Example

This project demonstrates Go applications built using clean architecture principles:
1. Movement-speed application: Sorts and displays champions by their movement speed
2. REST API: Provides champion movement speed data via HTTP endpoints

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
- Applications demonstrating League of Legends champion data retrieval:
  - Movement-speed application: Ranks champions by movement speed
  - REST API: Provides champion data through HTTP endpoints

## Prerequisites

### For Mac OS

1. Install Go
   ```bash
   # Using Homebrew
   brew install go
   
   # Verify installation
   go version  # Should show go1.18.1 or later
   ```

2. Install Docker Desktop for Mac
   - Download from [Docker Hub](https://hub.docker.com/editions/community/docker-ce-desktop-mac)
   - Follow the installation instructions

3. Obtain a Riot API Key from [Riot Developer Portal](https://developer.riotgames.com/)

## Running the Applications with Go

### Movement-Speed Application (Champions Sorted by Movement Speed)

```bash
# Set your Riot API key
export RIOT_API_KEY=your_api_key

# Run the application
go run cmd/movement-speed/main.go
```

This will fetch League of Legends champion data and display them sorted by movement speed (fastest to slowest).

### REST API (Champion Movement Speed Endpoint)

```bash
# Set your Riot API key
export RIOT_API_KEY=your_api_key

# Run the API server
go run cmd/api/main.go
```

This will start a REST API server on port 8080 with the following endpoints:

- `GET /api/champions/movement-speed`: Returns champions sorted by movement speed in JSON format
- `GET /health`: Health check endpoint that returns "OK" if the server is running

Example API response from `/api/champions/movement-speed`:

```json
{
  "count": 12,
  "champions": [
    {
      "rank": 1,
      "id": "11",
      "name": "Kassadin",
      "title": "the Void Walker",
      "movementSpeed": 355
    },
    {
      "rank": 2,
      "id": "6",
      "name": "Fizz",
      "title": "the Tidal Trickster",
      "movementSpeed": 350
    },
    ...
  ]
}
```

## Docker Setup

### Movement-Speed Application

#### Building the Docker Image

```bash
docker build -t movement-speed -f Dockerfile.movement-speed .
```

#### Running with Docker

```bash
docker run -e RIOT_API_KEY=your_api_key movement-speed
```

### REST API Application

#### Building the Docker Image

```bash
docker build -t api -f Dockerfile.api .
```

#### Running with Docker

```bash
docker run -e RIOT_API_KEY=your_api_key -p 8080:8080 api
```

### Running the Applications with Docker Compose

Docker Compose allows you to run both applications with a single command:

```bash
# Set your Riot API key in the environment
export RIOT_API_KEY=your_api_key

# Run the application
docker-compose up
```

Or pass the API key directly:

```bash
RIOT_API_KEY=your_api_key docker-compose up
```

Make sure to replace `your_api_key` with your actual Riot API key.
```bash
RIOT_API_KEY=your_api_key docker-compose up
```

Make sure to replace `your_api_key` with your actual Riot API key.

## Troubleshooting

### Common Issues on Mac OS

#### Environment Variables Not Persisting

**Problem**: Environment variables set in the terminal are not recognized by the application.

**Solution**: Set the environment variable in your shell profile:

```bash
# Add to ~/.zshrc or ~/.bash_profile
export RIOT_API_KEY=your_api_key

# Then reload your profile
source ~/.zshrc  # or source ~/.bash_profile
```

#### Docker Permission Issues

**Problem**: Docker command results in permission denied errors.

**Solution**: Make sure Docker Desktop is running. You might need to restart it:

```bash
# Check Docker status
docker info

# If it fails, restart Docker Desktop from the application menu
```

#### Architecture Differences (Apple Silicon Macs)

**Problem**: Building or running Docker images fails on M1/M2 Macs.

**Solution**: Specify the platform when building or running:

```bash
# For building
docker build --platform linux/amd64 -t movement-speed -f Dockerfile.movement-speed .

# For running
docker run --platform linux/amd64 -e RIOT_API_KEY=your_api_key movement-speed
```

#### API Key Issues

**Problem**: "RIOT_API_KEY environment variable not set" error.

**Solution**: Make sure you've set the RIOT_API_KEY environment variable correctly:

```bash
# Verify the variable is set
echo $RIOT_API_KEY

# If empty, set it
export RIOT_API_KEY=your_api_key
```
