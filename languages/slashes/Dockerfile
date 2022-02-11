FROM attemptthisonline/perl

ARG SLASHES_REV=460a0e01ea6df20812a2c952728c65e278718e29

RUN curl -L https://github.com/TryItOnline/slashes/archive/$SLASHES_REV.tar.gz | \
    tar -xz && \
    mv slashes-$SLASHES_REV /opt/slashes
