TAGS ?= ""
GO_BIN ?= go

install:
	$(GO_BIN) install -tags ${TAGS} -v ./...

build: 
	$(GO_BIN) build -v -o shoulders.bin .

test: 
	$(GO_BIN) test -race -cover -tags ${TAGS} ./...

lint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	golangci-lint run --enable-all

update:
	-rm go.*
	$(GO_BIN) mod init github.com/gobuffalo/shoulders
	$(GO_BIN) mod tidy
	make test
	make install
	shoulders -w
