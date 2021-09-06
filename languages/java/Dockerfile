FROM attemptthisonline/base

# check not already installed
RUN ! which java

ARG JAVA_URL=https://github.com/adoptium/temurin16-binaries/releases/download/jdk-16.0.2%2B7/OpenJDK16U-jdk_x64_linux_hotspot_16.0.2_7.tar.gz

RUN curl -Lo java.tar.gz "$JAVA_URL" && \
    tar -xf java.tar.gz && \
    mv -T jdk-* /opt/jdk

ENV JAVA_HOME /opt/jdk
ENV PATH $JAVA_HOME/bin:$PATH
