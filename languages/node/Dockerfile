#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG NODE_VERSION=v18.8.0

RUN curl https://nodejs.org/dist/$NODE_VERSION/node-$NODE_VERSION-linux-x64.tar.xz | \
    tar -xJ && \
    cd node-$NODE_VERSION-linux-x64 && \
    mv bin include lib /usr/local/ && \
    cd / && \
    rm -rf /tmp/node-$NODE_VERSION-linux-x64
