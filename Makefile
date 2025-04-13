test:
	(cd go/cmd/exercises && go tool gotestsum)

lint:
	(cd go/cmd/exercises && go tool golangci-lint run)

format:
	(cd go/cmd/exercises && go tool golangci-lint fmt)
