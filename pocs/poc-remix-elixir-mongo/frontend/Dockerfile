FROM node:21-bullseye-slim as base

FROM base as deps
WORKDIR /app
ADD package.json ./
RUN npm install --include=dev


FROM base as production-deps
WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
ADD package.json ./
RUN npm prune --omit=dev
RUN npm update

FROM base as build

WORKDIR /app
COPY --from=deps /app/node_modules ./node_modules
ADD . .
RUN npm run build

FROM base
WORKDIR /app

COPY --from=production-deps /app/node_modules ./node_modules
COPY --from=build /app/build ./build
COPY --from=build /app/public ./public
ADD . .

CMD npm run start