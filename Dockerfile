FROM golang:1.22-alpine

WORKDIR /app

COPY ./go.mod ./go.sum ./
RUN mkdir ./configs
RUN mkdir ./deploy
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -C ./cmd/app -o /deploy

CMD ["/deploy", "--secdist=configs/secdist.yaml", "--config=configs/config.yaml"]
