    # Docker Compose Watch App

This application is designed to monitor Docker Compose services and is built using Golang. The project runs on port 8080.

## Prerequisites

- Docker
- Docker Compose
- Golang

## Installation

1. Clone the repository:

2. Ensure Docker and Docker Compose are installed and running on your machine.

## Running the Application

To run the application, use the provided Makefile commands.

### Start the Services

To start the Docker Compose services and begin watching:

```sh
make up
```

### Stop the Services

To stop the Docker Compose services:

```sh
make down
```

### Accessing the Application

Once the services are running, the application can be accessed at:

```sh
http://localhost:8080
```

### Make changes

Make changes to main.go and save the file. Docker Compose will automatically rebuild the image, restart the container, and apply the changes.

Refresh the browser to see the updated response.
