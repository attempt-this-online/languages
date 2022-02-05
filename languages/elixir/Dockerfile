FROM attemptthisonline/base

ARG ELIXIR_VERSION=1.13.2

RUN curl -L https://github.com/elixir-lang/elixir/releases/download/v$ELIXIR_VERSION/Precompiled.zip -o elixir.zip && \
    mkdir /opt/elixir && \
    cd /opt/elixir && \
    unzip /tmp/elixir.zip && \
    rm -rf /tmp/elixir.zip
