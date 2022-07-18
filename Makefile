ifeq ($(shell test -f ./build/.env && echo yes), yes)
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif


deps:
	go mod tidy
	go mod vendor

run:
	go run ./cmd/api.go

### DOCKER ###
docker-infra-run:
	docker-compose -f build/dev/docker-compose-infra.yml up -d --build

docker-infra-down:
	docker-compose -f build/dev/docker-compose-infra.yml down --remove-orphans -v
