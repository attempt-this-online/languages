#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/node

# modified from https://github.com/elm/compiler/blob/master/installers/linux/README.md

ARG ELM_VERSION=0.19.1
ARG ELM_URL=https://github.com/elm/compiler/releases/download/${ELM_VERSION}/binary-for-linux-64-bit.gz

RUN curl -L -o elm.gz ${ELM_URL} && \
    gunzip elm.gz && \
    chmod +x elm && \
    mv elm /opt/elm && \
    npm install --global elm-line
