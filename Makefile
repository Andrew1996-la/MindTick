include .env
export

export PROJECT_ROOT=$(shell pwd)
  
env-up:
	@ docker compose up -d mind-tick-postgres

env-down:
	@docker compose down mind-tick-postgres

env-cleanup:
	@read -p "Очистить все voluem файлы окружения? Опастность утери данных![y/N]" ans; \
	if [ "$$ans" = "y" ]; then \
		docker compose down mind-tick-postgres && \
		rm -rf out/pgdata && \
		echo "Файлы окружения очищены"; \
	else \
		echo "Очистка окружения отменена"; \
	fi;

migrate-create:
	@if [ -z "$(seq)" ]; then \
		echo "отсутствует необходимый параметр seq. Пример migrate-create seq=init";\
		exit 1; \
	fi; \
	docker compose run --rm min-tick-postgres-migrate \
		create \
		-ext sql \
		-dir /migrations \
		-seq  "$(seq)"

migrate-up:
	@make migrate-action action=up

migrate-down:
	@make migrate-action action=down

migrate-action:
	@if [ -z "$(action)" ]; then \
		echo "отсутствует необходимый параметр action. Пример migrate-action action=up ";\
		exit 1; \
	fi; \
	docker compose run --rm min-tick-postgres-migrate \	
			-path /migrations \
			-database postgres://${POSTGRES_USER}:${POSTGRES_PASSWORD}@mind-tick-postgres:5432/${POSTGRES_DB}?sslmode=disable \
			"$(action)"
