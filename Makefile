.PHONY: docker_run_dev_dependencies
docker_run_dev_dependencies:
	docker-compose -f dependencies/docker-compose.yml up