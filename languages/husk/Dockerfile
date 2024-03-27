FROM attemptthisonline/haskell8

ARG HUSK_REV=03b3cf96e93824075cc70cef4bb3fd7242bd2ad0

RUN curl -L https://github.com/barbuz/Husk/archive/$HUSK_REV.tar.gz | \
    tar -xz && \
    mv Husk-$HUSK_REV /opt/husk && \
    cd /opt/husk && \
    ghc Husk.hs -o husk
