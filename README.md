# Simple Task List App in Go

This is a basic CRUD program in Go that implements a simple task list application. It uses the Gorilla Mux library to handle routes and HTTP requests.

## Project Initialization

Before running the program, make sure to initialize the Go module and install the Gorilla Mux library. You can do this by executing the following commands in your terminal:

```bash
go mod init [project_name]
go get github.com/gorilla/mux
```

## Using the Application

The application provides the following routes:

- `GET /tasks`: Get the list of tasks.
- `GET /tasks/{id}`: Get details of a specific task.
- `POST /tasks`: Create a new task.
- `PUT /tasks/{id}`: Update an existing task.
- `DELETE /tasks/{id}`: Delete a task.

Make sure to test these routes using tools like postman or curl 
