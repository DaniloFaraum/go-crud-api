# Go CRUD API

This repository contains a simple CRUD API built with Go, using Gin for routing and GORM for database management with SQLite.

## Features

- **CRUD Operations**: Create, Read, Update, and Delete functionality for managing resources.
- **SQLite Database**: Persistent storage with SQLite.
- **Modular Structure**: Clean and organized code structure for scalability.
- **Go-Gin**: Lightweight web framework for API routing.
- **GORM**: ORM for Go, providing database abstraction.

## Getting Started

### Prerequisites

- Go 1.19+
- SQLite

### Installation

Clone the repository:

```
git clone https://github.com/DaniloFaraum/go-crud-api.git
```
Install Dependencies:
```
go mod tidy
```
To run the API locally:
```
go run main.go
```
## Usage
After running the API, you can interact with it using tools like curl, Postman, or any other HTTP client. The API listens on http://localhost:8080 by default.

### API Endpoints
Here are the available endpoints:

- Create a Person

  - POST /api/persons
Request Body:
```
{
  "Name": "John Doe",
  "Age": 30,
  "Gender": "Male",
  "Ocupation": "Engineer",
  "Salary": 50000,
  "Alive": true
}
```
  - Response:
```
{
  "ID": 1,
  "CreatedAt": "2024-08-31T12:34:56Z",
  "UpdatedAt": "2024-08-31T12:34:56Z",
  "DeletedAt": null,
  "Name": "John Doe",
  "Age": 30,
  "Gender": "Male",
  "Ocupation": "Engineer",
  "Salary": 50000,
  "Alive": true
}
```
### Get All Persons

- GET /api/persons
  - Response:
```
[
  {
    "ID": 1,
    "Name": "John Doe",
    "Age": 30,
    "Gender": "Male",
    "Ocupation": "Engineer",
    "Salary": 50000,
    "Alive": true
  },
  ...
]
```
### Get a Person by ID

- GET /api/persons/{id}
  - Response:
```
{
  "ID": 1,
  "Name": "John Doe",
  "Age": 30,
  "Gender": "Male",
  "Ocupation": "Engineer",
  "Salary": 50000,
  "Alive": true
}
```
### Update a Person

- PUT /api/persons/{id}
  - WIP

### Delete a Person

- DELETE /api/persons/{id}
  - WIP

