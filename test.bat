REM creates pg_data directory for postgresdb volume
if exist .\pg_data\NUL echo "Folder already exists"
if not exist .\pg_data\NUL MKDIR "pg_data"

REM creates pg_data_test directory for postgres_test volume
if exist .\pg_data_test\NUL echo "Folder already exists"
if not exist .\pg_data_test\NUL MKDIR "pg_data_test"


