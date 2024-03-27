FROM attemptthisonline/java

ARG APL_REV=5eb0a4205e27afa6122096a25008474eec562dc0

RUN curl -L https://github.com/dzaima/APL/archive/$APL_REV.tar.gz | \
    tar -xz && \
    cd APL-$APL_REV && \
    ./build && \
    mv APL.jar /opt/ && \
    cd / && \
    rm -rf /tmp/APL-$APL_REV
