PYTHON_INTERPRETER := python3
PYTHON_VENV_ACTIVATION_COMMAND := source venv/bin/activate
SETENV_COMMAND := set
ifeq ($(OS),Windows_NT)
	PYTHON_VENV_ACTIVATION_COMMAND := venv\Scripts\activate
	SETENV_COMMAND := bash set
endif

PRE_COMMIT_CONFIG := ./build/ci/linters/.pre-commit-config.yaml
GOLANGCI_LINT_CONFIG := ./build/ci/linters/.golangci.yml

.PHONY: init_project
init_project:
	@make init_githooks

# Команда инициализирующая линтеры
.PHONY: init_githooks
init_githooks:
	(\
		$(PYTHON_INTERPRETER) -m venv venv; \
		$(PYTHON_VENV_ACTIVATION_COMMAND); \
		pip install pre-commit; \
		pre-commit install --config $(PRE_COMMIT_CONFIG); \
		pre-commit install -t pre-push --config $(PRE_COMMIT_CONFIG); \
	)


# Команда для запуска всех go тестов в папках test и internal
# с предварительным поднятием БД и установкой всех переменных окружения
.PHONY: go_all_tests
go_all_tests:
	go test -v ./tests/... ./internal/...

# При локальном запуске тестов необходимо поднять БД
export POSTGRES_DB:=medicine_db_test
export POSTGRES_USER:=medicine_user
export POSTGRES_PASSWORD:=medicine_password
export DATABASE_URL:=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)

.PHONY: local_go_all_tests
local_go_all_tests:
	docker-compose -f ./build/ci/tests/postgres.yml up -d
	@make go_all_tests
	docker-compose -f ./build/ci/tests/postgres.yml down

.PHONY: tools
tools:
	go generate tools/tools.go


DIRECTORIES_TO_LINT := ./internal/... ./pkg/... ./cmd/... ./tests/...
# Команда запускающая линтеры
.PHONY: lint
lint:
	bin/goimports -w cmd internal pkg tests tools
	bin/golangci-lint run --config $(GOLANGCI_LINT_CONFIG) $(DIRECTORIES_TO_LINT)

.PHONY: lint_with_fix
lint_fix:
	bin/goimports -w cmd internal pkg tests tools
	bin/golangci-lint run --config $(GOLANGCI_LINT_CONFIG) $(DIRECTORIES_TO_LINT) --fix

.PHONY: local_run
local_run:
	docker-compose -f ./build/local_development/service.yml up --build -d

.PHONY: mocks
mocks:
	bin/mockery --config="./build/ci/tests/.mockery.yml"
	touch mocks/.coverignore
