FROM attemptthisonline/base

RUN ! which jupiter
ARG JUPITER_URL=https://github.com/adoptium/temurin18-binaries/releases/download/jdk-18.0.1%2B10/OpenJDK18U-jdk_x64_linux_hotspot_18.0.1_10.tar.gz

RUN curl -L "$JUPITER_URL" | \
    unzip && \
    mv image /opt/jupiter
