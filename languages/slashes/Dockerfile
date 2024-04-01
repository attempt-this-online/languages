FROM registry.gitlab.pxeger.com/attempt-this-online/languages/perl

ARG SLASHES_REV=c3c84c8d98f5dbefe7d1b690c728dac3e5db1131

RUN curl -L "https://github.com/attempt-this-online/slashes/archive/$SLASHES_REV.tar.gz" | \
    tar -xz && \
    mv "slashes-$SLASHES_REV" /opt/slashes
