#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG STACKAGE_VERSION=lts-18.28

RUN curl -L https://downloads.haskell.org/~ghcup/x86_64-linux-ghcup -o ghcup && \
    chmod +x ghcup && \
    export GHCUP_INSTALL_BASE_PREFIX=/opt/ && \
    ./ghcup install ghc 8 --set && \
    ./ghcup install stack latest && \
    ./ghcup gc --cache --profiling-libs --tmpdirs --share-dir && \
    rm ./ghcup

ENV PATH=/opt/.ghcup/bin:$PATH 

RUN stack update --resolver $STACKAGE_VERSION && \
    stack config set install-ghc false --global && \
    stack config set system-ghc --global true && \
    stack install --resolver $STACKAGE_VERSION regex random
