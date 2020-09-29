all: definitions/falco-pdig

definitions/%: definitions-utils
	$(MAKE) -C $@ all

utils/%:
	$(MAKE) -C $(@D) $%

definitions-utils: utils/launcher/launcher utils/waitforever/waitforever utils/logshipper/logshipper
	mkdir -p definitions/falco-pdig/binaries
	cp $^ definitions/falco-pdig/binaries

