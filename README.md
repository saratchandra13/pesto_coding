# Project Overview

This project consists of three microservices implemented in Go: Order Service, Product Service, and User Service. Each of these services is designed to be deployed in a cloud environment such as AWS or Google Cloud.

## Architecture Decision Records (ADRs)

Each service has its own set of Architecture Decision Records (ADRs) that document the important decisions made during the development of the service. These ADRs can be found in the `adr` directory inside each service's repository.

## Testing

Each service comes with a suite of tests that can be run to verify its functionality. The instructions for running these tests, as well as example cURL commands for interacting with the service, can be found in the README file of each service's repository.

## Rate Limiting

Rate limiting has been implemented in each service to prevent abuse. This feature limits the number of requests a user can make within a specified time window. Tests for the rate limiting functionality are included in the test suite.

## Role Controller

Each service has its own Role Controller for handling roles. This could have been implemented as a separate service, but due to time constraints, it was included in each service.

## Concurrency Control

For concurrency control, optimistic concurrency control has been used in the Product Service. This helps to ensure that the service can handle multiple concurrent requests without data conflicts.

## Database

MySQL has been used to store all data for the services. The database models for each service can be found in the `models` directory inside each service's repository. Currently, the database credentials are stored in files, but in a production environment, these would be stored in environment variables. Each service could potentially have its own separate database instance.

## Monitoring and Alerting

Prometheus has been integrated into the Product Service for monitoring and alerting. Prometheus is an open-source monitoring and alerting toolkit that is widely used in the industry. By exposing metrics in a Prometheus-compatible format, we can collect, store, and visualize key performance indicators of the `product` service, such as request latency, error rates, and throughput.

# Deployment

Each of these services comes with a Dockerfile, allowing you to build a Docker image for each service. These Docker images can then be deployed to a Kubernetes cluster or an AWS ECS cluster.

To deploy these services, follow the steps below:

1. Clone the repositories for the Order Service, Product Service, and User Service.

2. Build the Docker images for each service using the `docker build` command.

3. Push the Docker images to a Docker registry that your Kubernetes cluster or ECS cluster can access.

4. Create a Kubernetes Deployment or an ECS Task Definition for each service, referencing the Docker images you pushed to your Docker registry.

5. Deploy each service to your Kubernetes cluster or ECS cluster.

Please refer to the specific cloud provider's documentation for more detailed instructions on deploying Docker images.

# Note

This README provides a high-level overview of the project. For more detailed information, please refer to the README and ADRs in each service's repository.