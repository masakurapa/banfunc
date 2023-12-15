install:
	go install ./cmd/banfunc/...

lint:
	golangci-lint run

test:
	go test ./...


.PHONY: test lint install
