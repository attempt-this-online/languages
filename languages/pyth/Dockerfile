FROM registry.gitlab.pxeger.com/attempt-this-online/languages/python

ARG PYTH_REV=bb7c848818caa96932a058ac047ef3b98f3527bf

RUN curl -L https://github.com/attempt-this-online/pyth/archive/$PYTH_REV.tar.gz | \
    tar -xz && \
    mv pyth-$PYTH_REV /opt/pyth
