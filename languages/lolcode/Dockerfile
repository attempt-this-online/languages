#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/base

ARG LOLCODE_VERSION=0.11.2

RUN curl -L https://github.com/justinmeza/lci/archive/refs/tags/v$LOLCODE_VERSION.tar.gz | \
	tar -xz && \
	cd lci-$LOLCODE_VERSION && \
	pacman -Syu --noconfirm cmake && \
	cmake . && \
	make && \
	make install && \
	cp lci /usr/local/bin/ && \ 
	cd /tmp/ && \
	rm -rf lci-$LOLCODE_VERSION
