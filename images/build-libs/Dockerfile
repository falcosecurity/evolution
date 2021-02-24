FROM debian:stable

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
    libtool \
    libelf-dev \
    wget \
    linux-headers-${ARCH} \
    && apt-get clean

COPY build.sh /
RUN chmod +x /build.sh

WORKDIR /workspace

ENTRYPOINT ["/build.sh"]