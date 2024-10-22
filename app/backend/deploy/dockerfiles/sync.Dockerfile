FROM golang:1.23-alpine

WORKDIR /app

COPY ./ ./

RUN go mod tidy

RUN go build -o sync cmd/sync/main.go

ENV SYNC_PORT=${SYNC_PORT}

EXPOSE ${SYNC_PORT}

# Install MongoDB client (for Alpine-based image)
RUN apk --no-cache add mongodb-tools

# Define environment variables for MongoDB connection
ENV MONGO_HOST=${MONGO_HOST}
ENV MONGO_PORT=${MONGO_PORT}

# Add the HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=5 CMD mongo --host $MONGO_HOST --port $MONGO_PORT --eval "db.adminCommand('ping')" || exit 1

CMD ["sh", "-c", "./sync -port $SYNC_PORT -env-path cmd/sync/.env"]

