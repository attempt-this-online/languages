#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

RUN curl -L https://downloads.haskell.org/~ghcup/x86_64-linux-ghcup -o ghcup && \
    chmod +x ghcup && \
    export GHCUP_INSTALL_BASE_PREFIX=/opt/ && \
    ./ghcup install ghc latest --set && \
    ./ghcup install cabal latest && \
    ./ghcup gc --cache --profiling-libs --tmpdirs --share-dir && \
    rm ./ghcup

ENV PATH=/opt/.ghcup/bin:$PATH 

ARG HGL_REV=09621e4a3856680eb510464e71b9a3693f19a376

RUN cabal --store-dir /opt/cabal update && \
    cabal --store-dir /opt/cabal install --lib --package-env /opt/ghc_env \
        random regex heredoc free hspec QuickCheck && \
    curl -L https://gitlab.com/WheatWizard/haskell-golfing-library/-/archive/$HGL_REV/haskell-golfing-library-$HGL_REV.tar.gz | \
    tar -xz && \
    cd haskell-golfing-library-* && \
    cabal --store-dir /opt/cabal install --lib --package-env /opt/ghc_env . && \
    cd / && \
    rm -rf /tmp/*
