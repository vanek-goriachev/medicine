1) Настроить конфиги (можно смотреть примеры конфигов в этом проекте)
   .././ci/linters/.pre-commit-config.yaml - конфиг для pre-commit хуков
   .././ci/linters/.golangci.yml - конфиг для golangci-lint
   .././ci/linters/.golangci.most_strict.yml - конфиг для golangci-lint с наиболее строгими правилами
2) Запустить команду инициализации линтеров из Makefile (см. ниже)
   ```bash
   make init_linters
   ```
3) Проверить, что все хуки работают корректно сделав коммит
   ```bash
   git add . && git commit -m "Initial commit"
   ```

# Makefile для инициализации линтеров и pre-commit хуков
```makefile
PYTHON_INTERPRETER := python3
PYTHON_VENV_ACTIVATION_COMMAND := venv/bin/activate
ifeq ($(OS),Windows_NT)
	PYTHON_VENV_ACTIVATION_COMMAND := venv\Scripts\activate
endif

PRE_COMMIT_CONFIG := ./ci/linters/.pre-commit-config.yaml

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

```

# То же самое, что в Makefile, но вручную с подробным описанием каждого шага
1) Создать виртуальное окружение python для проекта и активировать его
   ```bash
   python3 -m venv venv
   source venv/bin/activate
   # или для windows
   # venv\Scripts\activate
   ```
2) Установить утилиту pre-commit
   ```bash
   pip install pre-commit
   ```
3) Настроить конфиги
   .././ci/linters/.pre-commit-config.yaml - конфиг для pre-commit хуков
   .././ci/linters/.golangci.yml - конфиг для golangci-lint
   .././ci/linters/.golangci.most_strict.yml - конфиг для golangci-lint с наиболее строгими правилами
4) Установить pre-commit и pre-push хуки с использованием ранее написанного конфига
   ```bash
   pre-commit install --config ./ci/linters/.pre-commit-config.yaml
   pre-commit install -t pre-push --config ./ci/linters/.pre-commit-config.yaml
   ```
5) Проверить, что все хуки работают корректно сделав коммит
   ```bash
   git add .
   git commit -m "Initial commit"
   ```
