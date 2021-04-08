#!/bin/bash

# decompress and remove headers
zcat _output/bucket/* > _output/all.log
sed -i '/^#/ d' _output/all.log

