FROM attemptthisonline/clang

ARG K_REV=ba2c2bab7e533e8fdc4e4cc2ee0c0c939990d5f5

RUN curl -L https://codeberg.org/ngn/k/archive/$K_REV.tar.gz | \
    tar -xz && \
    cd k && \
    make -j $(nproc) CC=clang && \
    mv k /usr/local/bin/ && \
    cd / && \
    rm -rf /tmp/k
