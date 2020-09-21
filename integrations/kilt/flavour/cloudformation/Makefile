kilt.zip: build/handler
	cd build && zip ../kilt.zip ./handler

kilt-cfn:
	go build .

build/handler: kilt-cfn
	mkdir -p build
	cp kilt-cfn build/handler
	strip build/handler

clean:
	rm kilt
	rm -rf build


.PHONY: clean