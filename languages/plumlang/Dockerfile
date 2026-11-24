FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG PLUMLANG_REV=6ef73cdd01a7ac9a1cbe282d4b3378ea935ac74c

RUN mkdir plumlang && \
    curl -L https://github.com/plumlang/plumlang/archive/$PLUMLANG_REV.tar.gz | \
    tar -xz --strip-components=1 -C plumlang && \
    cd plumlang && \
    make && \
    install -m 755 plum /usr/local/bin/plum && \
    cd / && \
    rm -rf /tmp/plumlang
