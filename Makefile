NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
TEST_DEPS = \
	github.com/onsi/ginkgo \
	github.com/onsi/gomega
TEST_DIR = 'tests'
XC_ARCH = "darwin/amd64 darwin/386 linux/amd64 linux/386 windows/amd64 windows/386"

all: fmt test xctoolchain xc

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	@go get -d -v ./...
	@echo $(DEPS) | xargs -n1 go get -d
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

updatedeps:
	@echo "$(OK_COLOR)==> Updating all dependencies$(NO_COLOR)"
	@go get -d -v -u ./...
	@echo $(DEPS) | xargs -n1 go get -d -u
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

build: deps
	@echo "$(OK_COLOR)==> Building orange-cat$(NO_COLOR)"
	@go build -o out/orange
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

fmt:
	@echo "$(OK_COLOR)==> Fmt'ing source codes$(NO_COLOR)"
	@if scripts/gofmt.sh; \
		then echo "$(OK_COLOR) => Done$(NO_COLOR)"; \
		else echo "$(WARN_COLOR) => Fmt'ed$(NO_COLOR)"; exit 1; \
	fi

testdeps:
	@echo "$(OK_COLOR)==> Installing test dependencies$(NO_COLOR)"
	@- $(foreach DEP, $(TEST_DEPS), \
		go get -d -v $(DEP); \
	)
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

test: testdeps deps
	@echo "$(OK_COLOR)==> Testing modules$(NO_COLOR)"
	@cd $(TEST_DIR) && \
		go test -ginkgo.v
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

gox:
	@go get github.com/mitchellh/gox

xctoolchain: gox
	@gox -osarch=$(XC_ARCH) -build-toolchain

xc: deps gox
	@echo "$(OK_COLOR)==> Compiling into multiple targets$(NO_COLOR)"
	@go get github.com/mitchellh/gox
	@gox -osarch=$(XC_ARCH) -output="./out/{{.OS}}_{{.Arch}}/orange"
	@scripts/zip_output.sh
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

clean:
	@echo "$(OK_COLOR)==> Cleaning$(NO_COLOR)"
	@rm -rf ./out
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"
