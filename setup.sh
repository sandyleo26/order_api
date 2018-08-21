#!/bin/bash

export dbHost="localhost"
export dbPort="5432"
export dbName="sha"
export dbUser="postgres"
export dbPass="pass"
export dbSSL="disable"
export GOOGLE_DISTANCE_API_KEY=""

make setup

echo "build and start database"
make db-container && make db-container-start && sleep 3 && make db-migration

# mac
# make build && make start

# linux
echo "build and start api"
make build-linux && make start-linux

echo "Ready to go!"
echo "try curl -X GET "