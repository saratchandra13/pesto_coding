# pesto_coding
Coding file for pesto interviews

## API Testing

You can test the API using the following `curl` commands:

### Create a new user

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "ID": "1",
    "Username": "testuser",
    "Email": "testuser@example.com",
    "Password": "password123"
}' http://localhost:8080/users

```

### Authenticate a user

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "ID": "1",
    "Password": "password123"
}' http://localhost:8080/users/auth
```

### Get user details

```bash
curl -X GET http://localhost:8080/users/1
```

