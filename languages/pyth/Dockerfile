#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG PYTH_REV=b70e591254a79b2b6efaa9b05d1d56d1bd9c2512

RUN curl -L https://github.com/isaacg1/pyth/archive/$PYTH_REV.tar.gz | \
    tar -xz && \
    mv pyth-$PYTH_REV /opt/pyth
