build:
	go build ./cmd/app/

run: build
	go run ./cmd/app/

gen:
	go run github.com/deepmap/oapi-codegen/cmd/oapi-codegen --config ./scripts/codegen_public_api.yaml ./api/public.yaml

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down

start: docker-start run


stop: docker-stop

