FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG PICAT_VERSION=368

RUN curl http://picat-lang.org/download/picat${PICAT_VERSION}_linux64.tar.gz | \
    tar -xz && \
    mv Picat /opt/picat
