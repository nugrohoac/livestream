#unit test only, mysql is not running because integration test
unittest:
	@echo "Start running unit test"
	@go test ./... --short -cover -race
	@echo "Unit test done"

test:
	@echo "Start running Integration test"
	@go test ./...  -cover -race
	@echo "Integration test done"