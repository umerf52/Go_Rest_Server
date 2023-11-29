# REST Server in Go
A REST server written in Go to show basic CRUD operations.

- Uses:
  - [Gin](https://github.com/gin-gonic/gin/)
  - [Mongo Driver](https://pkg.go.dev/go.mongodb.org/mongo-driver)

## Usage
- Make sure MongoDB is installed and running.
- Create a database named `mydb` and a collection named `mycollection`.
- Clone the repo.
- Run `go run main.go`.
- API will be available at `http://127.0.0.1:8080/data/`.

## Endpoints
- `GET /data/` - Get all
- `GET /data/:id` - Get one
- `POST /data/` - Create
```
{
  "name": "John Doe"
}
```
- `PUT /data/:id` - Update
```
{
  "name": "Update Name"
}
```
- `DELETE /data/:id` - Delete

## Files
- `main.go` - Main file
- `config/db.go` - Code for connecting to MongoDB
- `handlers/handlers.go` - Handler methods for endpoints
- `models/models.go` - Model for representing data
