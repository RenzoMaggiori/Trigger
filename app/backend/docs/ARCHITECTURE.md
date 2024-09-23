# Trigger Project Structure

This repository follows a microservice architecture for Go applications. Each microservice is organized to ensure scalability, modularity, and ease of maintenance.

## Folder Structure Overview
```bash
/cmd/                    # Main applications for each service
├── /service1/           # Main entry point for Service 1
└── /service2/           # Main entry point for Service 2
│
/pkg/                    # Shared libraries and utilities
└── /commonlib/          # Common packages reusable across services
│
/internal/               # Private application code for each service
├── /service1/           # Internal logic for Service 1
└── /service2/           # Internal logic for Service 2
│
/api/                    # API definitions (OpenAPI/GRPC/Protobuf)
├── /service1/           # API definitions for Service 1
└── /service2/           # API definitions for Service 2
│
/config/                 # Configuration files
└── config.yaml          # Global or service-specific config files
│
/deploy/                 # Deployment-related files (Docker, Kubernetes, CI/CD)
├── /docker/             # Docker configuration files (Dockerfiles, docker-compose)
└── /k8s/                # Kubernetes deployment configurations
│
/scripts/                # Automation scripts
└── build.sh             # Example: Script for building the project
│
/docs/                   # Project documentation
└── architecture.md      # Architecture overview and documentation
│
/tests/                  # Unit and integration tests
├── /service1/           # Test files for Service 1
└── /service2/           # Test files for Service 2
```

## Directory Breakdown

`/cmd/`

This directory contains the entry point for each microservice. Every service has its own subdirectory under cmd/ that holds the main.go file, which is responsible for bootstrapping the service.

Example:

```bash
/cmd/service1/main.go
```

`/pkg/`

pkg/ is where you place shared libraries and utilities. These are packages that can be imported by other services within the project or by external projects.

Example:

```bash
/pkg/lib/logger.go  # A shared logging utility
```

`/internal/`

The internal/ directory contains the service-specific internal logic that should not be exposed to or imported by other services outside the repository. Go restricts import access to anything under the internal/ directory.

Example:

```bash
/internal/service1/handler.go  # Business logic for Service 1
```

`/api/`

The api/ directory contains the API definitions for each service. These could be OpenAPI (Swagger) specs, gRPC protobuf files, or any other service communication definitions.

Example:

```bash
/api/service1/openapi.json   # Swagger definition for Service 1
```

`/config/`

Centralized configuration files are placed in the config/ directory. These may include .yaml or .json files for service-specific configurations or global configurations shared across services.

Example:

```bash
/config/config.yaml  # Example global configuration file
```

`/deploy/`

The deploy/ directory contains Docker, Kubernetes, and CI/CD pipeline configuration files for deploying your services. This helps in containerizing and orchestrating services for production.

Example:

```bash
/deploy/docker/Dockerfile             # Dockerfile for Service 1
/deploy/k8s/service1-deployment.yaml  # Kubernetes deployment manifest for Service 1
```

`/scripts/`

Automation scripts for common tasks such as building, testing, or deploying services. This directory helps streamline the development and CI/CD process.

Example:

```bash
/scripts/build.sh  # Example build script
```

`/docs/`

The docs/ directory contains any project-related documentation. This is where architecture overviews, API documentation, or any other relevant docs live.

Example:

```bash
/docs/ARCITECTURE.md  # Architecture documentation for the project
```

`/tests/`

This directory contains unit and integration test files, usually separated by service. Having a dedicated tests/ folder helps to organize the test cases and test data for each service.

Example:

```bash
/tests/service1/handler_test.go  # Unit tests for Service 1
```

