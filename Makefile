PACKAGE_NAME=github.com/meisterunused/badgeify

bin/badgeify:
	go build -o bin/badgeify cmd/badgeify.go

.PHONY: run
run:
	@go run cmd/badgeify.go "coverage" "89%" > test.svg

.PHONY: test
test:
	@go test $(PACKAGE_NAME)

.PHONY: clean
clean:
	rm -f bin/badgeify:
