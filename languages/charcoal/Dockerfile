FROM registry.gitlab.pxeger.com/attempt-this-online/languages/python

ARG CHARCOAL_REV=c939a5e76f0470b8338200a54d41897e8bf6f69d

RUN pip install --no-cache-dir --upgrade regex brotli && \
    cd /opt && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.tar.gz | \
    tar -xz && \
    mv Charcoal-$CHARCOAL_REV charcoal
