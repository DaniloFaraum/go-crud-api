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
  {
    "ID": 1,
    "Name": "John Doe",
    "Age": 30,
    "Gender": "Male",
    "Ocupation": "Engineer",
    "Salary": 50000,
    "Alive": true
  },
  "message": "create-person operation sucessful"
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
  "message": "list-persons operation sucessful"
]
```
### Get a Person by ID

- GET /api/persons/{id}
  - Response:
```
{
  {
    "ID": 1,
    "Name": "John Doe",
    "Age": 30,
    "Gender": "Male",
    "Ocupation": "Engineer",
    "Salary": 50000,
    "Alive": true
  },
  "message": "get-person operation sucessful"
}
```
### Update a Person

- PUT /api/persons/{id}
  - WIP

### Delete a Person

- DELETE /api/persons/{id}
  - Reponse:
```
{
    "data": {
        "ID": 1,
        "CreatedAt": "2024-08-31T21:05:53.7595227-03:00",
        "UpdatedAt": "2024-08-31T21:05:53.7595227-03:00",
        "DeletedAt": "2024-08-31T21:53:44.5380916-03:00",
        "Name": "Danilo",
        "Age": 18,
        "Gender": "male",
        "Ocupation": "developer",
        "Salary": 1,
        "Alive": true
    },
    "message": "delete-person operation sucessful"
}
```
