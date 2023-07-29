test:
	@go test -race -failfast -count=1 -timeout=30s -v -cover

.PHONY: test
