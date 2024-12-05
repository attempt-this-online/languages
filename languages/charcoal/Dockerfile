FROM registry.gitlab.pxeger.com/attempt-this-online/languages/python

ARG CHARCOAL_REV=9d13b5f1c2fed1c9854b403094a53829456232d1

RUN pip install --no-cache-dir --upgrade regex brotli && \
    cd /opt && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.tar.gz | \
    tar -xz && \
    mv Charcoal-$CHARCOAL_REV charcoal
