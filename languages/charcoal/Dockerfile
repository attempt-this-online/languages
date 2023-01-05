#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG CHARCOAL_REV=49787e5e9011153d0dc26c712184603ae0c39641

RUN pip install --no-cache-dir --upgrade regex brotli && \
    cd /opt && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.tar.gz | \
    tar -xz && \
    mv Charcoal-$CHARCOAL_REV charcoal
