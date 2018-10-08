VERSION ?= $(shell cat ./VERSION)
RELEASE_OS=linux
BINARY := contactqr
ENTRYPOINT := cmd/contactqr/contactqr.go
BINARY_NAME=contactqr
PORT := 8080
GOCMD=go
GOBIN_DIR := $(GOPATH)/bin
UI_DIR := ./ui/public
RELEASE_DIR := ./release

all: test build

.PHONY: build
build: test
	mkdir -p build
	$(GOCMD) build -o build/$(BINARY_NAME)

.PHONY: clean
clean:
	rm -rf ./build
	rm -rf ./release
	cd ui && make clean

.PHONY: release_ui
release_ui:
	cd ui && make release

.PHONY: test
test:
	$(GOCMD) test ./...

.PHONY: release
release: clean release_ui
	mkdir -p $(RELEASE_DIR)
	GOOS=$(RELEASE_OS) GOARCH=amd64 go build -o $(RELEASE_DIR)/$(BINARY)-$(VERSION)-$(RELEASE_OS)-amd64 $(ENTRYPOINT)

.PHONY: dev
dev:
	mkdir -p build
	go build -o build/$(BINARY)-dev $(ENTRYPOINT)
	@echo "Expecting UI file path: $(UI_DIR), run 'make build' in the ui directory to generate static files."
	LOCALDEV=1 build/$(BINARY)-dev -path $(UI_DIR)

.PHONY: docker_build
docker_build:
	docker build --build-arg VERSION=$(VERSION) --build-arg PORT=$(PORT) -t $(BINARY):$(VERSION) .
	docker tag $(BINARY):$(VERSION) $(BINARY):latest

.PHONY: docker_release
docker_release:
	docker tag $(BINARY):$(VERSION) docker.io/aphexddb/contactqr:$(VERSION)
	docker push docker.io/aphexddb/contactqr:$(VERSION)
	docker tag $(BINARY):$(VERSION) docker.io/aphexddb/contactqr:latest
	docker push docker.io/aphexddb/contactqr:latest

.PHONY: heroku_release
heroku_release:
	heroku container:push web --arg VERSION=$(VERSION),PORT=$(PORT)
	heroku container:release web

.PHONY: run
run:
	@echo "Running container $(BINARY):latest"
	docker run --rm -it \
		-p $(PORT):$(PORT)/tcp \
		-e PORT=$(PORT) \
		$(BINARY):latest $*