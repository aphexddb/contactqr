VERSION ?= $(shell cat ./VERSION)
RELEASE_OS=linux
BINARY := contactqr
ENTRYPOINT := cmd/contactqr/contactqr.go
BINARY_NAME=contactqr
GOCMD=go
GOBINDIR := $(GOPATH)/bin
UIPATH := ./ui/public

all: test build

.PHONY: ui
ui:
	cd ui && make release

.PHONY: build
build: test
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME)

.PHONY: test
test:
	$(GOCMD) test ./...

.PHONY: release
release: ui
	mkdir -p release
	GOOS=$(RELEASE_OS) GOARCH=amd64 go build -o release/$(BINARY)-$(VERSION)-$(RELEASE_OS)-amd64 $(ENTRYPOINT)

.PHONY: dev
dev:
	mkdir -p build
	go build -o build/$(BINARY)-dev $(ENTRYPOINT)
	@echo "Expecting UI file path: $(UIPATH), run 'make build' in the ui directory to generate static files."
	build/$(BINARY)-dev -path $(UIPATH)

.PHONY: docker_build
docker_build: release
	docker build --build-arg VERSION=$(VERSION) -t $(BINARY):$(VERSION) .
	docker tag $(BINARY):$(VERSION) $(BINARY):latest

.PHONY: docker_release
docker_release:
	docker push $(BINARY):$(VERSION) CHANGEME_REPO.here

.PHONY: run
run:
	docker run --rm -it \
		-p 8080:8080/tcp \
		-e PORT=8080 \
		$(BINARY):latest $*