Here's an updated `README.md` to include the blockchain node as part of the project:

```markdown
# Startrix Task

Startrix Task is a multi-component project consisting of a backend API, a CLI tool, a frontend application, and a blockchain node. It is designed to demonstrate a blockchain-based wallet system with a modern web interface and command-line utilities.

## Project Structure

The project is organized into the following directories:

- **`backend/`**: A Go-based backend API built with [Fiber](https://gofiber.io/).
- **`cli/`**: A command-line interface (CLI) tool built with [Cobra](https://github.com/spf13/cobra) for managing wallets.
- **`frontend/`**: A React-based frontend application built with [Next.js](https://nextjs.org/).
- **`blockchain/`**: The blockchain node, responsible for verifying and processing transactions.

## Features

### Backend
- RESTful API built with Fiber.
- Example endpoint: `GET /` returns a "Hello, Fiber!" message.
- Configured to run on port `3000`.
- Connects to the blockchain node for transaction validation.

### CLI
- Wallet management commands:
  - `createWallet`: Creates a new wallet and saves it as a JSON file.
  - `signTransaction`: Placeholder for signing transactions.
- Built with Cobra for easy extensibility.

### Frontend
- React-based UI built with Next.js.
- TailwindCSS for styling.
- Includes a development server, build scripts, and production-ready configurations.

### Blockchain Node
- A custom blockchain node responsible for validating and processing transactions.
- Interacts with the backend API to validate transactions.
- Can be run locally for development or deployed for a testnet or mainnet environment.

## Prerequisites

Before you begin, ensure you have the following installed:

- **Go**: Version `1.22.2` or higher.
- **Node.js**: Version `16.x` or higher.
- **Docker**: For running MongoDB via `docker-compose`.
- **Blockchain Node Dependencies**: If running the blockchain node locally, ensure you have the required blockchain dependencies.

## Getting Started

### Backend

1. Navigate to the `backend` directory:
   ```bash
   cd backend
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. Run the backend server:
   ```bash
   go run main.go
   ```

   The API will be available at `http://localhost:3000`.

### CLI

1. Navigate to the `cli` directory:
   ```bash
   cd cli
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

3. To create a new wallet, use the `createWallet` command:
   ```bash
   go run main.go createWallet
   ```

4. To sign a transaction, use the `signTransaction` command:
   ```bash
   go run main.go signTransaction
   ```

### Frontend

1. Navigate to the `frontend` directory:
   ```bash
   cd frontend
   ```

2. Install dependencies:
   ```bash
   npm install
   ```

3. Run the development server:
   ```bash
   npm run dev
   ```

   The frontend will be available at `http://localhost:3000`.

### Blockchain Node

1. Navigate to the `blockchain` directory:
   ```bash
   cd blockchain
   ```

2. Install any necessary dependencies for the blockchain node.

3. Run the blockchain node:
   ```bash
   go run main.go
   ```

   The node will start and begin validating transactions sent by the backend API.

### Docker

The project includes a `docker-compose.yml` file to run MongoDB. To set up and run MongoDB in a container:

1. Run the following command in the root directory of the project:
   ```bash
   docker-compose up -d
   ```

   This will start a MongoDB container that you can use with the backend.

## Folder Structure

- **`backend/`**: Contains the Go backend code (Fiber API).
- **`cli/`**: Contains the Go CLI tool for managing wallets.
- **`frontend/`**: Contains the React frontend code (Next.js).
- **`blockchain/`**: Contains the blockchain node code that validates transactions and manages the blockchain.
- **`docker-compose.yml`**: Docker configuration for running MongoDB.
- **`README.md`**: Documentation for the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

### Updates:
- Added **Blockchain Node** section under **Features** to describe its role and interaction.
- Included **Blockchain Node** setup instructions under **Getting Started** to guide users through running the node locally.
- Updated **Folder Structure** to reflect the inclusion of the `blockchain/` directory.

Let me know if you need further adjustments or details added!