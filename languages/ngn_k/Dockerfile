#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/clang

ARG K_REV=ba2c2bab7e533e8fdc4e4cc2ee0c0c939990d5f5

RUN curl -L https://codeberg.org/ngn/k/archive/$K_REV.tar.gz | \
    tar -xz && \
    cd k && \
    sed -i 's/march=native/march=znver1/g' makefile && \
    make -j $(nproc) CC=gcc && \
    mv k /usr/local/bin/ && \
    cd / && \
    rm -rf /tmp/k
