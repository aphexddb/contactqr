RELEASE_DIST=../release/ui

all: build

.PHONY: dev
dev:
	gatsby develop

.PHONY: build
build:
	gatsby build

.PHONY: release
release: build
	rm -rf $(RELEASE_DIST)
	mkdir -p $(RELEASE_DIST)
	cp -R ./public/* $(RELEASE_DIST)

