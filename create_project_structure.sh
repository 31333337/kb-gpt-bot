#!/bin/bash

# Create the directory structure and empty Go files for the Keybase bot project

# Create directories
mkdir -p cmd
mkdir -p pkg/commands
mkdir -p pkg/openai
mkdir -p pkg/keybase

# Create main.go in cmd directory
touch cmd/main.go

# Create Go files in pkg/commands directory
touch pkg/commands/role.go
touch pkg/commands/temperature.go
touch pkg/commands/settings.go
touch pkg/commands/info.go
touch pkg/commands/config.go

# Create Go files in pkg/openai directory
touch pkg/openai/openai.go

# Create Go files in pkg/keybase directory
touch pkg/keybase/keybase.go

# Print the directory structure
echo "Created project structure:"
tree
