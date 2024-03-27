FROM attemptthisonline/base

RUN pacman -Syu --noconfirm go

ENV GOPATH=/opt/go

RUN mkdir /opt/go && \
    cd /opt && \
    go mod init ato.pxeger.com/go && \
    go get golang.org/x/exp
