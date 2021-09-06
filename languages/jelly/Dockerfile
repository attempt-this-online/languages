FROM attemptthisonline/python

ARG JELLY_REV=70c9fd93ab009c05dc396f8cc091f72b212fb188

RUN curl -Lo jelly.tar.gz \
        https://github.com/DennisMitchell/jellylanguage/archive/$JELLY_REV.tar.gz \
        && \
    pip install jelly.tar.gz
