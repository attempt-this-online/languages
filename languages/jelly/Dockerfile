FROM python:3-buster

WORKDIR /tmp

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
        wget \
        && \
    wget https://github.com/DennisMitchell/jellylanguage/archive/70c9fd93ab009c05dc396f8cc091f72b212fb188.tar.gz && \
    pip3 install https://github.com/DennisMitchell/jellylanguage/archive/70c9fd93ab009c05dc396f8cc091f72b212fb188.tar.gz

ENTRYPOINT /usr/local/bin/jelly
