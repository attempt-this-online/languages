FROM attemptthisonline/java

ARG SCALA_VERSION=2.13.8

RUN curl -L https://downloads.typesafe.com/scala/$SCALA_VERSION/scala-$SCALA_VERSION.tgz | \
    tar -xz && \
    mv -T scala-$SCALA_VERSION /usr/local/share/scala && \
    chown -R root:root /usr/local/share/scala && \
    chmod -R a=rX,o+w /usr/local/share/scala && \
    ln -s /usr/local/share/scala/bin/* /usr/local/bin/
