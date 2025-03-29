PYTHON_INTERPRETER := python3
PYTHON_VENV_ACTIVATION_COMMAND := source ./venv/bin/activate
SETENV_COMMAND := set
ifeq ($(OS),Windows_NT)
	PYTHON_VENV_ACTIVATION_COMMAND := venv\Scripts\activate
	SETENV_COMMAND := bash set
endif

PRE_COMMIT_CONFIG := ./ci/linters/.pre-commit-config.yaml

.PHONY: init_project
init_project:
	@make init_linters

# call to Init linters
# Add a line to every tutorial file which would say to edit init project command in makefile

# Команда инициализирующая линтеры
.PHONY: init_linters
init_linters:
	$(PYTHON_INTERPRETER) -m venv venv
	$(PYTHON_VENV_ACTIVATION_COMMAND)
	pip install pre-commit
	pre-commit install --config $(PRE_COMMIT_CONFIG)
	pre-commit install -t pre-push --config $(PRE_COMMIT_CONFIG)

# Команда запускающая линтеры
.PHONY: lint
lint:
	$(PYTHON_VENV_ACTIVATION_COMMAND)
	pre-commit run --all-files --config $(PRE_COMMIT_CONFIG)

# Команда для запуска всех go тестов в папках test и internal
# с предварительным поднятием БД и установкой всех переменных окружения
.PHONY: go_all_tests
go_all_tests:
	go test -v ./tests/... ./internal/...

# При локальном запуске тестов необходимо поднять БД
export POSTGRES_DB:=template_db_test
export POSTGRES_USER:=template_user
export POSTGRES_PASSWORD:=template_password
export DATABASE_URL:=postgres://$(POSTGRES_USER):$(POSTGRES_PASSWORD)@localhost:5432/$(POSTGRES_DB)

.PHONY: local_go_all_tests
local_go_all_tests:
	docker-compose -f ./ci/tests/postgres.yml up -d
	@make go_all_tests
	docker-compose -f ./ci/tests/postgres.yml down
