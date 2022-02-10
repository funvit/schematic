.PHONY:help
help:
	@echo Known rules:
	@echo
	@echo make install

.PHONY:install
install:
	go mod tidy
	go mod download
	go vet ./...
	go install ./cmd/...

.PHONY:check-generator
check-generator: install
	cd ./example && make gen
	#cd ./example && make test
