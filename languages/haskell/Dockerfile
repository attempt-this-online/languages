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

ARG HGL_REV=327f42477fbb9d9b4223c8689b06e4d24a536165

COPY hgl_base_4_18.patch /tmp/

RUN cabal --store-dir /opt/cabal update && \
    cabal --store-dir /opt/cabal install --lib --package-env /opt/ghc_env \
        base ghc ghc-boot template-haskell ghc-boot-th ghc-prim ghci \
        unix terminfo pretty time filepath directory process setenv \
        text binary array bytestring containers unordered-containers hashable integer-gmp utf8-string \
        transformers mtl semigroupoids free bifunctors profunctors comonad distributive primitive \
        deepseq hpc stm call-stack random regex heredoc hspec QuickCheck colour && \
    curl -L https://gitlab.com/WheatWizard/haskell-golfing-library/-/archive/$HGL_REV/haskell-golfing-library-$HGL_REV.tar.gz | \
    tar -xz && \
    cd haskell-golfing-library-* && \
    patch < /tmp/hgl_base_4_18.patch src/P/Functor.hs && \
    cabal --store-dir /opt/cabal install --lib --package-env /opt/ghc_env . && \
    cd / && \
    rm -rf /tmp/*
