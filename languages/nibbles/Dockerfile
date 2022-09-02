#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/haskell8

ARG NIBBLES_VERSION=1.00

RUN cd /opt && \
    curl -L http://golfscript.com/nibbles/nibbles-$NIBBLES_VERSION.tgz | \
    tar -xz && \
    cd nibbles && \
    rm all.hs && \
    ghc -O -package ghc *.hs
