FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG NIM_VERSION=2.2.6

RUN curl -L https://nim-lang.org/download/nim-$NIM_VERSION-linux_x64.tar.xz | \
    tar -xJ && \
    mv nim-$NIM_VERSION /opt/nim
