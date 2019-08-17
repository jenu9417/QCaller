all: clean build tar

fatty: clean build deps tar

clean: ## clean up target folder
	rm -rf ./target/*

build: ## Get required lib and build

	## Create folder structure
	mkdir -p $(CURDIR)/target/QCaller/bin
	
	## Compile and build QCaller go binary
	go clean
	go build
	
	## Copy QCaller binary and config to target
	mv $(CURDIR)/QCaller $(CURDIR)/target/QCaller/bin
	cp $(CURDIR)/config.json $(CURDIR)/target/QCaller/bin

	cp -R $(CURDIR)/external/postman $(CURDIR)/target/QCaller
	cp -R $(CURDIR)/external/scripts $(CURDIR)/target/QCaller
	
	## Copy README
	cp $(CURDIR)/README.md $(CURDIR)/target/QCaller

deps:
	## Copy external files to target
	cp -R $(CURDIR)/external/es $(CURDIR)/target/QCaller

tar:
	cd $(CURDIR)/target; tar -czvf ./QCaller.tar.gz *
	## Done. Package available under /target
