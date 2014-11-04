NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)
TEST_DIR = 'tests'

all: fmt testdeps test deps build

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
	@go build -o orange
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"

fmt:
	@echo "$(OK_COLOR)==> Fmt'ing source codes$(NO_COLOR)"
	@if scripts/gofmt.sh; \
		then echo "$(OK_COLOR) => Done$(NO_COLOR)"; \
		else echo "$(WARN_COLOR) => Fmt'ed$(NO_COLOR)"; exit 1; \
	fi

testdeps:
	@go get github.com/onsi/ginkgo
	@go get github.com/onsi/gomega

test: testdeps
	@echo "$(OK_COLOR)==> Testing modules$(NO_COLOR)"
	@cd $(TEST_DIR) && \
		go test
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"
