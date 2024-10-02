
FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG PARIGP_VERSION=2.17.0

RUN curl -L https://pari.math.u-bordeaux.fr/pub/pari/unix/pari-$PARIGP_VERSION.tar.gz | \
    tar -xz && \
    cd pari-$PARIGP_VERSION && \
    ./Configure --mt=pthread && \
    make all -j $(nproc) && \
    make install && \
    cd / && \
    rm -rf /tmp/pari-$PARIGP_VERSION
