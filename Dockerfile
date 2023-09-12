FROM golang:1.17 AS builder

RUN go version

COPY . /github.com/salesforceanton/news-tg-bot/
WORKDIR /github.com/salesforceanton/news-tg-bot/

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /github.com/salesforceanton/news-tg-bot/.bin/app .
COPY --from=builder /github.com/salesforceanton/news-tg-bot/configs configs/

EXPOSE 80

CMD ["./app"]