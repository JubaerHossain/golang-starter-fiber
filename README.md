# Attendance System Documentation

## Overview

The Attendance System is a web application designed to manage user attendance records. It includes features for creating, updating, and deleting user profiles, and recording attendance.

## Table of Contents

- [Project Structure](#project-structure)
- [Setup](#setup)
- [Usage](#usage)
- [Contributing](#contributing)
- [License](#license)

## Project Structure

The project follows a modular structure for better organization. Here's a breakdown of the main components:

- **app**: Contains the application logic and main entry point.
- **controllers**: Handles HTTP requests, business logic, and response formatting.
- **database**: Manages database connections and interactions.
- **models**: Defines data structures and entities used in the application.
- **repository**: Provides an abstraction for database operations.
- **routes**: Defines API endpoints and connects them to controller methods.
- **services**: Implements the core business logic and interacts with repositories.
- **utils**: Contains utility functions and shared components.
- **main.go**: The application's entry point.

## Setup <a name="setup"></a>

1. Clone the repository:

   ```bash
   git clone https://github.com/your-username/attendance-system.git
   cd attendance-system
   ```

2. Install dependencies:

   ```bash
    go mod tidy
   ```

3. Create a `.env` file in the root directory and add the following environment variables:

   ```bash
    # Database
    MONGOURI=mongodb://localhost:27017
   PORT=3000
   APP_URL=http://localhost
   DB_NAME=hrm
   ```

4. Start the server:

   ```bash
   go run main.go
   ```

## Usage <a name="usage"></a>

### API Endpoints

| Endpoint              | Method | Description                         |
| --------------------- | ------ | ----------------------------------- |
| /api/users            | GET    | Get all users                       |
| /api/users            | POST   | Create a new user                   |
| /api/users/:id        | GET    | Get a user by ID                    |
| /api/users/:id        | PUT    | Update a user by ID                 |
| /api/users/:id        | DELETE | Delete a user by ID                 |
| /api/users/:id/attend | POST   | Record attendance for a user        |
| /api/users/:id/attend | GET    | Get attendance records for a user   |
| /api/users/:id/attend | PUT    | Update attendance record for a user |
| /api/users/:id/attend | DELETE | Delete attendance record for a user |

## Build the Docker image:

```bash
docker build -t attendance-app .
```

## Run the Docker container:

```bash
docker run -p 3000:3000 attendance-app
```

### API Documentation

The API documentation is available at [http://localhost:3000/api-docs](http://localhost:3000/api-docs).

## Contributing <a name="contributing"></a>

Contributions are welcome! Please refer to the [contributing guide](CONTRIBUTING.md) for more details.

## License <a name="license"></a>

This project is licensed under the [MIT License](LICENSE).
