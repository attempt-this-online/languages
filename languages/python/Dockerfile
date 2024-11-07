FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG PYTHON_VERSION=3.13.0

RUN curl -L https://www.python.org/ftp/python/$PYTHON_VERSION/Python-$PYTHON_VERSION.tar.xz | \
    tar -xJ && \
    cd Python-$PYTHON_VERSION && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    ln -s python3 /usr/local/bin/python && \
    ln -s pip3 /usr/local/bin/pip && \
    cd / && \
    rm -rf /tmp/Python-$PYTHON_VERSION

# separate step so Python doesn't need to be rebuilt every time a pip thing is added or updated
RUN pip install --upgrade pip setuptools wheel
