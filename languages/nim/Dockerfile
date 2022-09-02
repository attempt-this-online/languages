#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG NIM_VERSION=1.6.6

RUN curl -L https://nim-lang.org/download/nim-$NIM_VERSION-linux_x64.tar.xz | \
    tar -xJ && \
    mv nim-$NIM_VERSION /opt/nim
