# Azion Storage Go

This project is a simple Go application that provides an edge storage service. It uses Chi router for handling HTTP requests and includes basic CRUD operations for managing buckets.

## Setup
1. Make sure you have Go installed.
2. Clone the repository.
3. Run `go mod tidy` to install dependencies.
4. Run `go run main.go` to start the server.

## Endpoints
- GET /buckets: List all buckets.
- POST /buckets: Create a new bucket.
- DELETE /buckets/{name}: Delete a specific bucket.

Feel free to explore and contribute!

[Azion Storage](https://www.azion.com/en/documentation/products/store/edge-storage/)
[Azion Go SDK](https://www.azion.com/en/documentation/devtools/sdk/go/)