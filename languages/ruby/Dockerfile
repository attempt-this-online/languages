FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG RUBY_VERSION=3.4.4

RUN pacman -Syu libyaml --noconfirm && \
    curl https://cache.ruby-lang.org/pub/ruby/$(echo $RUBY_VERSION | cut -d. -f-2)/ruby-$RUBY_VERSION.tar.gz | \
    tar -xz && \
    cd ruby-$RUBY_VERSION && \
    ./configure && \
    make && \
    make install && \
    cd / && \
    rm -rf /tmp/ruby-$RUBY_VERSION
