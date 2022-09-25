#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

RUN pacman -Syu --noconfirm go

ENV GOPATH=/opt/go

RUN mkdir /opt/go && \
    cd /opt && \
    go mod init ato.pxeger.com/go && \
    go get golang.org/x/exp
