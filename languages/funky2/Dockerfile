FROM attemptthisonline/csharp

ARG FUNKY2_VERSION=v1.1

RUN curl -L https://github.com/TehFlaminTaco/Funky2/archive/refs/tags/$FUNKY2_VERSION.tar.gz | \
    tar -xz && \
    cd Funky2-* && \
    dotnet publish -r linux-x64 -o /usr/local/bin/ && \
    cd / && \
    rm -rf /tmp/Funky2-*
