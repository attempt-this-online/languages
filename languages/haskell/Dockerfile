#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

RUN curl -L https://downloads.haskell.org/~ghcup/x86_64-linux-ghcup -o ghcup && \
    chmod +x ghcup && \
    export GHCUP_INSTALL_BASE_PREFIX=/opt/ && \
    ./ghcup install ghc 9.0 --set && \
    ./ghcup install stack latest && \
    ./ghcup gc --cache --profiling-libs --tmpdirs --share-dir && \
    rm ./ghcup

ENV PATH=/opt/.ghcup/bin:$PATH 

RUN stack update && \
    stack config set install-ghc false --global && \
    stack config set system-ghc --global true && \
    stack install --resolver nightly regex random heredoc
