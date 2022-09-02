#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/gcc

RUN curl -L https://github.com/ktye/i/releases/download/latest/k.c -o k.c && \
    gcc k.c -lm -o /usr/local/bin/k
