FROM registry.gitlab.pxeger.com/attempt-this-online/languages/ruby

ARG J_UBY_REV=d95e2d0786ee639a7cf63f48a1404d657c56fbcd

RUN curl -L https://github.com/cyoce/J-uby/archive/$J_UBY_REV.tar.gz | \
    tar -xz && \
    mv J-uby-$J_UBY_REV /opt/J-uby
