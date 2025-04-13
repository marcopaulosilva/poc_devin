# Go Clean Architecture Example

This project demonstrates two Go applications built using clean architecture principles:
1. Main application: Displays League of Legends champions in alphabetical order
2. Movement-speed application: Sorts and displays champions by their movement speed

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
- Two applications demonstrating League of Legends champion data retrieval:
  - Main application: Lists champions alphabetically
  - Movement-speed application: Ranks champions by movement speed

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

### Main Application (Alphabetical Champions List)

```bash
# Set your Riot API key
export RIOT_API_KEY=your_api_key

# Run the application
go run cmd/app/main.go
```

This will fetch League of Legends champion data from Riot Games API and display it alphabetically in the console.

### Movement-Speed Application (Champions Sorted by Movement Speed)

```bash
# Set your Riot API key
export RIOT_API_KEY=your_api_key

# Run the application
go run cmd/movement-speed/main.go
```

This will fetch League of Legends champion data and display them sorted by movement speed (fastest to slowest).

## Docker Setup

### Main Application

#### Building the Docker Image

```bash
docker build -t poc_devin .
```

#### Running with Docker

```bash
docker run -e RIOT_API_KEY=your_api_key poc_devin
```

### Movement-Speed Application

#### Building the Docker Image

```bash
docker build -t movement-speed -f Dockerfile.movement-speed .
```

#### Running with Docker

```bash
docker run -e RIOT_API_KEY=your_api_key movement-speed
```

### Running Both Applications with Docker Compose

Docker Compose allows you to run both applications with a single command:

```bash
# Set your Riot API key in the environment
export RIOT_API_KEY=your_api_key

# Run both applications
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
docker build --platform linux/amd64 -t poc_devin .

# For running
docker run --platform linux/amd64 -e RIOT_API_KEY=your_api_key poc_devin
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
