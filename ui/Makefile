PUBLIC=./public
RELEASE_DIST=../release/ui

all: build

.PHONY: clean
clean:
	rm -rf $(PUBLIC)
	rm -rf $(RELEASE_DIST)

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
	cp -R $(PUBLIC)/* $(RELEASE_DIST)

