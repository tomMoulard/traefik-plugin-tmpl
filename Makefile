.POSIX:

.PHONY: all
all: lint test

.PHONY: lint
lint:
	golangci-lint run

.PHONY: test
test:
	go test -v -cover ./...

.PHONY: yaegi_test
yaegi_test:
	yaegi test -v .

.PHONY: vendor
vendor:
	go mod vendor

.PHONY: entr
# https://github.com/eradman/entr
entr:
	find | entr -r -s "docker compose up --remove-orphans"

.PHONY: clean
clean:
	$(RM) -r ./vendor
