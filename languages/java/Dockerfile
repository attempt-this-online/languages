#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

# check not already installed
RUN ! which java

# go to https://adoptium.net/ to get the latest URL

ARG JAVA_URL=https://github.com/adoptium/temurin17-binaries/releases/download/jdk-17.0.2%2B8/OpenJDK17U-jdk_x64_linux_hotspot_17.0.2_8.tar.gz

RUN curl -L "$JAVA_URL" | \
    tar -xz && \
    mv -T jdk-* /opt/jdk

ENV JAVA_HOME /opt/jdk
ENV PATH $JAVA_HOME/bin:$PATH
