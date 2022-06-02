#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG TEXLIVE_VERSION=2022
ARG TEXLIVE_REPOSITORY=http://ftp.cvut.cz/tex-archive/systems/texlive/tlnet/

ENV PATH="$PATH:/usr/local/texlive/bin/linux"

COPY texlive.profile /etc/texlive.profile

RUN mkdir -p /usr/local/texlive/bin && \
    ln -s "/usr/local/texlive/bin/x86_64-linux" /usr/local/texlive/bin/linux && \
    curl -L "http://ftp.math.utah.edu/pub/tex/historic/systems/texlive/$TEXLIVE_VERSION/install-tl-unx.tar.gz" | \
    tar -xz && \
    cd install-tl-* && \
    ./install-tl --repository="$TEXLIVE_REPOSITORY" --profile=/etc/texlive.profile --scheme=minimal && \
    tlmgr install tex4ht && \
    cd / && \
    rm -rf /tmp/install-tl-*
