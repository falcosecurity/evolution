FALCO_IMAGE_NAME ?= "falcosecurity/falco-pdig"

all: binaries/logshipper binaries/launcher binaries/waitforever
	docker build -t "$(FALCO_IMAGE_NAME)" --no-cache .