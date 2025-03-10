.PHONY: docker_run_db_redis
docker_run_db_redis:
	docker-compose -f deployments/docker-compose.yml up

.PHONY: go_run_app
go_run_app:
	go run ./cmd/app

.PHONY: migration_up
migration_up:
	MIGRATION_UP=true go run ./cmd/migration

.PHONY: migration_down
migration_down:
	MIGRATION_UP=false go run ./cmd/migration

.PHONY: go_test
go_test:
	go test ./...

.PHONY: gen_jet_sql_builder
gen_jet_sql_builder:
	jet -source=mysql -dsn="recommendation-mysql-user:recommendation-mysql-password@tcp(localhost:3306)/recommendation?charset=utf8mb4&parseTime=true&multiStatements=true" \
		-ignore-tables=goose_db_version \
		-path=./internal/app/database/.jet_gen