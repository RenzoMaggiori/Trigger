FROM golang:1.23-alpine

WORKDIR /app

COPY ./backend ./

RUN go mod tidy

RUN go build -o discord cmd/discord/main.go

ENV DISCORD_PORT=${DISCORD_PORT}

EXPOSE ${DISCORD_PORT}

# Install MongoDB client (for Alpine-based image)
RUN apk --no-cache add mongodb-tools

# Define environment variables for MongoDB connection
ENV MONGO_HOST=${MONGO_HOST}
ENV MONGO_PORT=${MONGO_PORT}

# Add the HEALTHCHECK instruction
HEALTHCHECK --interval=30s --timeout=10s --start-period=5s --retries=5 CMD mongo --host $MONGO_HOST --port $MONGO_PORT --eval "db.adminCommand('ping')" || exit 1

CMD ["sh", "-c", "./discord -port $DISCORD_PORT -env-path cmd/discord/.env"]
