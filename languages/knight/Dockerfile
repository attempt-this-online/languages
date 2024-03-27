FROM attemptthisonline/node

ARG KNIGHT_VERSION=3105797369241d3f36613b4da2778c24f93f94bd

RUN cd /opt && \
    curl -L https://github.com/knight-lang/knight.js/archive/$KNIGHT_VERSION.tar.gz | \
    tar -xz && \
    mv knight.js-$KNIGHT_VERSION knight
