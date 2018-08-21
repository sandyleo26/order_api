#!/bin/bash

export DB_HOST="localhost"
export DB_PORT="5432"
export DB_NAME="sha"
export DB_USER="postgres"
export DB_PASS="pass"
export DB_SSL="disable"

make container && make container-start
