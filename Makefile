.PHONY: all build test %

GOBUILD=go build
GORUN= go run
GOFLAGS=
GOLDFLAGS=
OUT=./bin/
BIN=blg
BINOUT=$(OUT)$(BIN)
BUILDDIR=./cmd/
BUILDFILE=$(BUILDDIR)$(BIN)/$(BIN).go
SOFTRM=-rm -rf


all: clean build

clean:
	@$(SOFTRM) $(OUT)* 


build:
	@$(GOBUILD) -o $(BINOUT) $(BUILDFILE)