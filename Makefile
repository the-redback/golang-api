REGISTRY ?= maruftuhin
IMAGE_REPO = conways
IMAGE_VERSION ?= latest
IMAGE_NAME = $(REGISTRY)/$(IMAGE_REPO):$(IMAGE_VERSION)

.PHONY: build push run kube-run

build:
	docker build --no-cache -t $(IMAGE_NAME) .

push:
	docker push $(IMAGE_NAME)

run:
	docker run -p 12345:12345 -ti --net=host --rm $(IMAGE_NAME)

default: run
