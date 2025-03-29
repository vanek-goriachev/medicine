package rest

import (
	"context"
	"fmt"

	"template/internal/appcore"
)

const writeFail = "failed to write to stdout: %w"

func (a *App) initialize(ctx context.Context) error { //nolint:revive,unparam // going to fix later
	_, err := a.printStringToStdout("Initializing application\n")
	if err != nil {
		return fmt.Errorf(writeFail, err)
	}

	// Загрузить из файловой системы различные данные (конфиги, TLS сертификаты, креды итд)
	_, err = a.printStringToStdout("Loading filesystem data\n") // No logger here since it is not initialized yet
	if err != nil {
		return fmt.Errorf(writeFail, err)
	}

	// Собрать зависимости (IAM, телеметрия [логирование, трейсинг, метрики], БД [хранилище, кэш, очереди])
	applicationDependencies := appcore.NewDependencies()

	err = applicationDependencies.Initialize()
	if err != nil {
		return fmt.Errorf("failed to initialize dependencies: %w", err)
	}

	// Собрать ядро приложения (бизнес логику)
	a.appCore = appcore.NewCore(applicationDependencies)
	a.appCore.Initialize()

	// Мигрировать БД

	// Собрать сервер

	return nil
}
