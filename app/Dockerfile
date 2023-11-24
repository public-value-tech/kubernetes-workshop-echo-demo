FROM golang:1.21 AS builder

WORKDIR /app

ADD . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o echo-server .

FROM scratch

WORKDIR /app
COPY --from=builder /app/echo-server /app/echo-server

CMD [ "/app/echo-server" ]