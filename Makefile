COMPOSE_PATH=./deployments/docker-compose.yml

build:
	docker-compose -f $(COMPOSE_PATH) build --no-cache --parallel --force-rm
	docker-compose -f $(COMPOSE_PATH) up --remove-orphans --force-recreate -d

up:
	docker-compose -f $(COMPOSE_PATH) up --no-recreate -d

down:
	docker-compose -f$(COMPOSE_PATH) down

ps:
	docker-compose -f$(COMPOSE_PATH) ps

logs:
	docker logs go-notifier-app

