# List Real Command
RUN_CMD = go run main.go http ./config.yaml
TEST_CMD = go test -race -v -cover -covermode=atomic ./...
DOCKER_UP_CMD = docker-compose up -d --build app
DOCKER_DOWN_CMD = docker-compose stop app

.PHONY: run test up down

run:
	@$(RUN_CMD)

test:
	@$(TEST_CMD)

up:
	@$(DOCKER_UP_CMD)

down:
	@$(DOCKER_DOWN_CMD)

logs:
	@$(DOCKER_LOGS_CMD)

stats:
	@$(DOCKER_STATS_CMD)