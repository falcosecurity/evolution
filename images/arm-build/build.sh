#!/usr/bin/env sh

set -xeu

#Clone-dir drops the pwd inside the cloned directory so drop one back
cd ..

cmake -DUSE_BUNDLED_DEPS=On -DMUSL_OPTIMIZED_BUILD=On -DBUILD_DRIVER=Off /falco

while test $# -gt 0; do
    case "$1" in
        falco)
            make -j4 falco
            exit 0
            ;;
        grpc)
            make -j4 grpc
            exit 0
            ;;
        all)
            make -j4 all
            exit 0
            ;;
        tests)
            make tests
            exit 0
            ;;
        packages)
            make -j4 package
            exit 0
            ;;
    esac
done 