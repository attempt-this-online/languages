FROM debian:buster

WORKDIR /tmp

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y default-jre \
    apt-get install -y scala

ENTRYPOINT /usr/local/bin/scala
