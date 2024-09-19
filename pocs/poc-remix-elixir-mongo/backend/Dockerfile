# Elixir + Phoenix

FROM elixir:1.17.2
# Install debian packages
RUN apt-get update
RUN apt-get install --yes build-essential inotify-tools

# Install Phoenix packages
RUN mix local.hex --force
RUN mix local.rebar --force

# Install node
# RUN curl -sL https://deb.nodesource.com/setup_6.x -o nodesource_setup.sh
# RUN bash nodesource_setup.sh
# RUN apt-get install --yes nodejs npm

WORKDIR /app

CMD ["./run.sh"]