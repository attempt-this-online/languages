FROM registry.gitlab.pxeger.com/attempt-this-online/languages/haskell

ARG WSPACE_REV=5eb78c6afae9f5dd2d2012a1784f6483ec824135

RUN curl -L https://github.com/TryItOnline/WSpace/archive/$WSPACE_REV.tar.gz | \
    tar -xz && \
    cd WSpace-$WSPACE_REV && \
    make && \
    mv wspace /usr/local/bin && \
    cd / && \
    rm -rf /tmp/WSpace-*
