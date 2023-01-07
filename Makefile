# Target creators
.PHONY: go_get go_test_integration

#
# Go targets.
#
go_get:
	@echo '>>> Getting go modules.'
	@env go mod download

go_test_integration:
	@echo ">>> Running integration tests."
	@env go test -v -p 1 -tags="integration" ./tests/integration/...
		.
