FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

# Nightly:
# ARG FACTOR_VERSION=2023-07-30-13-46
#
# RUN curl -L https://downloads.factorcode.org/linux-x86-64/factor-linux-x86-64-$FACTOR_VERSION.tar.gz | \
#     tar -xz && \
#     mv factor /opt/

# Release:
ARG FACTOR_VERSION=0.100

RUN curl -L https://downloads.factorcode.org/releases/$FACTOR_VERSION/factor-linux-x86-64-$FACTOR_VERSION.tar.gz | \
    tar -xz && \
    mv factor /opt/
