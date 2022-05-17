#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/elixir

RUN curl -L https://github.com/Adriandmen/05AB1E/archive/refs/heads/master.zip -o 05ab1e.zip && \
    mkdir /opt/05ab1e && \
    cd /opt/05ab1e && \
    unzip /tmp/05ab1e.zip && \
    cd 05AB1E-master && \
    (yes | PATH=/opt/elixir/bin:$PATH mix local.hex) && \
    PATH=/opt/elixir/bin:$PATH mix deps.get && \
    (yes | PATH=/opt/elixir/bin:$PATH MIX_ENV=prod mix escript.build) && \
    mv osabie .. && \
    rm /tmp/05ab1e.zip