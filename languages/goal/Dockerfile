FROM registry.gitlab.pxeger.com/attempt-this-online/languages/go

ARG GOAL_REV=v1.0.0

RUN curl -L https://codeberg.org/anaseto/goal/archive/$GOAL_REV.tar.gz | tar -xz && \
    cd goal && \
    go build ./cmd/goal && \
    mv goal /usr/local/bin/ && \
    cd .. && \
    rm -rf goal
