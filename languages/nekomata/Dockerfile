#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/haskell

ARG NEKOMATA_VERSION=0.4.0.0

RUN curl -L https://github.com/AlephAlpha/Nekomata/archive/refs/tags/v$NEKOMATA_VERSION.tar.gz | \
    tar -xz && \
    cd Nekomata-$NEKOMATA_VERSION && \
    cabal --store-dir /opt/cabal install --package-env /opt/ghc_env --install-method=copy --installdir /usr/local/bin && \
    cd / && \
    rm -rf /tmp/Nekomata-$NEKOMATA_VERSION
