#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG FACTOR_VERSION=2023-06-03-00-14

RUN curl -L https://downloads.factorcode.org/linux-x86-64/factor-linux-x86-64-$FACTOR_VERSION.tar.gz | \
    tar -xz && \
    mv factor /opt/
