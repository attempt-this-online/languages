FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG ZSH_VERSION=5.9

RUN curl -L https://www.zsh.org/pub/zsh-$ZSH_VERSION.tar.xz | tar -xJ && \
    cd zsh-$ZSH_VERSION && \
    ./configure \
        --with-tcsetpgrp \
        --enable-pcre \
        && \
    make CC="gcc -DHAVE_BOOLCODES" && \
    make install && \
    cd / && \
    rm -rf /tmp/zsh-$ZSH_VERSION
