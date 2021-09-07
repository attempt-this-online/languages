FROM attemptthisonline/base

# check not already installed
RUN ! which java

ARG JAVA_URL=https://github.com/adoptium/temurin16-binaries/releases/download/jdk-16.0.2%2B7/OpenJDK16U-jdk_x64_linux_hotspot_16.0.2_7.tar.gz

RUN curl -L "$JAVA_URL" | \
    tar -xz && \
    mv -T jdk-* /opt/jdk

ENV JAVA_HOME /opt/jdk
ENV PATH $JAVA_HOME/bin:$PATH
