FROM golang:1.23-alpine

# Install MongoDB client (for Alpine-based image)
RUN apk --no-cache add mongodb-tools

WORKDIR /app

COPY ./backend ./

RUN go mod tidy

RUN go build -o bitbucket cmd/bitbucket/main.go

ENV BITBUCKET_PORT=${BITBUCKET_PORT}

EXPOSE ${BITBUCKET_PORT}

# Define environment variables for MongoDB connection
ENV MONGO_HOST=${MONGO_HOST}
ENV MONGO_PORT=${MONGO_PORT}

# Add the HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=5 CMD mongo --host $MONGO_HOST --port $MONGO_PORT --eval "db.adminCommand('ping')" || exit 1

CMD ["sh", "-c", "./bitbucket -port $BITBUCKET_PORT -env-path cmd/bitbucket/.env"]

