#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/node

ARG ELM_VERSION=0.19.1

RUN curl -L https://github.com/elm/compiler/releases/download/$ELM_VERSION/binary-for-linux-64-bit.gz | \
    gzip -d > /usr/local/bin/elm && \
    chmod +x /usr/local/bin/elm && \
    npm install --global elm-line && \
    echo 'main = Interact.onLines identity' > dummy.elm && \
    elm-line dummy.elm
