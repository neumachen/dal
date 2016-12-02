NO_COLOR=\033[0m
OK_COLOR=\033[32;01m
ERROR_COLOR=\033[31;01m
WARN_COLOR=\033[33;01m
DEPS = $(go list -f '{{range .TestImports}}{{.}} {{end}}' ./...)

deps:
	@echo "$(OK_COLOR)==> Installing dependencies$(NO_COLOR)"
	go get -u github.com/golang/lint/golint
	go get github.com/rnubel/pgmgr
	go get -t -d -v ./...

format:
	@echo "$(OK_COLOR)==> Formatting$(NO_COLOR)"
	go fmt ./...

test:
	@echo "$(OK_COLOR)==> Testing$(NO_COLOR)"
	@./scripts/coverage

testcoverage:
	@echo "$(OK_COLOR)==> Testing$(NO_COLOR)"
	@./scripts/coverage --html

testcoveralls:
	@echo "$(OK_COLOR)==> Testing with coveralls$(NO_COLOR)"
	@./scripts/coverage --coveralls

lint:
	@echo "$(OK_COLOR)==> Linting$(NO_COLOR)"
	@golint .

resetdb:
	@echo "$(OK_COLOR)==> Resetting the Database$(NO_COLOR)"
	pgmgr db drop || true
	pgmgr db create || true
	pgmgr db migrate || true

setup: deps format lint resetdb

runtest: setup test

runtestcoveralls: setup testcoveralls
