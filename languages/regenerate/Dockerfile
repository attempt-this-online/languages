#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG REGENERATE_REV=b130e9ae0d32d0f601ea126d9096013064e3186b

RUN curl -L https://github.com/dloscutoff/Esolangs/raw/$REGENERATE_REV/Regenerate/regenerate.py -o /opt/regenerate.py
