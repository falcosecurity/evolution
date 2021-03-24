SHELL := /bin/bash

IMG_SLUG := test-infra
IMG_NAME := update-falco-k8s-manifests
IMG_TAG ?= latest

ACCOUNT := 292999226676
DOCKER_PUSH_REPOSITORY = dkr.ecr.eu-west-1.amazonaws.com

IMAGE := "$(ACCOUNT).$(DOCKER_PUSH_REPOSITORY)/$(IMG_SLUG)/$(IMG_NAME):$(IMG_TAG)"

build-push: build-image push-image

build-image:
	docker build --no-cache -t "$(IMG_SLUG)/$(IMG_NAME)" .

push-image:
	docker tag "$(IMG_SLUG)/$(IMG_NAME)" $(IMAGE)
	docker push $(IMAGE)

local-registry:
	docker tag "$(IMG_SLUG)/$(IMG_NAME)" localhost:5000/$(IMG_NAME)
	docker push localhost:5000/$(IMG_NAME)