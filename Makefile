build: gen
	go build -C ./cmd/app/ -o ./../../deploy/feedgram

run: build
	export ENVIRONMENT=production && go run ./cmd/app/

run-debug: build
	export ENVIRONMENT=debug && go run ./cmd/app/

test:
	go test ./...

docker-start:
	docker-compose up -d

docker-stop:
	docker-compose down

docker-build: 
	docker build --tag=ilyazh/feedgram:latest .

start: docker-start

stop: docker-stop

mock-gen:
	mkdir -p ./internal/caches/configs/mocks; \
	mockgen -source ./internal/caches/configs/cache.go -package mocks -destination ./internal/caches/configs/mocks/component.go; \
	\
	mkdir -p ./internal/components/message_dispatcher/mocks; \
	mockgen -source ./internal/components/message_dispatcher/component.go -package mocks -destination ./internal/components/message_dispatcher/mocks/component.go; \
	\
	mkdir -p ./internal/components/message_sender/mocks; \
	mockgen -source ./internal/components/message_sender/component.go -package mocks -destination ./internal/components/message_sender/mocks/component.go; \
	\
	mkdir -p ./internal/components/news_checker/mocks; \
	mockgen -source ./internal/components/news_checker/component.go -package mocks -destination ./internal/components/news_checker/mocks/component.go; \
	\
	mkdir -p ./internal/components/rss_reader/mocks; \
	mockgen -source ./internal/components/rss_reader/component.go -package mocks -destination ./internal/components/rss_reader/mocks/component.go; \
	\
	mkdir -p ./internal/components/storage/mocks; \
	mockgen -source ./internal/components/storage/component.go -package mocks -destination ./internal/components/storage/mocks/component.go; \
	\
	mkdir -p ./internal/components/telegram/mocks; \
	mockgen -source ./internal/components/telegram/component.go -package mocks -destination ./internal/components/telegram/mocks/component.go \ 
	mockgen -source ./internal/components/telegram/tg_api.go -package mocks -destination ./internal/components/telegram/mocks/tg_api.go

gen: mock-gen

