#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/elixir

ARG OSABIE_REV=1d532f3aad3af265691c9c1bc601c246bfa2944e

RUN curl -L https://github.com/Adriandmen/05AB1E/archive/$OSABIE_REV.zip -o 05ab1e.zip && \
    mkdir /opt/05ab1e && \
    cd /opt/05ab1e && \
    unzip /tmp/05ab1e.zip && \
    cd 05AB1E-$OSABIE_REV && \
    (yes | PATH=/opt/elixir/bin:$PATH mix local.hex) && \
    PATH=/opt/elixir/bin:$PATH mix deps.get && \
    (yes | PATH=/opt/elixir/bin:$PATH MIX_ENV=prod mix escript.build) && \
    mv osabie .. && \
    rm /tmp/05ab1e.zip || true
