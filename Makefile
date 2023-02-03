GO ?= go
GOLINT ?= golint
GOFMT ?= gofmt
GOSEC ?= gosec
GOVULNCHECK ?= govulncheck

all:	fodselsnr


fodselsnr:	$(wildcard *.go) $(wildcard */*.go) $(wildcard cmd/fodselsnr/*.go)
		$(GO) mod tidy
		$(GO) build -v -o fodselsnr cmd/fodselsnr/main.go

clean:
	rm -rf fodselsnr

vet:
	$(GO) vet ./...

lint:
	$(GOLINT) ./...

fmt:
	$(GOFMT) -w .

gosec:
	$(GOSEC) ./...

check:
	$(GOVULNCHECK) ./...

.PHONY: all clean vet lint fmt gosec check
