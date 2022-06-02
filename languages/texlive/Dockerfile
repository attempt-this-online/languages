#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG TEXLIVE_VERSION=2022

RUN mkdir -p /usr/local/texlive/bin && ln -s "/usr/local/texlive/bin/x86_64-linux" /usr/local/texlive/bin/linux

ENV PATH="$PATH:/usr/local/texlive/bin/linux"

COPY texlive.profile /etc/texlive.profile

RUN curl -L "http://ftp.math.utah.edu/pub/tex/historic/systems/texlive/${TEXLIVE_VERSION}/install-tl-unx.tar.gz" > /tmp/install-tl-unx.tar.gz \
    && tar -xvzf /tmp/install-tl-unx.tar.gz -C /tmp \
    && if [ "$TEXLIVE_VERSION" = 2022 ]; then repository="http://ftp.cvut.cz/tex-archive/systems/texlive/tlnet/"; else repository="http://ftp.math.utah.edu/pub/texlive/historic/systems/texlive/$TEXLIVE_VERSION/tlnet-final"; fi \
    && /tmp/*/install-tl --repository="$repository" --profile=/etc/texlive.profile --scheme=minimal && rm -rf /tmp/*

RUN tlmgr install tex4ht && rm -rf /tmp/*
