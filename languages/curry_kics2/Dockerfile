FROM registry.gitlab.pxeger.com/attempt-this-online/languages/curry_pakcs as build

ARG KICS2_VERSION=3.4.0

ENV PATH=/opt/kics2/bin:$PATH

RUN pacman -Syu --noconfirm python stack && \
    cd /opt && \
    curl -L https://www.curry-lang.org/kics2/download/kics2-$KICS2_VERSION-x86_64-linux.tar.gz | \
    tar -xz && \
    mv kics2-* kics2 && \
    cd kics2 && \
    make -j $(nproc)

FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

COPY --from=build /opt/kics2 /opt/kics2

ENV PATH=/opt/cpm/bin:/opt/kics2/bin:$PATH

RUN pacman -Syu --noconfirm stack git && \
    mkdir /opt/cpm && \
    ln -s /opt/cpm /root/.cpm && \
    cypm update && \
    cypm install runcurry
