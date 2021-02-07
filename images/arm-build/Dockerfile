FROM multiarch/alpine:aarch64-v3.12

RUN apk update
RUN apk add \
    bash \
    g++ \
    gcc \
    cmake \
    cmake \
    make \
    ncurses-dev \
    git \
    perl \
    linux-headers \
    autoconf \
    automake \
    m4 \
    libtool \
    elfutils-dev \
    libelf-static \
    patch \
    binutils

COPY build.sh /build.sh

RUN chmod 755 /build.sh

ENTRYPOINT ["/build.sh"]