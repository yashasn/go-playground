PACKAGES := $(shell go list ./...)
VERSION:=1.0.0
LDFLAGS := -ldflags "-X main.Version=${VERSION}"
MODULE = $(shell go list -m)


.PHONY: test
test: ##run unit tests
	@go test $(PACKAGES)

.PHONY: test-coverage
test-coverage: ## run unit tests with coverage
	@echo "mode: count" > coverage-all.out
	@$(foreach pkg,$(PACKAGES), \
		go test -p=1 -cover -covermode=count -coverprofile=coverage.out ${pkg}; \
		tail -n +2 coverage.out >> coverage-all.out;)

.PHONY: build
build:  ## build the API server binary
	CGO_ENABLED=0 go build ${LDFLAGS} -a -o gameserver $(MODULE)

.PHONY: run
run: ## run the API server
	@go run ${LDFLAGS} main.go

.PHONY: clean
clean: ## remove temporary files
	rm -rf server coverage.out coverage-all.out 

.PHONY: lint
lint: ## run golint on all Go package
	@golangci-lint run ./...

.PHONY: fmt
fmt: ## run "go fmt" on all Go packages
	@go fmt $(PACKAGES)