FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

RUN pacman -Syu --noconfirm rustup && \
    rustup toolchain install beta --profile minimal && \
    rustup component remove cargo && \
    mv /root/.rustup /opt/rust && \
    chmod -R a+rX /opt/rust

ENV RUSTUP_HOME=/opt/rust
