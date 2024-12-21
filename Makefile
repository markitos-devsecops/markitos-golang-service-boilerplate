#!/bin/bash
#:[.''.]:>-------------------------------------------
#:[.''.]:> Author:
#:[.''.]:> Marco Antonio Rubio Lopez
#:[.''.]:> markitos.es.info@gmail.com
#:[.''.]:> diciembre 2024
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><runs>
run:
	go run .
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><testing>
test:
	go test ./...
testv:
	go test -v ./...
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><database - migrations - sqlc>
createdb:
	docker exec -it markitos-svc-postgres createdb --username=admin --owner=admin markitos-svc-boilerplate
dropdb: 
	docker exec -it markitos-svc-postgres psql -U admin -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'markitos-svc-boilerplate' AND pid <> pg_backend_pid();"
	docker exec -it markitos-svc-postgres dropdb -U admin markitos-svc-boilerplate
migrate-init:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "postgresql://admin:admin@markitos-svc-postgres:5432/markitos-svc-boilerplate?sslmode=disable" create \
		-ext sql -dir internal/infrastructure/database/migrations/ -seq init_schema 
migrate-up:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-svc-postgres:5432/markitos-svc-boilerplate?sslmode=disable" -verbose up		
migrate-down:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-svc-postgres:5432/markitos-svc-boilerplate?sslmode=disable" -verbose down $(or $(VERSION),1)
migrate-version:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-svc-postgres:5432/markitos-svc-boilerplate?sslmode=disable" version
migrate-goto:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-svc-postgres:5432/markitos-svc-boilerplate?sslmode=disable" goto $(or $(VERSION),1)
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:><security> TODO: ver como hacerlo con otros
appsec-sast-sca:
	docker run --rm -v $(shell pwd):/src returntocorp/semgrep semgrep --config=auto /src --verbose
appsec-gitleaks:
	docker run --rm -v $(shell pwd):/repo zricethezav/gitleaks:latest detect --source /repo
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:><building>
image-build:
	docker build -t ghcr.io/markitos-es/markitos-svc-boilerplate:$(or $(TAG),1.0.0) .
image-push:
	docker push ghcr.io/markitos-es/markitos-svc-boilerplate:$(or $(TAG_SEMVER),1.0.0) && \
	docker image rm --force ghcr.io/markitos-es/markitos-svc-boilerplate:$(or $(TAG_SEMVER),1.0.0)
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:> License: OpenSource :) to pa ti!
#:[.''.]:> markitos.es
#:[.''.]:>-------------------------------------------
