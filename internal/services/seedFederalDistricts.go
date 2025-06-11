package services

import (
	"encoding/json"
	"os"

	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func SeedFederalDistricts() []models.FederalDistrict {

	data, err := os.ReadFile("./data/districts.json")

	if err != nil {
		logger.Fatal("Ошибка при чтении файла districts.json: ", err)
	}

	var districts []models.FederalDistrict

	if err := json.Unmarshal(data, &districts); err != nil {
		logger.Fatal("Ошибка при разборе JSON округов: ", err)
	}

	if err != nil {
		logger.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	// Начинаем транзакцию
	tx, err := database.DB.Begin()
	if err != nil {
		logger.Fatal("Ошибка начала транзакции: ", err)
	}

	defer tx.Rollback()

	QUERY := `
	INSERT INTO federal_districts (	
		short_name, name
	) VALUES (
		$1, $2
	)
	ON CONFLICT (short_name) DO UPDATE 
	SET name = EXCLUDED.name
	RETURNING id`

	stmt, err := tx.Prepare(QUERY)
	if err != nil {
		logger.Fatal("Ошибка подготовки запроса: %w", err)
	}
	defer stmt.Close()

	// Вставляем данные
	for i := range districts {
		var id int
		err = stmt.QueryRow(
			districts[i].ShortName,
			districts[i].Name,
		).Scan(&id)

		if err != nil {
			logger.Error("Ошибка вставки федерального округа ", i, err)
		}

		districts[i].ID = id

	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		logger.Fatal("Ошибка завершения транзакции: %w", err)
	}

	logger.Success("Импорт федеральных округов завершен. Всего импортировано: ", len(districts))

	return districts
}
