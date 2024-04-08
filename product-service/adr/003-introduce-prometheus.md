# 1. Introduce Prometheus Metrics

Date: 2024-04-04

## Status

Implemented

## Context

We have identified a potential issue with the `product` service. Currently, the service does not expose any metrics or monitoring information, making it difficult to track the performance and health of the service. Metrics are essential for monitoring the performance, availability, and reliability of a service, as well as for troubleshooting issues and identifying bottlenecks.

## Decision

We have decided to introduce Prometheus metrics in the `product` service. Prometheus is an open-source monitoring and alerting toolkit that is widely used in the industry. By exposing metrics in a Prometheus-compatible format, we can collect, store, and visualize key performance indicators of the `product` service, such as request latency, error rates, and throughput.

## Consequences

- The `product` service will now expose Prometheus metrics in a format that can be scraped by a Prometheus server.