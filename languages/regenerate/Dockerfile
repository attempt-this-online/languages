#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG REGENERATE_REV=ed650aa2956998e1ed623f8c10e87dcfc78b703c

RUN curl -L https://github.com/dloscutoff/Esolangs/raw/$REGENERATE_REV/Regenerate/regenerate.py -o /opt/regenerate.py
