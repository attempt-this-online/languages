FROM registry.gitlab.pxeger.com/attempt-this-online/languages/python

ARG CHARCOAL_REV=796761f03b9d4861eb07c4a505c414d7fd747c30

RUN pip install --no-cache-dir --upgrade regex brotli && \
    cd /opt && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.tar.gz | \
    tar -xz && \
    mv Charcoal-$CHARCOAL_REV charcoal
