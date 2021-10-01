GO ?= go

all:	fodselsnr


fodselsnr:	$(wildcard *.go) $(wildcard */*.go) $(wildcard cmd/appapi-push/*.go)
		$(GO) mod tidy
		$(GO) build -v -o fodselsnr cmd/fodselsnr/main.go

clean:
	rm -rf fodselsnr

vet:
	$(GO) vet ./...

lint:
	golint ./...

fmt:
	gofmt -w .

gosec:
	gosec ./...

.PHONY: clean vet lint fmt gosec all
