#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/csharp AS build

ARG RETINA_VERSION=v1.2.1

RUN curl -L https://github.com/attempt-this-online/retina/archive/refs/tags/$RETINA_VERSION.tar.gz | \
    tar -xz && \
    cd retina-*/Retina && \
    mkdir /opt/retina && \
    dotnet publish --sc -r linux-x64 -o /opt/retina

FROM attemptthisonline/base

ENV DOTNET_SYSTEM_GLOBALIZATION_INVARIANT=1

COPY --from=build /opt/retina /opt/retina
