FROM attemptthisonline/java

ARG SCALA_VERSION=3.1.3

RUN curl -L https://github.com/lampepfl/dotty/releases/download/$SCALA_VERSION/scala3-$SCALA_VERSION.tar.gz | \
    tar -xz && \
    mv -T scala3-$SCALA_VERSION /usr/local/share/scala && \
    chown -R root:root /usr/local/share/scala && \
    chmod -R a=rX,o+w /usr/local/share/scala && \
    ln -s /usr/local/share/scala/bin/* /usr/local/bin/
