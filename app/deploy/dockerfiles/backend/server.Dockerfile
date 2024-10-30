FROM nginx:latest

COPY ./backend/config/nginx.template.conf /etc/nginx/nginx.template.conf

COPY ./deploy/entrypoints/server.sh /usr/local/bin/docker-entrypoint.sh

ENV DOLLAR="$"

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

EXPOSE 80

ENTRYPOINT ["/usr/local/bin/docker-entrypoint.sh"]

CMD ["nginx", "-g", "daemon off;"]
