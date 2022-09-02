#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG RUBY_VERSION=3.1.2

RUN curl https://cache.ruby-lang.org/pub/ruby/$(echo $RUBY_VERSION | cut -d. -f-2)/ruby-$RUBY_VERSION.tar.gz | \
    tar -xz && \
    cd ruby-$RUBY_VERSION && \
    ./configure && \
    make && \
    make install && \
    cd / && \
    rm -rf /tmp/ruby-$RUBY_VERSION
