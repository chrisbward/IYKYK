#--- Help ---
help:
	@echo 
	@echo Makefile targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' Makefile | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo

test: ## Runs the test cases
	@echo "+ $@"
	@bash -c "go test -v ./pkg/..."

mockgen: ## Generates the mocks
	@echo "+ $@"
	bash -c "mockgen -source=pkg/entities/interfaces.go -destination=pkg/entities/_mocks/interfaces_mocked.go -package mockedinterfaces"
