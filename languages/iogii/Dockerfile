FROM registry.gitlab.pxeger.com/attempt-this-online/languages/haskell8

ARG IOGII_VERSION=1.2

RUN cd /opt && \
    curl -L https://golfscript.com/iogii/$IOGII_VERSION/iogii -o iogii && \
    chmod +x ./iogii && \
    curl -L https://golfscript.com/iogii/$IOGII_VERSION/engine.hs -o engine.hs && \
    ghc -O -package ghc ./engine.hs