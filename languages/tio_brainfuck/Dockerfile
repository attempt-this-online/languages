FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG BRAINFUCK_REV=9ef9a92460bae60188ce45a0ac64c8fb8d0898f8

RUN curl -L https://github.com/TryItOnline/brainfuck/archive/$BRAINFUCK_REV.tar.gz | \
    tar -xz && \
    mv brainfuck-$BRAINFUCK_REV /opt/brainfuck
