devRun:
	@echo "start to build"
	@go build -o ./bin/main ./src/cmd/web/
	@echo "start to run"
	@./bin/main

.PHONY:
	devRun