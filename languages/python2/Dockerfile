FROM attemptthisonline/base

ARG PYTHON_VERSION=2.7.18

RUN curl -L https://www.python.org/ftp/python/$PYTHON_VERSION/Python-$PYTHON_VERSION.tar.xz | \
    tar -xJ && \
    cd Python-$PYTHON_VERSION && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    cd / && \
    rm -rf /tmp/Python-$PYTHON_VERSION
