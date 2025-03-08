.PHONY: docker_run_db_redis
docker_run_db_redis:
	docker-compose -f deployments/docker-compose.yml up

.PHONY: migration_up
migration_up:
	MIGRATION_UP=true go run ./cmd/migration

.PHONY: migration_down
migration_down:
	MIGRATION_UP=false go run ./cmd/migration