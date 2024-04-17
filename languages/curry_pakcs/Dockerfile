FROM registry.gitlab.pxeger.com/attempt-this-online/languages/swi_prolog

ARG PAKCS_VERSION=3.6.0

ENV PATH=/opt/cpm/bin:/opt/pakcs/bin:$PATH

RUN pacman -Syu --noconfirm git && \
    curl -L https://www.informatik.uni-kiel.de/~pakcs/download/pakcs-$PAKCS_VERSION-amd64-Linux.tar.gz | \
    tar -xz && \
    mv pakcs-* /opt/pakcs && \
    cd /opt/pakcs && \
    make -j $(nproc) && \
    mkdir /opt/cpm && \
    ln -s /opt/cpm /root/.cpm && \
    cypm update && \
    cypm install runcurry
