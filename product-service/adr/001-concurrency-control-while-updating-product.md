# 1. Update product function to use optimistic concurrency control

Date: 2024-04-04

## Status

Implemented

## Context

We have identified a potential issue with the `updateProduct` function in the `product` service. Currently, the function updates a product in the database without checking if the product has been modified by another process since it was last read. This can lead to lost updates or inconsistent data if two processes try to update the same product at the same time.

## Decision

We have decided to update the `updateProduct` function to use optimistic concurrency control. This means that before updating the product in the database, the function will check if the product has been modified since it was last read. If the product has been modified, the update will fail, and the function will return an error. This will prevent lost updates and ensure data consistency.

## Consequences

- The `updateProduct` function will now check for concurrent modifications before updating the product in the database.