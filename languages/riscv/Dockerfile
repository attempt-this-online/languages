FROM attemptthisonline/base

RUN ! which jupiter
ARG JUPITER_URL=https://github.com/andrescv/Jupiter/releases/download/v3.1/Jupiter-3.1-linux.zip
RUN curl -L "$JUPITER_URL" | \
    unzip && \
    mv image /opt/jupiter
