# TODO: run golint, errcheck, staticcheck, unused, ineffassign
.PHONY: default
default: deps checkgofmt vet predeclared staticcheck unused golint errchack ineffassign test

.PHONY: ci
ci: deps checkgofmt vet predeclared staticcheck unused golint errchack ineffassign testcover

.PHONY: deps
deps:
	go get -d -v -t ./...

.PHONY: updatedeps
updatedeps:
	go get -d -v -t -u -f ./...

.PHONY: install
install:
	go install ./...

.PHONY: checkgofmt
checkgofmt:
	@if [ -n "$$(go version | awk '{ print $$3 }' | grep -v devel)" ]; then \
		gofmt -s -l . ; \
		if [ -n "$$(gofmt -s -l .)"  ]; then \
			echo "Run gofmt on the above files!"; \
			exit 1; \
		fi; \
	fi

# workaround https://github.com/golang/protobuf/issues/214 until in master
.PHONY: vet
vet:
	go vet ./...

.PHONY: staticcheck
staticcheck:
	@go get honnef.co/go/tools/cmd/staticcheck
	staticcheck ./...

.PHONY: unused
unused:
	@go get honnef.co/go/tools/cmd/unused
	unused ./...

.PHONY: ineffassign
ineffassign:
	@go get github.com/gordonklaus/ineffassign
	ineffassign .

.PHONY: predeclared
predeclared:
	@go get github.com/nishanths/predeclared
	predeclared .

# Intentionally omitted from CI, but target here for ad-hoc reports.
.PHONY: golint
golint:
	@go get github.com/golang/lint/golint
	golint -min_confidence 0.9 -set_exit_status ./...

# Intentionally omitted from CI, but target here for ad-hoc reports.
.PHONY: errchack
errcheck:
	@go get github.com/kisielk/errcheck
	errcheck ./...

.PHONY: test
test:
	go test -cover -race ./...

.PHONY: generate
generate:
	go generate

.PHONY: testcover
testcover:
	@echo go test -covermode=atomic ./... 
	@echo "mode: atomic" > coverage.out
	@for dir in $$(go list ./...); do \
		go test -race -coverprofile profile.out -covermode=atomic $$dir ; \
		if [ -f profile.out ]; then \
			tail -n +2 profile.out >> coverage.out && rm profile.out ; \
		fi \
	done
