#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/ruby

ARG BRAINFLAK_REV=9665dec1bd07aadbfd5889ff742d5bec6d19d031

RUN curl -L https://github.com/DJMcMayhem/Brain-Flak/archive/$BRAINFLAK_REV.tar.gz | \
    tar -xz && \
    mv Brain-Flak-* /opt/brain-flak
