# soft docker
FROM golang:1.19-alpine

RUN go version
WORKDIR /hakaton-bot-api
COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./

RUN go build -o hakaton-bot-api ./cmd/main.go
# EXPOSE 8081
CMD ["./hakaton-bot-api"]