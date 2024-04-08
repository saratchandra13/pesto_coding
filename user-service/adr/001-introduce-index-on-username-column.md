# 1. Add Index on Username

Date: 2024-04-04

## Status

Proposed

## Context

We have identified a performance issue in our application. When querying for users based on their username, the database has to scan the entire `users` table, which can be slow if the table contains a large number of records. This is because the `username` field is not indexed.

## Decision

We have decided to add an index on the `username` field in the `users` table. This will allow the database to quickly locate the data without having to scan the entire table, which will significantly improve the performance of user queries.

## Consequences

Adding an index will consume more storage space and may slightly slow down write operations (insert, update, delete) because the database will need to update the index. However, we believe that the benefits of faster query performance outweigh these costs. We will monitor the performance and adjust as necessary.