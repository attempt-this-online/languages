#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/clang

ARG CHAPEL_VERSION=1.28.0

ENV CHPL_HOME=/opt/chapel CHPL_LLVM=system
RUN pacman -Syu --noconfirm python cmake && \
    curl -L https://github.com/chapel-lang/chapel/releases/download/$CHAPEL_VERSION/chapel-$CHAPEL_VERSION.tar.gz | \
    tar -xz && \
    mv chapel-* /opt/chapel && \
    cd /opt/chapel && \
    make -j $(nproc) && \
    ln -s $CHPL_HOME/bin/chapel /usr/local/bin/
