FROM attemptthisonline/clang AS build

ARG BQN_REV=5e865589b6e644fb2a275f26486ccf0ebc73d74b

RUN pacman -Syu --noconfirm git && \
    git clone https://github.com/dzaima/CBQN.git && \
    cd CBQN && \
    git checkout $BQN_REV && \
    make

FROM attemptthisonline/base

COPY --from=build /tmp/CBQN/BQN /usr/local/bin/
