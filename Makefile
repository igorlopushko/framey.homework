run-all:
	@LOG_LEVEL=debug bash -c 'go run main.go'
run-speed:
	@LOG_LEVEL=debug bash -c 'go run main.go -p speedtest.net'
run-fast:
	@LOG_LEVEL=debug bash -c 'go run main.go -p fast.com'
test:
	@go test -v ./... -coverprofile cover.out