FROM debian:buster

WORKDIR /tmp

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
        wget \
        python3 \
        python-pip3 \
        && \
    wget https://github.com/DennisMitchell/jellylanguage/archive/master.tar.gz && \
    pip3 install --upgrade https://github.com/DennisMitchell/jellylanguage/archive/master.tar.gz

ENTRYPOINT /usr/local/bin/jelly
