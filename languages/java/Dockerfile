FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

# check not already installed
RUN ! which java

# go to https://adoptium.net/ to get the latest URL

ARG JAVA_URL=https://github.com/adoptium/temurin18-binaries/releases/download/jdk-18.0.1%2B10/OpenJDK18U-jdk_x64_linux_hotspot_18.0.1_10.tar.gz

RUN curl -L "$JAVA_URL" | \
    tar -xz && \
    mv -T jdk-* /opt/jdk

ENV JAVA_HOME /opt/jdk
ENV PATH $JAVA_HOME/bin:$PATH
