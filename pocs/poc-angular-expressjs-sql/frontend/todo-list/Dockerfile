FROM node:21.6.0 AS dev

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

RUN npm install -g @angular/cli

COPY . ./

EXPOSE 4200

CMD ["npm", "run", "start", "--", "--host", "0.0.0.0"]

FROM node:21.6.0 AS build

WORKDIR /app

COPY package.json package-lock.json ./
RUN npm install

RUN npm install -g @angular/cli

COPY . ./

RUN npm run build --configuration=production

FROM nginx:latest AS prod

COPY --from=build /app/dist/todo-list /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
