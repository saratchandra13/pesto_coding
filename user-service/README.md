# pesto_coding
Coding file for pesto interviews

## API Testing

You can test the API using the following `curl` commands:

### Create a new user

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "Username": "testuser",
    "Email": "testuser@example.com",
    "Password": "password123"
}' http://localhost:8080/users

```

### Authenticate a user

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "Username": "testuser",
    "Password": "password123"
}' http://localhost:8080/users/auth
```

### Get user details

```bash
curl -X GET http://localhost:8080/users/1
```

### Get role details

```bash
curl -X GET http://localhost:8080/roles/1
```

### create role

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "RoleName": "admin"
}' http://localhost:8080/roles
```

### Assign role to user

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "UserId": 1,
    "RoleId": 1
}' http://localhost:8080/roles/assign
```

