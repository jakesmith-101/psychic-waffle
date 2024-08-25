# psychic-waffle
An example of my Full-Stack potential.

- Frontend built in Svelte, initialised with sveltekit
- Go
- PostgreSQL
- Docker ofc, maybe look into other areas, kubernetes???

## Quickstart

- build: `docker-compose build`
- start: `docker-compose up`
- stop:  `docker-compose down`

### env vars for postgres containers
Generate a .env by running the test script. <br />
Or create the .env in root with these vars:

```
# Postgres
DB_HOST=
DB_DRIVER=
DB_USER=
DB_PASSWORD=
DB_NAME=
DB_PORT=5432

# Postgres Test
TEST_DB_HOST=
TEST_DB_DRIVER=
TEST_DB_USER=
TEST_DB_PASSWORD=
TEST_DB_NAME=
TEST_DB_PORT=5432
```
