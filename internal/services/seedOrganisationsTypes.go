package services

import (
	"encoding/json"
	"os"

	"github.com/lib/pq"
	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func SeedOrganisationsTypes() {

	data, err := os.ReadFile("./data/org_types.json")

	if err != nil {
		logger.Fatal("[SEED ORGANISATIONS TYPES] Ошибка при чтении файла org_types.json: ", err)
	}

	var orgTypes []models.EducationType

	if err := json.Unmarshal(data, &orgTypes); err != nil {
		logger.Fatal("[SEED ORGANISATIONS TYPES] Ошибка при разборе JSON типов образовательных учреждений: ", err)
	}

	// Начинаем транзакцию
	tx, err := database.DB.Begin()
	if err != nil {
		logger.Fatal("Ошибка начала транзакции: ", err)
	}

	defer tx.Rollback()

	QUERY := `
    INSERT INTO education_types (    
        key, title, level, ownership_forms, keywords
    ) VALUES (
        $1, $2, $3, $4, $5
    )
    ON CONFLICT (key) DO UPDATE SET
        title = EXCLUDED.title,
        level = EXCLUDED.level,
        ownership_forms = EXCLUDED.ownership_forms,
        keywords = EXCLUDED.keywords`

	stmt, err := tx.Prepare(QUERY)
	if err != nil {
		logger.Fatal("[SEED ORGANISATIONS TYPES] Ошибка подготовки запроса: %w", err)
	}
	defer stmt.Close()

	for i, orgType := range orgTypes {

		_, err := stmt.Exec(
			orgType.Key,
			orgType.Title,
			orgType.Level,
			pq.Array(orgType.OwnershipForms),
			pq.Array(orgType.Keywords),
		)

		if err != nil {
			logger.Fatal("[SEED ORGANISATIONS TYPES] Ошибка вставки ", i, err)
		}

	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		logger.Fatal("[SEED ORGANISATIONS TYPES] Ошибка завершения транзакции: %w", err)
	}
	logger.Success("Импорт типов образовательных учреждений завершен")
}
