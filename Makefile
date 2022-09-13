run-all:
	@LOG_LEVEL=debug bash -c 'go run main.go'
run-ookla:
	@LOG_LEVEL=debug bash -c 'go run main.go -p ookla'
run-netflix:
	@LOG_LEVEL=debug bash -c 'go run main.go -p netflix'
test:
	@go test -v ./... -coverprofile cover.out
bench:
	@go test -v ./... -bench=.  -benchmem -run=^#
godoc:
	@godoc -http=0.0.0.0:6060 -v -timestamps=true -links=true -play=true
lint:
	@golangci-lint run -v --config ./.golangci.yml