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

