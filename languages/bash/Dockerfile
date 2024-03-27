FROM attemptthisonline/base

# THE_ prefix because $BASH_VERSION is a special variable
ARG THE_BASH_VERSION=5.2.21

RUN curl -L https://ftp.gnu.org/gnu/bash/bash-$THE_BASH_VERSION.tar.gz | \
    tar -xz && \
    cd bash-$THE_BASH_VERSION && \
    ./configure && \
    make -j $(nproc) && \
    make install && \
    cd / && \
    rm -rf /tmp/bash-$THE_BASH_VERSION
