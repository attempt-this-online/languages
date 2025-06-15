FROM registry.gitlab.pxeger.com/attempt-this-online/languages/node

ARG KNIGHT_VERSION=199c203aa8e4b7917db5a8ca560ad11bda4586f5

RUN cd /opt && \
    curl -L https://github.com/knight-lang/javascript/archive/$KNIGHT_VERSION.tar.gz | \
    tar -xz && \
    mv javascript-$KNIGHT_VERSION knight
