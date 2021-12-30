FROM attemptthisonline/python

ARG PIP_VERSION=v1.0.2

RUN curl -L https://github.com/dloscutoff/pip/archive/refs/tags/$PIP_VERSION.tar.gz | \
    tar -xz && \
    mv -T pip-* /opt/pip && \
    ln -s /opt/pip/pip.py /usr/local/bin/piplang
