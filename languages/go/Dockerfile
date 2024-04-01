FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

RUN pacman -Syu --noconfirm go

ENV GOPATH=/opt/go

RUN mkdir /opt/go && \
    cd /opt && \
    go mod init ato.pxeger.com/go && \
    go get golang.org/x/exp
