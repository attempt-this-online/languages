FROM attemptthisonline/base

ARG ZSH_VERSION=5.8

RUN curl https://www.zsh.org/pub/zsh-$ZSH_VERSION.tar.xz | tar -xJ && \
    curl https://www.zsh.org/pub/zsh-$ZSH_VERSION-doc.tar.xz | tar -xJ && \
    cd zsh-$ZSH_VERSION && \
    ./configure --with-tcsetpgrp && \
    make && \
    make install && \
    cd / && \
    rm -rf /tmp/zsh-$ZSH_VERSION
