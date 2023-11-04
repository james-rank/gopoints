LINTER=golangci-lint

.PHONY: show-%
show-%:
	@echo ${$*}

.PHONY: lint
lint:
	betteralign -apply ./...
	$(LINTER) run --fix

.PHONY: test
test:
	go test -v -cover -coverprofile=coverage.out ./...
	sed -i '/mock_/d' coverage.out

.PHONY: generate-coverage
generate-coverage: test
	gocover-cobertura < coverage.out > coverage.xml
	mkdir -p coverage-reports
	go tool cover -html=coverage.out -o coverage-reports/coverage.html

.PHONY: view-coverage
view-coverage: generate-coverage
	go tool cover -html=coverage.out