#!/bin/bash

# Script to generate a secure JWT_SECRET and update .env file
# Usage: ./generate_jwt_secret.sh

ENV_FILE="./.env"
NEW_JWT_SECRET=$(openssl rand -base64 32)

# Check if .env file exists
if [ ! -f "$ENV_FILE" ]; then
    echo "Error: $ENV_FILE file not found."
    exit 1
fi

# Update JWT_SECRET in .env file
if grep -q "^JWT_SECRET=" "$ENV_FILE"; then
    # If JWT_SECRET line exists, replace it
    sed -i "s|^JWT_SECRET=.*|JWT_SECRET=$NEW_JWT_SECRET|" "$ENV_FILE"
else
    # If JWT_SECRET line doesn't exist, add it
    echo "JWT_SECRET=$NEW_JWT_SECRET" >> "$ENV_FILE"
fi

echo "JWT_SECRET has been generated and updated in $ENV_FILE"
echo "New JWT_SECRET: $NEW_JWT_SECRET"
