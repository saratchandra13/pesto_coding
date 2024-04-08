# pesto_coding
Coding file for pesto interviews

## API Testing

You can test the API using the following `curl` commands:

### Create a new product

```bash
curl -X POST -H "Content-Type: application/json" -d '{
    "Name": "Test Product",
    "Description": "This is a test product",
    "Price": 99.99,
    "Category": "Test Category",
    "ImageURL": "http://example.com/test.jpg"
}' http://localhost:8080/products
```

### Get a product by ID

```bash
curl http://localhost:8080/products/1
```
