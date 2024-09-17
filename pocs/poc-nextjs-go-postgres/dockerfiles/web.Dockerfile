FROM node:18-alpine

RUN apk add --no-cache g++ make py3-pip libc6-compat

WORKDIR /app

COPY ./frontend/package*.json .

EXPOSE 3000

WORKDIR /app

COPY ./frontend ./frontend

RUN npm run build

RUN npm ci

RUN addgroup -g 1001 -S nodejs

RUN adduser -S nextjs -u 1001

USER nextjs

ENV NODE_ENV=production

CMD npm start

