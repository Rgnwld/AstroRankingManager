# Environment

COMPOSE=docker-environment/docker-compose.yml
DB_LOCAL=docker-environment/database/compose-local.yml
API_LOCAL=docker-environment/go/compose-local.yml

.PHONY: help
help:
	@echo 'help	- Displays information about available commands.'
	@echo 'up	- Start services defined in local development'
	@echo 'down	- Stops and removes defined services'
	@echo 'logs	- Displays the records (logs) of the defined services '

.PHONY: up
up:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} up -d --build --remove-orphans

.PHONY: down
down:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} down

.PHONY: logs
logs:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} logs

