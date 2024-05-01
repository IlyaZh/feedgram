build:
	go build -C ./cmd/app/ -o ./../../deploy/feedgram

run: build
	export ENVIRONMENT=production && go run ./cmd/app/

run-debug: build
	export ENVIRONMENT=debug && go run ./cmd/app/

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down

start: docker-start run

stop: docker-stop

