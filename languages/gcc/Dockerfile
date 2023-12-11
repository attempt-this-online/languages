#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG BOOST_VERSION=1.80.0

# Install Boost
RUN cd /tmp && \
    BOOST_VERSION_MOD=$(echo $BOOST_VERSION | tr . _) && \
    wget https://boostorg.jfrog.io/artifactory/main/release/${BOOST_VERSION}/source/boost_${BOOST_VERSION_MOD}.tar.bz2 && \
    tar --bzip2 -xf boost_${BOOST_VERSION_MOD}.tar.bz2 && \
    cd boost_${BOOST_VERSION_MOD} && \
    ./bootstrap.sh --prefix=/usr/local && \
    ./b2 install && \
    rm -rf /tmp/*

RUN pacman -Syu --noconfirm gcc-{objc,go,fortran,ada,d}
