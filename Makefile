GOFMT_FILES?=$$(find . -not -path "./vendor/*" -type f -name '*.go')
TEST?=./...

fmt:
	gofmt -w $(GOFMT_FILES)
	go vet ./...

test:
	go test $(TEST)