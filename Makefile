COMPOSE_PATH=./deployments/dev/docker-compose.yml
LOC_COMPOSE_PATH=./deployments/loc/docker-compose.yml
TEST_COMPOSE_PATH=./deployments/test/docker-compose.yml

#loc
build-loc:
	docker-compose -f $(LOC_COMPOSE_PATH) build --no-cache --parallel --force-rm
	docker-compose -f $(LOC_COMPOSE_PATH) up --remove-orphans --force-recreate -d

up-loc:
	docker-compose -f $(LOC_COMPOSE_PATH) up --no-recreate -d

down-loc:
	docker-compose -f$(LOC_COMPOSE_PATH) down

exec-loc:
	docker exec -it go-notifier bash

#dev
build:
	docker-compose -f $(COMPOSE_PATH) build --no-cache --parallel --force-rm
	docker-compose -f $(COMPOSE_PATH) up --remove-orphans --force-recreate -d

up:
	docker-compose -f $(COMPOSE_PATH) up --no-recreate -d

down:
	docker-compose -f$(COMPOSE_PATH) down

ps:
	docker-compose -f$(COMPOSE_PATH) ps

exec:
	docker exec -it go-notifier bash

logs:
	docker logs go-notifier-app

#test
test:
	docker-compose -f $(TEST_COMPOSE_PATH) up --build --abort-on-container-exit
	docker-compose -f$(TEST_COMPOSE_PATH) down --volumes
down-test:
	docker-compose -f$(TEST_COMPOSE_PATH) down