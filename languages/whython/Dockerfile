FROM attemptthisonline/base

ARG WHYTHON_VERSION=3.10.0y2

RUN curl -L https://github.com/pxeger/whython/archive/refs/tags/v$WHYTHON_VERSION.tar.gz | \
    tar -xz && \
    cd whython-$WHYTHON_VERSION && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    ln -s python3 /usr/local/bin/python && \
    ln -s pip3 /usr/local/bin/pip && \
    cd / && \
    rm -rf /tmp/whython-$WHYTHON_VERSION

# separate step so Python doesn't need to be rebuilt every time a pip thing is added or updated
RUN pip install --upgrade pip setuptools wheel
