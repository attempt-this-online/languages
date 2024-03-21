#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG PIP_VERSION=v1.2.0

RUN curl -L https://github.com/dloscutoff/pip/archive/refs/tags/$PIP_VERSION.tar.gz | \
    tar -xz && \
    mv -T pip-* /opt/pip && \
    chmod +x /opt/pip/pip.py && \
    ln -s /opt/pip/pip.py /usr/local/bin/piplang
