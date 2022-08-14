#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/erlang

ARG ELIXIR_VERSION=1.13.4

RUN curl -L https://github.com/elixir-lang/elixir/releases/download/v$ELIXIR_VERSION/Precompiled.zip -o elixir.zip && \
    mkdir /opt/elixir && \
    cd /opt/elixir && \
    unzip /tmp/elixir.zip && \
    rm -rf /tmp/elixir.zip
