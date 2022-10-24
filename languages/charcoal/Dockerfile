#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG CHARCOAL_REV=b84c6c5249b3cfafa7082bf7943de9bbff188660

RUN pip install --no-cache-dir --upgrade regex brotli && \
    cd /opt && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.tar.gz | \
    tar -xz && \
    mv Charcoal-$CHARCOAL_REV charcoal
