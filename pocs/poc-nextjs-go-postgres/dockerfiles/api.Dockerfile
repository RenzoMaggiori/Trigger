FROM golang:1.22-alpine

WORKDIR /app

COPY ./backend/api ./api

COPY ./scripts ./scripts

WORKDIR /app/api

RUN go mod tidy

RUN go build -o api .

WORKDIR /app

CMD ["./scripts/api.sh"]

