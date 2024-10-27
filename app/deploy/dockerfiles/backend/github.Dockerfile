FROM golang:1.23-alpine

WORKDIR /app

COPY ./backend ./

RUN go mod tidy

RUN go build -o github cmd/github/main.go

ENV GITHUB_PORT=${GITHUB_PORT}

EXPOSE ${GITHUB_PORT}

# Install MongoDB client (for Alpine-based image)
RUN apk --no-cache add mongodb-tools

# Define environment variables for MongoDB connection
ENV MONGO_HOST=${MONGO_HOST}
ENV MONGO_PORT=${MONGO_PORT}

# Add the HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=5 CMD mongo --host $MONGO_HOST --port $MONGO_PORT --eval "db.adminCommand('ping')" || exit 1

CMD ["sh", "-c", "./github -port $GITHUB_PORT -env-path cmd/github/.env"]

