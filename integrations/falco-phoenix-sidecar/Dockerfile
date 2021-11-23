FROM python:3.8.10-alpine as builder
#RUN echo http://dl-8.alpinelinux.org/alpine/v3.7/main > /etc/apk/repositories; \
#    echo http://dl-8.alpinelinux.org/alpine/v3.7/community >> /etc/apk/repositories
RUN apk add --update --no-cache \
    build-base \
    g++ \
    gcc \
    git \
    libc-dev \
    ca-certificates \
    py3-lxml \
    py3-pip \
    python3-dev \
    openssl-dev \
    libevent-dev \
    libffi-dev \
    libxslt-dev \
    libxml2-dev \
    make \
    cmake \
    dbus-glib-dev 

WORKDIR /install

COPY ./requirements.txt /
RUN pip install wheel
RUN pip install --prefix="/install" --no-warn-script-location -r /requirements.txt


###
# Main stage
###
FROM python:3.8.10-alpine

RUN apk add --no-cache \
    bash \
    sshpass \
    ca-certificates \
    openssl \
    libevent \
    libffi \
    libxslt \
    libxml2 \
    libstdc++

# copy the python dependencies
COPY --from=builder /install /usr/local

WORKDIR /app
COPY ./sidecar.py .

CMD [ "python3", "sidecar.py"]
