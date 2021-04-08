#!/bin/bash

# decompress and remove headers
zcat _output/bucket/* > _output/all.log
sed -i '/^#/ d' _output/all.log


# alternative file for visidata processing

zcat _output/bucket/*  | head -n 2 | tail -n 1 | sed 's/#Fields: //' | sed 's/ /\t/g' > _output/visidata.log
cat _output/all.log >> _output/visidata.log
