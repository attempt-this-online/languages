FROM attemptthisonline/base

ARG ZSH_VERSION=5.8.1

RUN curl -L https://www.zsh.org/pub/zsh-$ZSH_VERSION.tar.xz | tar -xJ && \
    cd zsh-$ZSH_VERSION && \
    ./configure \
        --with-tcsetpgrp \
        --enable-pcre \
        && \
    make && \
    make install && \
    cd / && \
    rm -rf /tmp/zsh-$ZSH_VERSION
