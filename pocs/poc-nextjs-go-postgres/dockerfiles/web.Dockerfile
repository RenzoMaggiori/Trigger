FROM node:18-alpine as base

RUN apk add --no-cache g++ make py3-pip libc6-compat

WORKDIR /app

COPY package*.json .

EXPOSE 3000

# Production
FROM base as production

WORKDIR /app

COPY . .

RUN npm run build
RUN npm ci
RUN addgroup -g 1001 -S nodejs
RUN adduser -S nextjs -u 1001

USER nextjs

ENV NODE_ENV=production

CMD npm start

# Development
FROM base as dev

RUN npm i

COPY . .

ENV NODE_ENV=development

CMD npm run dev
