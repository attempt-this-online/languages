#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG J_VERSION=903

COPY setup /tmp/

RUN curl -L http://www.jsoftware.com/download/j$J_VERSION/install/j${J_VERSION}_linux64.tar.gz | \
    tar -xz && \
    mv j$J_VERSION /opt/j && \
    cd /opt/j && \
    ./jconsole.sh /tmp/setup && \
    rm -rf /tmp/*
