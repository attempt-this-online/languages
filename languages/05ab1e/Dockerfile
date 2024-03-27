FROM attemptthisonline/elixir

ARG OSABIE_REV=1d532f3aad3af265691c9c1bc601c246bfa2944e

RUN curl -L https://github.com/Adriandmen/05AB1E/archive/$OSABIE_REV.tar.gz | \
    tar -xz && \
    cd 05AB1E-$OSABIE_REV && \
    (echo Y | mix local.hex) && \
    mix deps.update --all && \
    (echo Y | MIX_ENV=prod mix escript.build) && \
    mv osabie /usr/local/bin && \
    rm -rf 05AB1E-$OSABIE_REV
