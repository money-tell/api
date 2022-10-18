ifeq ($(shell test -f ./build/.env && echo yes), yes)
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

deps:
	go mod tidy
	go mod vendor

run:
	go run ./cmd/api.go

migrations-create:
	@read -p "Name of the migration: " migration \
	&& echo "Create migrations $$migration at postgres ${MONEY_TELL_POSTGRES_MASTER_DSN}" \
	&& goose -dir migrations postgres "${MONEY_TELL_POSTGRES_MASTER_DSN}" create $$migration sql

migrations-up:
	@goose -dir migrations postgres "${MONEY_TELL_POSTGRES_MASTER_DSN}" up

migrations-down:
	@goose -dir migrations postgres "${MONEY_TELL_POSTGRES_MASTER_DSN}" down

#gowrap gen -p ./app/generated/db/ -i Querier -o app/generated/db/with_prometheus.go -t queries/prometheus.gotmpl
gen-sql:
	@echo "Generate sql..."
	@rm -rf ./app/generated/db
	@mkdir -p ./app/generated/db
	@sqlc generate
	@mockery --dir app/generated/db/ --output app/generated/db/mocks --name Querier

gen: gen-sql


### DOCKER ###
docker-infra-run:
	docker-compose -f build/dev/docker-compose-infra.yml up -d --build

docker-infra-down:
	docker-compose -f build/dev/docker-compose-infra.yml down --remove-orphans -v
