#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/csharp AS build

ARG FUNKY2_VERSION=v1.1

RUN curl -L https://github.com/TehFlaminTaco/Funky2/archive/refs/tags/$FUNKY2_VERSION.tar.gz | \
    tar -xz && \
    cd Funky2-* && \
    dotnet publish --sc -r linux-x64 -o /opt/

FROM attemptthisonline/base

COPY --from=build /opt/Funky2 /usr/local/bin/
