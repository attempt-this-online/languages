FROM attemptthisonline/clang

ARG COGNATE_REV=7582199a03830dceb803c294686e09e3746cf814

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
