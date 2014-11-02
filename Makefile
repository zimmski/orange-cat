NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

all: checkfmt deps build

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

checkfmt:
	@echo "$(OK_COLOR)==> Checking if there's any file to fmt$(NO_COLOR)"
	@scripts/gofmt.sh
	@echo "$(OK_COLOR) => Done$(NO_COLOR)"
