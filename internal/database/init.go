package database

import (
	"gitlab.com/scdb/updater/internal/config"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/utils"
)

func Init() {
	logger.Info("________________Инициализация базы данных Postgres________________")
	config.LoadConfig()
	db, error := Connect()

	if error != nil {
		logger.Error("Ошибка подключения к DB")
	}

	tables := []struct {
		Name     string
		FilePath string
	}{
		{
			Name:     "education_organizations",
			FilePath: "internal/database/sql/create/educationOrganizationsTable.sql",
		},
		// Добавляй сюда другие таблицы по мере надобности
	}

	for _, table := range tables {
		query := utils.LoadSQLFile(table.FilePath)

		_, err := db.Exec(query)
		if err != nil {
			logger.Fatal("Ошибка при создании таблицы", table.Name+":", err)
		}

		logger.Info("Таблица успешно создана: ", table.Name)
	}

	// Отключаемся от db в конце завершения работы функции
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Ошибка при закрытии базы данных:", err)
		}
	}()
}
