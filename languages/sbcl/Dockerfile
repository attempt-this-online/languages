FROM attemptthisonline/base

ARG SBCL_VERSION=2.2.1

RUN curl -L http://prdownloads.sourceforge.net/sbcl/sbcl-$SBCL_VERSION-x86-64-linux-binary.tar.bz2 | \
    tar -xj && \
    cd sbcl-$SBCL_VERSION-x86-64-linux && \
    ./install.sh
