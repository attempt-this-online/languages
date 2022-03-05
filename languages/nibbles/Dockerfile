#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM haskell:8 AS build

WORKDIR /tmp

ARG NIBBLES_VERSION=0.25

RUN curl -L http://golfscript.com/nibbles/nibbles-$NIBBLES_VERSION.tgz | \
    tar -xz && \
    cd nibbles && \
    ghc -O -package ghc *.hs && \
    mv ./nibbles /

FROM attemptthisonline/base

COPY --from=build /nibbles /usr/local/bin/
