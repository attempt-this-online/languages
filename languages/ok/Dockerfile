FROM attemptthisonline/node

ARG OK_REV=97956bf6f2ea795178cd19bd7ae80ebe3a4111de

RUN curl -L https://github.com/JohnEarnest/ok/archive/$OK_REV.tar.gz | \
    tar -xz && \
    mv ok-$OK_REV /opt/ok
