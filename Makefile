run-base-compose:
	docker compose -f ./compose/mongodb.yaml -f ./compose/postgres.yaml up -d

stop-base-compose:
	docker compose -f ./compose/mongodb.yaml -f ./compose/postgres.yaml down --remove-orphans