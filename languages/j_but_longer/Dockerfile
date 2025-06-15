FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG J_VERSION=9.6

COPY setup /tmp/

RUN pacman -U --noconfirm https://archive.archlinux.org/packages/g/glibc/glibc-2.40-1-x86_64.pkg.tar.zst && \
    curl -L http://www.jsoftware.com/download/j$J_VERSION/install/j${J_VERSION}_linux64.tar.gz | \
    tar -xz && \
    mv j$J_VERSION /opt/j && \
    cd /opt/j && \
    ./jconsole.sh /tmp/setup && \
    rm -rf /tmp/*
