#!/bin/bash
cat <&0 > chart.yaml
kustomize edit add resource chart.yaml
kustomize build . && rm chart.yaml
