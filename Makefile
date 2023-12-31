init:
	@echo 'Initial Setup'
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin v1.55.2
	git config core.hooksPath githooks

test:
	@echo 'Run Unit Testing'
	go test -v -coverpkg=./... -coverprofile=coverage.out ./...