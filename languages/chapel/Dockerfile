FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG CHAPEL_VERSION=2.7.0


ENV CHPL_HOME=/opt/chapel CHPL_LLVM=none
RUN pacman -Syu --noconfirm python cmake && \
    curl -L https://github.com/chapel-lang/chapel/releases/download/$CHAPEL_VERSION/chapel-$CHAPEL_VERSION.tar.gz | \
    tar -xz && \
    mv chapel-* /opt/chapel && \
    cd /opt/chapel && \
    make -j $(nproc) && \
    ln -s $CHPL_HOME/bin/*/chpl /usr/local/bin/
