#syntax=docker/dockerfile-upstream:1.4.0-rc1
FROM attemptthisonline/python

ARG CHARCOAL_REV=bd84513685ac54187a849c01dd2d5ad52d15d3b8

RUN pip install --no-cache-dir --upgrade regex brotli && \
    curl -L https://github.com/somebody1234/Charcoal/archive/$CHARCOAL_REV.zip -o /opt/charcoal.zip && \
    cd /opt && \
    unzip charcoal.zip && \
    mv Charcoal-$CHARCOAL_REV charcoal && \
    rm charcoal.zip
