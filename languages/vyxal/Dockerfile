FROM attemptthisonline/python

ARG VYXAL_REV=v2.5.3

RUN curl -L https://github.com/Vyxal/Vyxal/archive/$VYXAL_REV.tar.gz | \
    tar -xz && \
    cd Vyxal-*/ && \
    pip install -r requirements.txt && \
    mv vyxal "$(python -c 'import site; print(site.getsitepackages()[0])')" && \
    printf '#!/bin/sh\npython -m vyxal.interpreter "$@"' > /usr/local/bin/vyxal && \
    chmod a+x /usr/local/bin/vyxal

# TODO(pxeger,Lyxal): package Vyxal properly so it can be installed with pip
