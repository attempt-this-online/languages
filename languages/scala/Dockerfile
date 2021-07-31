FROM openjdk:11

WORKDIR /tmp

RUN apt-get update && \
    apt-get upgrade -y && \
    curl -fsL https://downloads.typesafe.com/scala/2.13.6/scala-2.13.6.tgz | tar xfz - -C /usr/share && \
    mv /usr/share/scala-2.13.6 /usr/share/scala && \
    chown -R root:root /usr/share/scala && \
    chmod -R 755 /usr/share/scala && \
    ln -s /usr/share/scala/bin/* /usr/local/bin

ENTRYPOINT /usr/local/bin/scala
