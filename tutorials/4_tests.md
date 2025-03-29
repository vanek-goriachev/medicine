# Настройка тестов
В данном репозитории примеры тестов расположены в папках 'tests/example' и 'internal/tooling/example'
Для запуска тестов в Makefile добавлена команда go_all_tests
```makefile
.PHONY: go_all_tests
go_all_tests:
	go test -v ./tests/... ./internal/...
```

# Интеграционные тесты с Базой данных
Если вашим тестам необходима реальная база данных, то вы можете поднять ее при помощи docker-compose.
Пример такого теста можно посмотреть в файле internal/tooling/example/example_db_test.go
Для этого создайте файл ci/tests/postgres.yml
```yml
version: '3.9'

services:

  db:
    image: postgres
    restart: unless-stopped
    ports:
      - "5432:5432"
    environment:
      POSTGRES_DB: ${POSTGRES_DB}
      POSTGRES_USER: ${POSTGRES_USER}
      POSTGRES_PASSWORD: ${POSTGRES_PASSWORD}
```
И добавьте в Makefile команду для запуска тестов с указанием переменных окружения и запуском тестов с предварительным поднятием БД
```makefile
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
```
Переменные POSTGRES_DB, POSTGRES_USER, POSTGRES_PASSWORD попадают в контейнер после запуска команды docker-compose
Переменная DATABASE_URL читается кодом при подключении к БД
```go
func (t *ExampleDBTestSuite) Test_GreenPath() {
    ...
    conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
    ...
}
```
