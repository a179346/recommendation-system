## Prerequisite

- [docker](https://www.docker.com/) & `docker-compose`: Run the dev MySQL and Redis server
- [go](https://go.dev/) version >= 1.23
- `make`: optional
- [Postman](https://www.postman.com/): Test recommendation system

## How to get the server running

> ⭐ The following commands should be executed from **the project's root**.

#### 1. Run the dev MySQL and Redis server

```sh
make docker_run_db_redis

or

docker-compose -f deployments/docker-compose.yml up
```

Please wait until MySQL is ready.

#### 2. Run the database migration

```sh
make migration_up

or

MIGRATION_UP=true go run ./cmd/migration
```

The above command will create the database table and insert initial data into the database automatically.

#### 3. Start recommendation system

```sh
make go_run_app

or

go run ./cmd/app
```

You have successfully brought the server online! Let’s test it.

## How to test recommendation system

Here is the [API specification](../api/swagger.yml). You can [import it to Postman](https://learning.postman.com/docs/getting-started/importing-and-exporting/importing-from-swagger/).

Now you can test the application however you like.
