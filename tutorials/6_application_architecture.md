На текущем этапе подготовки шаблонного репозитория я вижу архитектуру REST приложения следующим образом

Внимание: код ниже "схематичен". В нем опущена обработка ошибок, вывод сообщений в консоль и т.д.
Этот код получит свое развитие в следующих туториалах.

В некоторых местах код можно улучшить из сниппета 1 в сниппет 2, избавившись от комментариев

```go
func F() {
	// Сделать первое
	// Сделать второе
	// Сделать третье
}
```

```go
func F() {
    doFirst()
    doSecond()
    doThird()
}

func doFirst() {}
func doSecond() {}
func doThird() {}

```


# Предлагаемая структура приложения
1) Файл main.go
```go
package main

imports ...

func main() {
	terminal := os.NewTerminal()

	app := rest.NewApp(terminal)

	app.InitializeAndRun()
}
```
2) Пакет rest
```go
package rest

import (...)

type App struct {...}

func NewApp(...) *App {
	...
}

func (a *App) InitializeAndRun() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	a.initialize(ctx)

	go a.cancelContextOnReceiveInterruptingSignals(cancel)

	a.run(ctx)
}

func (a *App) initialize(ctx context.Context) error {
	// Загрузить из файловой системы различные данные (конфиги, TLS сертификаты, креды итд)
	...

	// Собрать зависимости (IAM, телеметрия [логирование, трейсинг, метрики], БД [хранилище, кэш, очереди])
	applicationDependencies := appcore.NewDependencies()
	err = applicationDependencies.Initialize()

	// Собрать ядро приложения (бизнес логику)
	a.appCore = appcore.NewCore(applicationDependencies)
	a.appCore.Initialize()

	// Мигрировать БД
	...

	// Собрать сервер
	...

	return nil
}

func (a *App) cancelContextOnReceiveInterruptingSignals(cancelFunc context.CancelFunc) {
	terminal := a.systemDependencies.Terminal
	<-terminal.SignalCh(1, syscall.SIGINT, syscall.SIGTERM)
	cancelFunc()
}

func (a *App) run(ctx context.Context) error {
	// Запустить сервер
	// Настроить Graceful Shutdown (в том числе телеметрии)
	return nil
}

```
3) Пакет appcore
```go
type Dependencies struct {
	// Telemetry
	// IAM
	// DB
}

func NewDependencies() *Dependencies {
	return &Dependencies{}
}

func (*Dependencies) Initialize() error {
	// Initialize all dependencies
	return nil
}

type Core struct {
	Dependencies *Dependencies
	// Gateways
	// Mappers
	// Actions
	// UserScenarios
}

func NewCore(dependencies *Dependencies) *Core {
	return &Core{
		Dependencies: dependencies,
	}
}

func (*Core) Initialize() {
	// Initialize core
}

```
