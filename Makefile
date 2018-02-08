.PHONY: all build install install.deps test test.setup cloc release todo size clean

# build
OUTPUT = test/out/*.go
# testing
F  = elegont_test.go
T  = TestTmp

all: build

build:
	@go build -o $(OUTPUT)

install:
	@go install

test:
	@go test $(F) -run $(T) -v

clean:
	@rm -rf $(OUTPUT)

rebuild: clean all;