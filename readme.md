
# Startrix Task

Startrix Task is a multi-component project consisting of a backend API, a CLI tool, a frontend application, and a blockchain node. It is designed to demonstrate a blockchain-based wallet system with a modern web interface and command-line utilities.

![Startrix Task Overview](https://coffeecoin.muhammedsirajudeen.in/preview.png)

## Project Structure

The project is organized into the following directories:

- **`backend-blockchain/`**: A Go-based backend API built with [Fiber](https://gofiber.io/).
- **`cli/`**: A command-line interface (CLI) tool built with [Cobra](https://github.com/spf13/cobra) for managing wallets.
- **`frontend/`**: A React-based frontend application built with [Next.js](https://nextjs.org/).

## Features

### Backend
# API Documentation

## Overview
This is a RESTful API built with the Fiber framework. It provides endpoints for managing transactions, querying balances, and performing airdrops in a blockchain-like system. The API is configured to run on port `3000`.

---

## Endpoints

### 1. **Create Transaction**
   - **URL:** `/transaction`
   - **Method:** `POST`
   - **Description:** Verifies and records a transaction between two accounts.
   - **Request Body:**
     ```json
     {
       "sender": "string",        // Hex-encoded public key of the sender
       "recipient": "string",     // Address of the recipient
       "amount": 100.0,           // Amount to transfer
       "signature": "string",     // Base64-encoded signature
       "previous_block": "string" // (Optional) Previous block hash
     }
     ```
   - **Responses:**
     - `200 OK`: Transaction verified and recorded.
     - `400 Bad Request`: Invalid JSON or insufficient balance.
     - `401 Unauthorized`: Invalid signature.

---

### 2. **Get Account Balance**
   - **URL:** `/balance/:address`
   - **Method:** `GET`
   - **Description:** Retrieves the balance of a specific account.
   - **Path Parameter:**
     - `address` (string): The address of the account.
   - **Responses:**
     - `200 OK`: Returns the balance of the account.
       ```json
       {
         "address": "string",
         "balance": 100.0
       }
       ```

---

### 3. **Get All Transactions**
   - **URL:** `/transactions`
   - **Method:** `GET`
   - **Description:** Retrieves the entire transaction history.
   - **Responses:**
     - `200 OK`: Returns a list of all transactions.
       ```json
       [
         {
           "sender": "string",
           "recipient": "string",
           "amount": 100.0,
           "signature": "string",
           "previous_block": "string"
         }
       ]
       ```

---

### 4. **Get Transactions by Address**
   - **URL:** `/transactions/:address`
   - **Method:** `GET`
   - **Description:** Retrieves all transactions associated with a specific address (as sender or recipient).
   - **Path Parameter:**
     - `address` (string): The address to filter transactions by.
   - **Responses:**
     - `200 OK`: Returns a list of transactions associated with the address.
       ```json
       {
         "transactions": [
           {
             "sender": "string",
             "recipient": "string",
             "amount": 100.0,
             "signature": "string",
             "previous_block": "string"
           }
         ]
       }
       ```

---

### 5. **Airdrop Coins**
   - **URL:** `/airdrop`
   - **Method:** `POST`
   - **Description:** Adds 100 coins to the balance of a specified account.
   - **Request Body:**
     ```json
     {
       "address": "string" // Address to receive the airdrop
     }
     ```
   - **Responses:**
     - `200 OK`: Airdrop successful.
       ```json
       {
         "message": "Airdropped 100 coins",
         "address": "string",
         "balance": 200.0
       }
       ```
     - `400 Bad Request`: Invalid request body.

---

## Notes
- **Transaction Verification:** Transactions are verified using Ed25519 signatures. The sender's public key, recipient address, amount, and signature are validated to ensure authenticity.
- **Genesis Block:** The system initializes with a genesis block to establish the blockchain structure.
- **Concurrency:** Mutex locks are used to ensure thread-safe operations on shared resources like transaction history and account balances.
- RESTful API built with Fiber.
- Example endpoint: `GET /` returns a "Hello, Fiber!" message.
- Configured to run on port `3000`.
- Connects to the blockchain node for transaction validation.


## Build Instructions

### Backend

1. Navigate to the `backend-blockchain` directory:
   ```bash
   cd backend-blockchain
   ```

2. Build the backend binary:
   ```bash
   go build -o backend-server main.go
   ```

   This will generate an executable file named `backend-server`.

3. Run the binary:
   ```bash
   ./backend-server
   ```

### CLI

# CLI Documentation

The `cli` directory contains the command-line interface (CLI) implementation for interacting with the application. This CLI allows users to perform various operations such as creating wallets and managing tasks.

## Prerequisites

Ensure you have the following installed on your system:
- [Go](https://golang.org/dl/) (version 1.16 or later)

## Usage

1. **Navigate to the CLI directory**  
   Change your working directory to the `cli` folder where the CLI source code is located.

2. **Build the CLI binary**  
   Use the `go build` command to compile the CLI source code into an executable binary. The output binary will be named `wallet-cli`.

3. **Execute CLI commands**  
   Run the generated binary to execute commands. For example, you can create a new wallet using the `createWallet` command.

## Example Commands

- **Install Startrix-AI Binary (Linux)**  
   If you are using Linux, you can install the `startrix-ai` binary by running:
   ```bash
   sudo install -m 755 wallet-cli /usr/local/bin/startrix-ai
   ```
   This will allow you to use the `startrix-ai` command globally.

## Notes

- Ensure the `cli` directory contains the `main.go` file, which serves as the entry point for the CLI application.
- The `wallet-cli` binary must have execute permissions. If not, you can grant permissions using:
  ```bash
  chmod +x wallet-cli
  ```
- Refer to the application's documentation or source code for additional commands and their usage.
   ```

   The API will be available at `http://localhost:3000`.
   """
   This module provides functionality for creating and managing cryptocurrency wallets, 
   as well as signing transactions. It includes the following key functions:

   Functions:
   ----------
   - createWallet():
      Creates a new cryptocurrency wallet. This function generates a new private key 
      and derives the corresponding public key and wallet address. The wallet can be 
      used to send and receive cryptocurrency transactions.
           
      ### Example: Create Wallet

      To create a new wallet using the `startrix-cli`, run the following command:

      ```bash
         startrix-cli createWallet
      ```

      This will generate a new wallet with a unique address and private key. The output will look similar to:

      ```plaintext
      Wallet created successfully!
      Address: 0x123456789abcdef
      Private Key: abcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890
      ```

      Make sure to securely store the private key, as it is required to sign transactions.

   - signTransaction(transaction, privateKey):
      Signs a cryptocurrency transaction using the provided private key. This ensures 
      the authenticity and integrity of the transaction. The function takes the 
      transaction data and the private key as input, and returns the signed transaction.

   Other Commands:
   ---------------
   - Additional helper functions or utilities may be included in the module to support 
     wallet creation and transaction signing, such as key generation, address derivation, 
     and cryptographic operations.
   """

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

   The frontend will be available at `http://localhost:3001`.



## Folder Structure

- **`backend/`**: Contains the Go backend code (Fiber API).
- **`cli/`**: Contains the Go CLI tool for managing wallets.
- **`frontend/`**: Contains the React frontend code (Next.js).
- **`blockchain/`**: Contains the blockchain node code that validates transactions and manages the blockchain.
- **`README.md`**: Documentation for the project.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.