#!/bin/bash

# creates pg_data directory for postgresdb volume
mkdir -p pg_data

# creates pg_data_test directory for postgres_test volume
mkdir -p pg_data_test

# check if the file exists already
if [ -f ".env" ]; then
    echo "File .env already exists."
    exit 1
fi

# auto generate env vars
printf "# Postgres\nDB_HOST=postgresdb\nDB_DRIVER=postgres\nDB_USER=spuser\nDB_PASSWORD=SPuser\nDB_NAME=psychic-waffle\nDB_PORT=5432\n\n# Postgres Test\nTEST_DB_HOST=postgres_test\nTEST_DB_DRIVER=postgres\nTEST_DB_USER=spuser\nTEST_DB_PASSWORD=SPuser_test\nTEST_DB_NAME=psychic-waffle_test\nTEST_DB_PORT=5432" > .env