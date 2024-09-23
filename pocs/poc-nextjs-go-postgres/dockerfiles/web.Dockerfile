FROM node:18-alpine AS base

# Install dependencies only when needed
FROM base AS deps
# Check https://github.com/nodejs/docker-node/tree/b4117f9333da4138b03a546ec926ef50a31506c3#nodealpine to understand why libc6-compat might be needed.
RUN apk add --no-cache libc6-compat
WORKDIR /app

# Install dependencies based on the preferred package manager
COPY ./frontend/package.json ./frontend/package-lock.json* ./

RUN npm ci

FROM base AS runner

WORKDIR /app

COPY --from=deps /app/node_modules ./node_modules
COPY ./frontend .

EXPOSE 3000

ENV PORT=3000

CMD ["npm", "run", "dev"]
