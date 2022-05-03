#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG SBCL_VERSION=2.2.1

RUN curl -L https://kumisystems.dl.sourceforge.net/project/sbcl/sbcl/$SBCL_VERSION/sbcl-$SBCL_VERSION-x86-64-linux-binary.tar.bz2 | \
    tar -xj && \
    cd sbcl-$SBCL_VERSION-x86-64-linux && \
    ./install.sh
