FROM registry.gitlab.pxeger.com/attempt-this-online/languages/clang

ARG K_REV=a1dac7940934fd23d403b462d6f0ee09853d55a4

RUN curl -L https://codeberg.org/ngn/k/archive/$K_REV.tar.gz | \
    tar -xz && \
    cd k && \
    sed -i 's/march=native/march=znver1/g' makefile && \
    make -j $(nproc) CC=gcc && \
    mv k /usr/local/bin/ && \
    cd / && \
    rm -rf /tmp/k
