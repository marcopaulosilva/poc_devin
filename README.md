# Go Clean Architecture Example

This project demonstrates Go applications built using clean architecture principles:
1. Movement-speed application: Sorts and displays champions by their movement speed
2. REST API: Provides champion movement speed data via HTTP endpoints
3. Consumer application: Retrieves data from the API and persists it to AWS RDS (PostgreSQL)
4. Kubernetes deployment: Run the applications in separate Kubernetes clusters

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
  - Consumer application: Retrieves data from the API and persists it to AWS RDS (PostgreSQL)

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

Docker Compose allows you to run the movement speed API with a single command:

```bash
# Set your Riot API key in the environment
export RIOT_API_KEY=your_api_key

# Run the application
docker-compose up
```

### Running the Consumer Application with LocalStack

To run the consumer application with LocalStack for AWS RDS emulation:

```bash
# Set your Riot API key in the environment
export RIOT_API_KEY=your_api_key

# Run the local stack with consumer application
docker-compose -f docker-compose.local.yml up
```

This will:
- Start the movement speed API
- Start LocalStack with RDS emulation
- Start the consumer application that fetches data from the API and stores it in RDS


## Kubernetes Setup

This project can be deployed to Kubernetes clusters, allowing for scalable and resilient operation of both the movement speed API and the consumer application with RDS integration.

### Prerequisites for Mac

1. **Install Docker Desktop**

   Download and install [Docker Desktop for Mac](https://www.docker.com/products/docker-desktop)

2. **Enable Kubernetes in Docker Desktop**

   - Open Docker Desktop
   - Go to Preferences > Kubernetes
   - Check "Enable Kubernetes"
   - Click "Apply & Restart"

3. **Install kubectl**

   ```bash
   brew install kubectl
   ```

4. **Install Skaffold** (for local development)

   ```bash
   brew install skaffold
   ```

### Configuring the Kubernetes Deployment

1. **Create a Secret for the Riot API Key**

   ```bash
   # Replace YOUR_API_KEY with your actual Riot API key
   kubectl create secret generic riot-api-secret --from-literal=api-key=YOUR_API_KEY
   ```

   Alternatively, you can use the provided YAML file:

   ```bash
   # First, base64 encode your API key
   echo -n "YOUR_API_KEY" | base64
   
   # Edit the kubernetes/riot-api-secret.yaml file and replace the placeholder value
   # Then apply the secret
   kubectl apply -f kubernetes/riot-api-secret.yaml
   ```

2. **Deploy the API Application (First Cluster)**

   ```bash
   # Create a namespace for the API application
   kubectl create namespace api-cluster
   
   # Apply the manifests to the api-cluster namespace
   kubectl apply -f kubernetes/api-deployment.yaml -n api-cluster
   kubectl apply -f kubernetes/api-service.yaml -n api-cluster
   kubectl apply -f kubernetes/riot-api-secret.yaml -n api-cluster
   ```
   
3. **Deploy the Consumer Application with LocalStack (Second Cluster)**

   ```bash
   # Create a namespace for the consumer application
   kubectl create namespace consumer-cluster
   
   # Apply the LocalStack deployment
   kubectl apply -f kubernetes/consumer/localstack.yaml -n consumer-cluster
   
   # Apply the consumer deployment
   kubectl apply -f kubernetes/consumer/deployment.yaml -n consumer-cluster
   ```

### Running with Skaffold (Local Development)

Skaffold automates the workflow for building, pushing, and deploying your applications:

#### API Application (First Cluster)

```bash
# Make sure you're in the project root directory
skaffold dev -n api-cluster
```

#### Consumer Application (Second Cluster)

```bash
# Make sure you're in the project root directory
skaffold -f skaffold.consumer.yaml dev -n consumer-cluster
```

This will:
- Build the Docker image
- Deploy to your local Kubernetes cluster
- Stream logs from deployed pods
- Automatically redeploy when files change

### Accessing the Applications

#### Accessing the API (First Cluster)

Once deployed, you can access the API using:

```bash
# Get the service URL (if using LoadBalancer type)
kubectl get service movement-speed-api -n api-cluster

# For local development with Docker Desktop, the service is usually available at:
curl http://localhost:80/api/champions/movement-speed

# Health check
curl http://localhost:80/health
```

#### Accessing LocalStack RDS (Second Cluster)

You can interact with the LocalStack RDS instance using the AWS CLI:

```bash
# Configure AWS CLI to use LocalStack endpoint
export AWS_ENDPOINT_URL=http://localhost:4566

# List RDS instances
aws rds describe-db-instances --endpoint-url=$AWS_ENDPOINT_URL

# Connect to the PostgreSQL database (requires psql client)
psql -h localhost -p 5432 -U postgres -d champions
```

### Monitoring the Deployments

#### Monitoring the API Deployment (First Cluster)

```bash
# Check deployment status
kubectl get deployments -n api-cluster

# Check pods
kubectl get pods -n api-cluster

# View logs
kubectl logs -l app=movement-speed-api -n api-cluster

# Describe a pod for detailed information
kubectl describe pod <pod-name> -n api-cluster
```

#### Monitoring the Consumer Deployment (Second Cluster)

```bash
# Check deployment status
kubectl get deployments -n consumer-cluster

# Check pods
kubectl get pods -n consumer-cluster

# View logs for the consumer application
kubectl logs -l app=champion-consumer -n consumer-cluster

# View logs for LocalStack
kubectl logs -l app=localstack -n consumer-cluster

# Describe a pod for detailed information
kubectl describe pod <pod-name> -n consumer-cluster
```

#### Cross-Cluster Communication

The consumer application is configured to communicate with the API in the first cluster using Kubernetes DNS. The API URL is set to:

```
http://movement-speed-api.api-cluster.svc.cluster.local
```

This allows the consumer application to access the API across cluster boundaries.


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
