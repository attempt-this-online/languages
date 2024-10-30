FROM registry.gitlab.pxeger.com/attempt-this-online/languages/base

ARG POWERSHELL_VERSION=7.4.6

RUN mkdir -p /opt/microsoft/powershell/7 && \
    curl -L https://github.com/PowerShell/PowerShell/releases/download/v$POWERSHELL_VERSION/powershell-$POWERSHELL_VERSION-linux-x64.tar.gz | \
    tar -zx -C /opt/microsoft/powershell/7 && \
    chmod +x /opt/microsoft/powershell/7/pwsh && \
    ln -s /opt/microsoft/powershell/7/pwsh /usr/local/bin/
