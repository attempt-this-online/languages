FROM registry.gitlab.pxeger.com/attempt-this-online/languages/gcc

ARG BRACMAT_REV=f151843735bf5cb60ef056c0023110329845d8f1

RUN curl -L https://github.com/BartJongejan/Bracmat/archive/$BRACMAT_REV.tar.gz | \
    tar -xz && \
    cd Bracmat-*/src && \
    gcc -O2 bracmat.c xml.c json.c unicaseconv.c unichartypes.c -o /usr/local/bin/bracmat && \
    cd / && \
    rm -rf /tmp/Bracmat-*
