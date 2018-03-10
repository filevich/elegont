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

# go test -v config.go syntax.go utils.go syntax_YAML.go elegont.go syntax_utils.go elegont_test.go