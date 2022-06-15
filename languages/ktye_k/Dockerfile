#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/go

RUN curl -L https://github.com/ktye/i/releases/download/latest/kg.go -o kg.go && \
    go build -o /usr/local/bin/kg kg.go
