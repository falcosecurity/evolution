IMG_NAME = test-infra/arm-build
ACCOUNT=292999226676
DOCKER_PUSH_REPOSITORY=dkr.ecr.eu-west-1.amazonaws.com
TAG?=latest
IMAGE=$(ACCOUNT).$(DOCKER_PUSH_REPOSITORY)/$(IMG_NAME)

build-push: build-image push-image

build-image:
	docker build -t $(IMG_NAME) .

push-image:
	docker tag $(IMG_NAME) $(IMAGE):$(TAG)
	docker push $(IMAGE):$(TAG)