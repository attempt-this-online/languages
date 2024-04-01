FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG SBCL_VERSION=2.2.8

RUN curl -L https://sourceforge.net/projects/sbcl/files/sbcl/$SBCL_VERSION/sbcl-$SBCL_VERSION-x86-64-linux-binary.tar.bz2/download | \
    tar -xj && \
    cd sbcl-$SBCL_VERSION-x86-64-linux && \
    ./install.sh
