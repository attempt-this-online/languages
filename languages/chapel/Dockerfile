#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG CHAPEL_VERSION=1.28.0

# install requirements to build Chapel from source
RUN pacman -Syy --noconfirm wget python cmake llvm clang

# set Chapel env vars
ENV CHPL_HOME=/opt/chapel/$CHAPEL_VERSION
ENV CHPL_LLVM=system

# get sources
WORKDIR $CHPL_HOME
RUN wget -q -O - https://github.com/chapel-lang/chapel/releases/download/$CHAPEL_VERSION/chapel-$CHAPEL_VERSION.tar.gz | tar -xzC /opt/chapel --transform 's/chapel-//'

# build
RUN make -j `nproc`

# hack to get access to Chapel binaries
RUN cd $CHPL_HOME/bin && ln -s */* .
ENV PATH="${PATH}:${CHPL_HOME}/bin:${CHPL_HOME}/util"
