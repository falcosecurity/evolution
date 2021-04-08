#!/bin/bash

rm _output/bucket
mkdir -p _output/bucket
aws s3 cp  --recursive "s3://logging-falco-distribution/falco-distribution/" _output/bucket


