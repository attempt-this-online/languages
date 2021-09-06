FROM attemptthisonline/java

# setup install dir
RUN mkdir -p /opt/coursier_bin && \
    chmod a=rX,u+w /opt /opt/coursier_bin && \
    mkdir -p /root/.local/share/coursier/ && \
    ln -s /opt/coursier_bin /root/.local/share/coursier/bin # trick coursier into installing binaries into /opt/

ENV PATH="/opt/coursier_bin:$PATH"

# install coursier
RUN curl -fLo cs https://git.io/coursier-cli-linux && \
    chmod +x cs && \
    ./cs setup -y --env --apps scala3-compiler && \
    scala3-compiler --version
#     ^ forces Scala compiler to be downloaded
