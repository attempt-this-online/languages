#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/clang

ARG COGNATE_REV=07673f8fe2cbdc17ac0729943c51c8704ddfceeb

# vim is needed because cognate uses xxd
RUN pacman -Syu --noconfirm vim libblocksruntime && \
    curl -L https://github.com/cognate-lang/cognate/archive/$COGNATE_REV.tar.gz | \
    tar -xz && \
    cd cognate-$COGNATE_REV && \
    make && \
    make PREFIX=/usr/local install && \
    cd / && \
    rm -rf /tmp/cognate-$COGNATE_REV && \
    pacman -Rns --noconfirm vim
