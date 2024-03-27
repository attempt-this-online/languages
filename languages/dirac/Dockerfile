FROM attemptthisonline/gcc

ARG DIRAC_REV=c1a2b9c42f2dba5c685da0769093e08fc27878f2

RUN curl -L https://github.com/kspalaiologos/dirac/archive/$DIRAC_REV.tar.gz | \
    tar -xz && \
    cd dirac-$DIRAC_REV && \
    gcc -O2 dirac.c -o /usr/local/bin/dirac && \
    cd / && \
    rm -rf /tmp/dirac-$DIRAC_REV
