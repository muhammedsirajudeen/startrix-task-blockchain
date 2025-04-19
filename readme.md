# Startrix Task

Startrix Task is a multi-component project consisting of a backend API, a CLI tool, and a frontend application. It is designed to demonstrate a blockchain-based wallet system with a modern web interface and command-line utilities.

## Project Structure

The project is organized into the following directories:

- **`backend/`**: A Go-based backend API built with [Fiber](https://gofiber.io/).
- **`cli/`**: A command-line interface (CLI) tool built with [Cobra](https://github.com/spf13/cobra) for managing wallets.
- **`frontend/`**: A React-based frontend application built with [Next.js](https://nextjs.org/).

## Features

### Backend
- RESTful API built with Fiber.
- Example endpoint: `GET /` returns a "Hello, Fiber!" message.
- Configured to run on port `3000`.

### CLI
- Wallet management commands:
  - `createWallet`: Creates a new wallet and saves it as a JSON file.
  - `signTransaction`: Placeholder for signing transactions.
- Built with Cobra for easy extensibility.

### Frontend
- React-based UI built with Next.js.
- TailwindCSS for styling.
- Includes a development server, build scripts, and production-ready configurations.

## Prerequisites

- **Go**: Version `1.22.2` or higher.
- **Node.js**: Version `16.x` or higher.
- **Docker**: For running MongoDB via `docker-compose`.

## Getting Started

### Backend

1. Navigate to the `backend` directory:
   ```bash
   cd backend