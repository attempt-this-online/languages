FROM attemptthisonline/haskell8

ARG NIBBLES_VERSION=1.01

RUN cd /opt && \
    curl -L http://golfscript.com/nibbles/nibbles-$NIBBLES_VERSION.tgz | \
    tar -xz && \
    cd nibbles && \
    ghc -O -package ghc *.hs
