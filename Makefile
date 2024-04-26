build:
	go build -C ./cmd/app/ -o ./../../deploy/feedgram

run: build
	go run ./cmd/app/

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down

start: docker-start run

stop: docker-stop

