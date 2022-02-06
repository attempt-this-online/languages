FROM attemptthisonline/base

RUN pacman -Syu --noconfirm rustup && \
    rustup toolchain install beta --profile minimal && \
    rustup component remove cargo
