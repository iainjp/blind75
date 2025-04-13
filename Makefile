test:
	(cd go && go tool gotestsum)

lint:
	(cd go && go tool golangci-lint run)

format:
	(cd go && go tool golangci-lint fmt)
