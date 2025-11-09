FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG TCL_VERSION=9.0.2

RUN curl -L http://prdownloads.sourceforge.net/tcl/tcl$TCL_VERSION-src.tar.gz | \
    tar -xz && \
    cd tcl* && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    cd / && \
    rm -rf /tmp/tcl*
