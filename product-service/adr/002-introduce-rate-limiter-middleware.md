# 1. Introduce Rate Limiter Middleware

Date: 2024-04-04

## Status

Implemented

## Context

We have identified a potential issue with the `product` service. Currently, the service does not have any rate limiting in place, which can lead to abuse by malicious users or unintentional misuse by legitimate users. Rate limiting is a common technique used to prevent abuse of APIs and services by limiting the number of requests a user can make in a given time period.

## Decision

We have decided to introduce a rate limiter middleware in the `product` service. This middleware will limit the number of requests a user can make to the service within a specified time window. If a user exceeds the limit, the middleware will return a `429 Too Many Requests` response, indicating that the user has exceeded the rate limit.

## Consequences

- Users of the `product` service will be limited in the number of requests they can make within a specified time window.