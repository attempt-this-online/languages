FROM registry.gitlab.pxeger.com/attempt-this-online/languages/ruby

ARG LABYRINTH_REV=d3cc04f398594549c1c85c8116ef146d461a5b6f

RUN curl -L https://github.com/m-ender/labyrinth/archive/$LABYRINTH_REV.tar.gz | \
    tar -xz && \
    mv labyrinth-* /opt/labyrinth
