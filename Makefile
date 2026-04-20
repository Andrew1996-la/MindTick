include .env
export

export PROJECT_ROOT=$(shell pwd)
  
env-up:
	docker compose up -d mind-tick-postgres

env-down:
	docker compose down mind-tick-postgres

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
