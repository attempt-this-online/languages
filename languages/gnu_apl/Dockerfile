#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG APL_VERSION=1.8

RUN curl -L https://mirrors.gethosted.online/gnu/apl/apl-$APL_VERSION.tar.gz | \
    tar -xz && \
    cd apl-$APL_VERSION && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    cd / && \
    rm -rf /tmp/apl-$APL_VERSION
