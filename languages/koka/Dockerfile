#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

RUN pacman -Sy --noconfirm && \
    curl -fsSL https://github.com/koka-lang/koka/releases/latest/download/install.sh | sh
