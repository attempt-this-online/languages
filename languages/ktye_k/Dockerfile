FROM registry.gitlab.pxeger.com/attempt-this-online/languages/gcc

RUN curl -L https://github.com/ktye/i/releases/download/latest/k.c -o k.c && \
    gcc k.c -lm -o /usr/local/bin/k
