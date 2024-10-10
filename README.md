# CLI Single-Sign-On

[![Go Version](https://img.shields.io/badge/go-1.15-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## Description

This CLI is an example of implementing OAuth2 authentication (authorization code flow) by command line. It provides a secure method of authentication by delegating the login functionality to an identity provider. It provides a straightforward way for users to authenticate via their web browser and securely handle authentication tokens.

## Features

- OAuth2 authentication support
- Secure handling of authentication tokens
- Easy integration with external OAuth2 providers

## Getting Started

### Prerequisites

- Go 1.22 or higher

### Setting Up the Authentication Environment

Before running the CLI, you need to set up the authentication environment which includes Keycloak and PostgreSQL.

1. **Start Keycloak and PostgreSQL Containers:**

   Run the following command to start the containers using Docker Compose:

   ```bash
   docker compose up -d
   ```

    Ensure your docker-compose.yml file is configured to set up Keycloak and PostgreSQL correctly.

2. **Configure Keycloak:**

    - Access the Keycloak admin console at http://localhost:8080/admin.
    - Log in with the admin credentials (default: admin/admin if not changed).
    - Navigate to the "dev" realm.
    - Under Clients, find the client auth-cli and regenerate the secret.
    - Update the ClientSecret in your login.go file with the new secret.

3. **Set Up AdminCLI User:**

    Still within the Keycloak admin console:

    - Go to Users, add a new user called admincli.
    - Set a password for admincli under the Credentials tab and disable Temporary to make it a permanent password.

### Installing

To get started with this CLI tool, clone the repository and build the executable:

```bash
git clone https://github.com/wallanaq/go-auth-cli.git
cd go-auth-cli/src
go build -o ../authcli
```

### Usage

To start the authentication process, run:

```bash
./authcli login
```

### Contributing

Contributions are welcome! Please feel free to submit a pull request.

### License

This project is licensed under the MIT License 
