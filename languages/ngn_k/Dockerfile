FROM attemptthisonline/base

ARG K_REV=de8bc7db10b0c5944e5491c3aa968dea0d8b8d41

RUN curl -L https://codeberg.org/ngn/k/archive/$K_REV.tar.gz | \
    tar -xz && \
    cd k && \
    make -j $(nproc) CC=gcc && \
    mv k /usr/local/bin/ && \
    cd / && \
    rm -rf /tmp/k
