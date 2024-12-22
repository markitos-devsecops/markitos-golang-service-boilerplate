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
	docker exec -it markitos-service-postgres createdb --username=admin --owner=admin markitos-service-boilerplate
dropdb: 
	docker exec -it markitos-service-postgres psql -U admin -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'markitos-service-boilerplate' AND pid <> pg_backend_pid();"
	docker exec -it markitos-service-postgres dropdb -U admin markitos-service-boilerplate
migrate-init:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "postgresql://admin:admin@markitos-service-postgres:5432/markitos-service-boilerplate?sslmode=disable" create \
		-ext sql -dir internal/infrastructure/database/migrations/ -seq init_schema 
migrate-up:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-service-postgres:5432/markitos-service-boilerplate?sslmode=disable" -verbose up		
migrate-down:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-service-postgres:5432/markitos-service-boilerplate?sslmode=disable" -verbose down $(or $(VERSION),1)
migrate-version:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-service-postgres:5432/markitos-service-boilerplate?sslmode=disable" version
migrate-goto:
	docker run --user $(id -u):$(id -g) -v ./internal/infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@markitos-service-postgres:5432/markitos-service-boilerplate?sslmode=disable" goto $(or $(VERSION),1)
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:><security>
appsec-sast:
	@if [ -f .semgrepignore ]; then mv .semgrepignore .semgrepignore.bak; fi
	docker run --rm -v $(shell pwd):/src returntocorp/semgrep semgrep scan --config=auto /src --verbose
	@if [ -f .semgrepignore.bak ]; then mv .semgrepignore.bak .semgrepignore; fi
appsec-sca:
	docker run --rm -v $(shell pwd):/repo zricethezav/gitleaks:latest detect --source /repo
appsec: appsec-sast appsec-sca
#:[.''.]:>-------------------------------------------



#:[.''.]:>-------------------------------------------
#:[.''.]:><building>
image-push:
	@echo "" && \
	echo "" && \
	echo "Building image markitos-service-boilerplate:$(or $(TAG),1.0.0)" && \
	docker build -t ghcr.io/markitos-es/markitos-service-boilerplate:$(or $(TAG),1.0.0) . && \
	echo "" && \
	echo "Pushing image markitos-service-boilerplate:$(or $(TAG),1.0.0)" && \
	docker push ghcr.io/markitos-es/markitos-service-boilerplate:$(or $(TAG),1.0.0) && \
	echo "" && \
	echo "Image markitos-service-boilerplate:$(or $(TAG),1.0.0) pushed" && \
	docker image rm --force ghcr.io/markitos-es/markitos-service-boilerplate:$(or $(TAG),1.0.0)
	echo "" && \
	echo "Image markitos-service-boilerplate:$(or $(TAG),1.0.0) removed"
	echo "" && \
	echo "Done!"
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:> License: OpenSource :) to pa ti!
#:[.''.]:> markitos.es
#:[.''.]:>-------------------------------------------
