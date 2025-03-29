# Настройка CI
1) Создайте папку .github/workflows в корне вашего репозитория.
2) В этой папке создайте файл с именем, например, lint_and_test.yml
3) Пример конфигурации для запуска тестов и линтера вы можете найти ниже или в файле .github/workflows/lint_and_test.yml
```yml
name: "Go project CI"

on:
  push:
    branches: ["main"]
  pull_request:
    branches: ["*"]

jobs:
  lint:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      - name: Install Go 1.24.x
        uses: actions/setup-go@v5
        with:
          go-version: '1.24.x'
      - name: Install python 3.13.x
        uses: actions/setup-python@v5
        with:
          python-version: '3.13.x'
      - name: Init linters
        run: |
          make init_linters
      - name: Run linters
        run: |
          make lint

  test:
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v4
      - name: Install Go 1.24.x
        uses: actions/setup-go@v5
        with:
          go-version: '1.24.x'
      - name: Run tests
        run: |
          make go_all_tests
```

# Тесты использующие базу данных
1)  Если вашим тестам необходима реальная база данных, то вы можете поднять ее при помощи docker-compose.
    Пример такого теста можно посмотреть в файле internal/tooling/example/example_db_test.go
2)  Для корректной работы тестов в CI следует поднять базу данных внутри соответствующей джобы.
    Для этого в джобу следует добавить код ниже (блоки services и env)

   Переменные внутри test:services:postgres:env отвечают за переменные окружения при поднятии базы данных
   Переменные внутри test:env отвечают за переменные окружения при запуске тестов

   ```yml
     test:
       runs-on: ...
       services:
         postgres:
           image: postgres:17
           ports:
             - "5432:5432"
           env:
             POSTGRES_DB: template_db_test
             POSTGRES_USER: template_user
             POSTGRES_PASSWORD: template_password
       env:
         DATABASE_URL: "postgres://template_user:template_password@localhost:5432/template_db_test"
       steps:
         ...
   ```

   В примере теста в файле ./internal/tooling/example/example_db_test.go происходит чтение только переменной DATABASE_URL
   ```go
    func (t *ExampleDBTestSuite) Test_GreenPath() {
        ...
        conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
        ...
    }
```
