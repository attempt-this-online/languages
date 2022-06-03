#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

RUN curl -L https://github.com/PyGamer0/flax/archive/refs/heads/main.zip -o flax.git && \
    mkdir /opt/flax && \
    cd /opt/flax && \
    unzip /tmp/flax.git && \
    mv flax-main/* . && \
    rm -rf flax-main && \
    pip install mpmath more-itertools
