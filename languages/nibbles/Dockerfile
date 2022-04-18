#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/haskell8

ARG NIBBLES_VERSION=0.25

COPY out.patch /tmp/

RUN cd /opt && \
    curl -L http://golfscript.com/nibbles/nibbles-$NIBBLES_VERSION.tgz | \
    tar -xz && \
    cd nibbles && \
    patch < /tmp/out.patch && \
    rm /tmp/out.patch && \
    ghc -O -package ghc *.hs
