FROM debian:buster

LABEL usage="docker run -i -t -v /path/to/source:/workspace test-infra/build-libs [cmake options]"

ARG ARCH=amd64

RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates \
    cmake \
    build-essential \
    clang \
    llvm \
    git \
    libncurses-dev \
    pkg-config \
    autoconf \ 
    automake \
    libtool \
    libelf-dev \
    wget \
    libb64-dev \
    libc-ares-dev \
    libcurl4-openssl-dev \
    libssl-dev \
    libtbb-dev \
    libjq-dev \
    libjsoncpp-dev \
    libgrpc++-dev \
    protobuf-compiler-grpc \
    libgtest-dev \
    linux-headers-${ARCH} \
    && apt-get clean

COPY build.sh /
RUN chmod +x /build.sh

WORKDIR /workspace

ENTRYPOINT ["/build.sh"]
