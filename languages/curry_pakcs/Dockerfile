#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/swi_prolog

ARG PAKCS_VERSION=3.3.0

RUN curl -L https://www.informatik.uni-kiel.de/~pakcs/download/pakcs-$PAKCS_VERSION-amd64-Linux.tar.gz | \
    tar -xzC /opt && \
    cd /opt/pakcs-$PAKCS_VERSION && \
    make -j $(nproc) && \
    ln -s /opt/pakcs-$PAKCS_VERSION/bin/pakcs /usr/local/bin/
