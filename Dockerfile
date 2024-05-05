FROM golang:1.22-alpine

WORKDIR /app

COPY ./go.mod ./go.sum ./Makefile ./
RUN mkdir ./configs
RUN mkdir ./deploy
RUN go mod download
RUN go install go.uber.org/mock/mockgen@latest
RUN export PATH=$PATH:$(go env GOPATH)/bin
RUN apk update && apk add --no-cache make

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux make build

CMD ["/deploy/feedgram", "--secdist=configs/secdist.yaml", "--config=configs/config.yaml"]
