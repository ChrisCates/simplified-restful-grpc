#!/bin/bash

# Set GRPC_ROOT variable
export GRPC_ROOT=$PWD

# Run the build script
go run ./bin/server/server.go
