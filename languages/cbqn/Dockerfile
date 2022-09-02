#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/clang AS build

ARG BQN_REV=76141f6fe642f87364be1c61894d60c5d22ca0bf

RUN pacman -Syu --noconfirm git && \
    git clone https://github.com/dzaima/CBQN.git && \
    cd CBQN && \
    git checkout $BQN_REV && \
    make

FROM attemptthisonline/base

COPY --from=build /tmp/CBQN/BQN /usr/local/bin/
