# Todo App

This is a TODO application project developed in Go. It uses a modular architecture with a clear directory structure to handle databases, models, and routes.

## Project Structure

```
todo/
│── cmd/
│   └── main/
│       └── main.go  # Application entry point
│
│── internal/
│   ├── database/
│   │   ├── dbCon.go  # Database connection
│   │   ├── tasks.db  # SQLite database file
│   │   └── verifyDBExist.go  # Database existence verification
│   │
│   ├── models/
│   │   └── model.go  # Data model definitions
│   │
│   ├── routes/
│   │   └── routes.go  # API route definitions
│
│── .gitignore  # Git ignored files
│── go.mod  # Go modules
```

## Installation

1. Clone the repository:

   ```sh
   git clone https://github.com/LeonibelDev/go-todo/
   cd todo
   ```

2. Install dependencies:

   ```sh
   go mod tidy
   ```

3. Run the project:

   ```sh
   go run cmd/main/main.go
   ```

## Features

- SQLite database connection.
- Task model definitions.
- RESTful routes for CRUD operations.

## Contribution

If you wish to contribute, please fork the repository and submit a pull request with your improvements.

## License

This project is under the MIT license.

