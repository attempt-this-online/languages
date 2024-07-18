FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

RUN curl -L https://downloads.haskell.org/~ghcup/x86_64-linux-ghcup -o ghcup && \
    chmod +x ghcup && \
    export GHCUP_INSTALL_BASE_PREFIX=/opt/ && \
    ./ghcup install ghc 9.4 --set && \
    ./ghcup install cabal recommended && \
    ./ghcup gc --cache --profiling-libs --tmpdirs --share-dir && \
    rm ./ghcup

ENV PATH=/opt/.ghcup/bin:$PATH

ARG HGL_REV=702406d06b681ab6930245b4e2b5f75ee43dd389

COPY hgl_base_4_17.patch /tmp/

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
    patch < /tmp/hgl_base_4_17.patch src/P/Functor.hs && \
    cabal --store-dir /opt/cabal install --lib --package-env /opt/ghc_env . && \
    cd / && \
    rm -rf /tmp/*
