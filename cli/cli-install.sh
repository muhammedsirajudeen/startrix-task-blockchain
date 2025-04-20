#!/bin/bash

# Function to check if a command is available
check_command() {
    command -v "$1" &> /dev/null
}

# Function to install wget
install_wget() {
    echo "wget is not installed. Installing wget..."
    sudo apt-get update && sudo apt-get install -y wget
}

# Check if Go is installed
if check_command "go"; then
    echo "Go is installed."
    go version
else
    echo "Go is not installed. Please install Go version 1.22.2."
    exit 1
fi

# Check if wget is installed, install if not
if ! check_command "wget"; then
    install_wget
else
    echo "wget is already installed."
fi

# Create a temporary directory
TEMP_DIR=$(mktemp -d)

# Download the ZIP file containing the CLI
echo "Downloading CLI..."
wget -P "$TEMP_DIR" https://github.com/muhammedsirajudeen/startrix-task-blockchain/raw/refs/heads/main/cli.zip

# Unzip the downloaded file
echo "Extracting CLI..."
unzip "$TEMP_DIR/cli.zip" -d "$TEMP_DIR"

# Inspect the extracted files to determine the correct directory structure
echo "Inspecting extracted files..."
ls -R "$TEMP_DIR"

# Move the contents of the extracted folder to the current directory
echo "Moving extracted files..."

# Change directory to the correct CLI folder
cd "$TEMP_DIR/cli" || { echo "Failed to change directory to $TEMP_DIR/cli"; exit 1; }

# Build the CLI
echo "Building the CLI..."
go build -o startrix-cli
if [ $? -eq 0 ]; then
    echo "CLI built successfully."
else
    echo "Failed to build the CLI."
    exit 1
fi

# Add the CLI to PATH if it's not already there
echo "Adding CLI to PATH..."
CLI_PATH=$(pwd)
if [[ ":$PATH:" != *":$CLI_PATH:"* ]]; then
    echo "export PATH=\$PATH:$CLI_PATH" >> ~/.bashrc
    export PATH=$PATH:$CLI_PATH
    echo "CLI added to PATH. Please restart your terminal or run 'source ~/.bashrc' to apply changes."
else
    echo "CLI is already in PATH."
fi
