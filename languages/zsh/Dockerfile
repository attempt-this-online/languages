FROM debian:buster

WORKDIR /tmp

RUN apt-get update && \
    apt-get upgrade -y && \
    apt-get install -y \
        wget \
        gcc \
        xz-utils \
        libncurses-dev \
        make \
        && \
    wget https://www.zsh.org/pub/zsh-5.8.tar.xz && \
    wget https://www.zsh.org/pub/zsh-5.8-doc.tar.xz && \
    tar -xf zsh-5.8.tar.xz && \
    tar -xf zsh-5.8-doc.tar.xz && \
    cd zsh-5.8 && \
    ./configure --with-tcsetpgrp && \
    make && \
    make install

ENTRYPOINT /usr/local/bin/zsh
