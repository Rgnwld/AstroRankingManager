# Environment

COMPOSE=docker-environment/docker-compose.yml
COMPOSEDBONLY=docker-environment/docker-compose-db-only.yml
DB_LOCAL=docker-environment/database/compose-local.yml
API_LOCAL=docker-environment/go/compose-local.yml

.PHONY: help
help:
	@printf '\033[1;33mNOTE: for local development, please provide a .env file on root'
	@printf '\nYou can see more about the .env in the sample.env\n'
	@printf '\n\033[mhelp	- Displays information about available commands.'
	@printf '\nbuild	- Build services'
	@printf '\ndev	- Start services'
	@printf '\ndb	- Start database container for local development'
	@printf '\ndown	- Stops and removes defined services'
	@printf '\nlogs	- Displays the records (logs) of the defined services '

.PHONY: build
build:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} build --no-cache

.PHONY: dev
dev:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} --env-file .env up -d --build --remove-orphans

.PHONY: db
db:
	docker-compose -f ${COMPOSEDBONLY} -f ${DB_LOCAL} --env-file .env up -d --remove-orphans

.PHONY: down
down:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} down

.PHONY: logs
logs:
	docker-compose -f ${COMPOSE} -f ${DB_LOCAL} -f ${API_LOCAL} logs

