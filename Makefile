.PHONY: all build test %

all: clean build

%:
	go build -o ./bin/$@ ./cmd/blg/.