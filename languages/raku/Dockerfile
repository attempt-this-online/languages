FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG RAKU_VERSION=2023.12-01

ENV PATH=/opt/raku/bin:/opt/raku/share/perl6/site/bin:$PATH

RUN curl https://rakudo.org/dl/rakudo/rakudo-moar-$RAKU_VERSION-linux-x86_64-gcc.tar.gz | \
    tar -xz && \
    mv rakudo-* /opt/raku
