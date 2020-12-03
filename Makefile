GOCMD=go
GOBUILD := $(GOCMD) build
GOFILES := $(shell find . -maxdepth 1 -name '*.go' -not -name '*_test.go')
TESTFILES := $(shell find . -maxdepth 1 -name '*_test.go')
BINNAME := bin/aoc2020

test:
	go test

bench: clean
	go test -bench=.

run: bin/aoc2020
	./bin/aoc2020

build: bin/aoc2020

$(BINNAME): $(GOFILES)
	$(GOBUILD) -o $(BINNAME) $(GOFILES)

clean:
	rm -f $(BINNAME)