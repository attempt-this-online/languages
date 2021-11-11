FROM attemptthisonline/base

ARG DOTNET_SDK_VERSION=6.0.100

RUN mkdir /opt/dotnet && \
    curl -SL https://dotnetcli.azureedge.net/dotnet/Sdk/$DOTNET_SDK_VERSION/dotnet-sdk-$DOTNET_SDK_VERSION-linux-x64.tar.gz | \
    tar -xzC /opt/dotnet && \
    ln -s /opt/dotnet/dotnet /usr/local/bin/dotnet
