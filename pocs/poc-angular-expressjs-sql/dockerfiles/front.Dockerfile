FROM node:18.19.1 AS build

WORKDIR /app

COPY ./frontend/todo-list/package*.json ./

RUN npm install

RUN npm install -g @angular/cli

COPY ./frontend/todo-list/ ./

RUN npm run build

EXPOSE 4200

CMD ["npm", "run", "start", "--", "--host", "0.0.0.0"]