#!/bin/bash

# Set the package name and version
PACKAGE_NAME=dollbox
PACKAGE_VERSION=1.0.0

# Set the Go module path
MODULE_PATH=https://github.com/DollBoxPM/DollBoxPM

# Create the Go module file
go mod init $MODULE_PATH
go mod init

go mod init github.com/DollBoxPM/DollBoxPM
go get github.com/go-git/go-git/v5
go install github.com/go-git/go-git/v5@latest

# Build the package
go build -o $PACKAGE_NAME main.go

# Create the module file
go mod tidy

# Create the sum file
go mod verify

# Move the executable
