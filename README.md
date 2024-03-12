# Blog App REST API with Clean Architecture

This project aims to develop a REST API for blog-app using Golang and PostgreSQL, adhering to a clean architecture approach. The project is easily runnable using Docker and Docker Compose.

## Installation

1. Clone the project:

    ```bash
    git clone https://github.com/mehmet-dogru/go-blog-app.git
    ```

2. Navigate to the project directory:

    ```bash
    cd go-blog-app
    ```

3. Run the project using Docker Compose:

    ```bash
    docker-compose up --build
    ```

This command starts the PostgreSQL database and the Golang REST API server.

## Usage

You can access the API using an API client (e.g., Postman or cURL). The fundamental endpoints of the API are:

- **POST /users/register** : Creates a new user registration.
- **POST /users/login**    : Logs in with an existing user and creates a session (JWT token).

## Architecture

This project follows the principles of Clean Architecture. The core components include:

- **Domain**: The layer containing fundamental rules and data structures representing the business logic.
- **Use Cases**: The layer implementing application functionality and business workflows.
- **Interfaces**: The layer containing external components for interaction, such as HTTP Server and Database.
- **Frameworks & Drivers**: The layer directly interacting with external components and connecting lower layers.

## Contribution

If you wish to contribute to this project, please discuss your changes by opening an issue before submitting a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
