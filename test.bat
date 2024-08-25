REM creates pg_data directory for postgresdb volume
if exist pg_data\NUL echo "Folder already exists"
if not exist pg_data\NUL MKDIR "pg_data"

REM creates pg_data_test directory for postgres_test volume
if exist pg_data_test\NUL echo "Folder already exists"
if not exist pg_data_test\NUL MKDIR "pg_data_test"

REM auto generate env vars
if exist .env echo "Env already exists"
if not exist .env (
    echo "# Postgres"
    echo "DB_HOST=postgresdb"
    echo "DB_DRIVER=postgres"
    echo "DB_USER=spuser"
    echo "DB_PASSWORD=SPuser"
    echo "DB_NAME=psychic-waffle"
    echo "DB_PORT=5432"

    echo "# Postgres Test"
    echo "TEST_DB_HOST=postgres_test"
    echo "TEST_DB_DRIVER=postgres"
    echo "TEST_DB_USER=spuser"
    echo "TEST_DB_PASSWORD=SPuser_test"
    echo "TEST_DB_NAME=psychic-waffle_test"
    echo "TEST_DB_PORT=5432"
) > .env