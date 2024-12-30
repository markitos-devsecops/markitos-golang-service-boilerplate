#!/bin/bash
#:[.''.]:>-------------------------------------------
#:[.''.]:> Author...: Marco Antonio Rubio Lopez
#:[.''.]:> Contact..: markitos.es.info@gmail.com
#:[.''.]:> createdAt: diciembre 2024
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><docker-compose>
docker-up:
	docker compose up -d
docker-down:
	docker compose down
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><runs>
run:
	go run .
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><testing>
test-cache-clear:
	go clean -testcache
test:
	go test ./...
testv:
	go test -v ./...
testc:
	go test ./testsuite/... -cover -coverpkg=./internal/...
testcv:
	go test -v ./testsuite/... -cover -coverpkg=./internal/...
#:[.''.]:>-------------------------------------------

#:[.''.]:>-------------------------------------------
#:[.''.]:><database - migrations - gorm>
createdb:
	docker exec markitos-golang-service-postgres createdb --username=admin --owner=admin markitos-golang-service-boilerplate || true
	docker exec markitos-golang-service-postgres psql -U admin -d markitos-golang-service-boilerplate -c "CREATE USER \"markitos-golang-service-boilerplate\" WITH PASSWORD 'markitos-golang-service-boilerplate';"
	docker exec markitos-golang-service-postgres psql -U admin -d markitos-golang-service-boilerplate -c "GRANT ALL PRIVILEGES ON DATABASE \"markitos-golang-service-boilerplate\" TO \"markitos-golang-service-boilerplate\";"
	docker exec markitos-golang-service-postgres psql -U admin -d markitos-golang-service-boilerplate -c "GRANT ALL PRIVILEGES ON SCHEMA public TO \"markitos-golang-service-boilerplate\";"

dropdb: 
	docker exec markitos-golang-service-postgres psql -U admin -c "SELECT pg_terminate_backend(pg_stat_activity.pid) FROM pg_stat_activity WHERE pg_stat_activity.datname = 'markitos-golang-service-boilerplate' AND pid <> pg_backend_pid();"
	docker exec markitos-golang-service-postgres dropdb -U admin markitos-golang-service-boilerplate
	docker exec markitos-golang-service-postgres psql -U admin -c "DROP USER IF EXISTS \"markitos-golang-service-boilerplate\";"
migrate-init:
	docker run --user $(id -u):$(id -g) -v ./infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ \
		-database "postgresql://admin:admin@localhost:5432/markitos-golang-service-boilerplate?sslmode=disable" create \
		-ext sql -dir infrastructure/database/migrations/ -seq init_schema 
migrate-up:
	docker run --user $(id -u):$(id -g) -v ./infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@localhost:5432/markitos-golang-service-boilerplate?sslmode=disable" -verbose up		
migrate-down:
	docker run --user $(id -u):$(id -g) -v ./infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@localhost:5432/markitos-golang-service-boilerplate?sslmode=disable" -verbose down $(or $(VERSION),1)
migrate-version:
	docker run --user $(id -u):$(id -g) -v ./infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@localhost:5432/markitos-golang-service-boilerplate?sslmode=disable" version
migrate-goto:
	docker run --user $(id -u):$(id -g) -v ./infrastructure/database/migrations:/migrations --network host migrate/migrate \
		-path=/migrations/ -database "postgresql://admin:admin@localhost:5432/markitos-golang-service-boilerplate?sslmode=disable" goto $(or $(VERSION),1)
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
docker-login:
	docker login ghcr.io -u markitos-es --password $(GITHUB_TOKEN)
docker-publish-tag: docker-login
	@echo "" && \
	echo "" && \
	echo "Building image markitos-golang-service-boilerplate:$(or $(TAG),1.0.0)" && \
	docker build -t ghcr.io/markitos-devsecops/markitos-golang-service-boilerplate:$(or $(TAG),1.0.0) . && \
	echo "" && \
	echo "Pushing image markitos-golang-service-boilerplate:$(or $(TAG),1.0.0)" && \
	docker push ghcr.io/markitos-devsecops/markitos-golang-service-boilerplate:$(or $(TAG),1.0.0) && \
	echo "" && \
	echo "Image markitos-golang-service-boilerplate:$(or $(TAG),1.0.0) pushed" && \
	docker image rm --force ghcr.io/markitos-devsecops/markitos-golang-service-boilerplate:$(or $(TAG),1.0.0)
	echo "" && \
	echo "Image markitos-golang-service-boilerplate:$(or $(TAG),1.0.0) removed"
	echo "" && \
	echo "Done!"
docker-publish-postgres: docker-login
	@echo "" && \
	echo "" && \
	echo "Building image markitos-golang-service-postgres:$(or $(TAG),1.0.0)" && \
	docker build --file Dockerfile.postgres -t ghcr.io/markitos-devsecops/markitos-golang-service-postgres:$(or $(TAG),1.0.0) . && \
	echo "" && \
	echo "Pushing image markitos-golang-service-postgres:$(or $(TAG),1.0.0)" && \
	docker push ghcr.io/markitos-devsecops/markitos-golang-service-postgres:$(or $(TAG),1.0.0) && \
	echo "" && \
	echo "Image markitos-golang-service-postgres:$(or $(TAG),1.0.0) pushed" && \
	docker image rm --force ghcr.io/markitos-devsecops/markitos-golang-service-postgres:$(or $(TAG),1.0.0)
	echo "" && \
	echo "Image markitos-golang-service-postgres:$(or $(TAG),1.0.0) removed"
	echo "" && \
	echo "Done!"
#:[.''.]:>-------------------------------------------


#:[.''.]:>-------------------------------------------
#:[.''.]:> License: OpenSource :) to pa ti!
#:[.''.]:> markitos.es
#:[.''.]:>-------------------------------------------
