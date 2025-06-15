package database

import (
	"github.com/lib/pq"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func GetOrgTypesMap() map[string]models.EducationTypeShortInfo {
	rows, err := DB.Query("SELECT key, keywords FROM education_types")
	if err != nil {
		logger.Fatal("[GET ORG TYPES] Ошибка при получении типов образовательных учреждений: ", err)
	}

	orgTypesMap := make(map[string]models.EducationTypeShortInfo)

	for rows.Next() {
		var orgType models.EducationTypeShortInfo
		var keywords []string
		err := rows.Scan(&orgType.Key, pq.Array(&keywords))
		if err != nil {
			logger.Fatal("[GET ORG TYPES] Ошибка при сканировании типа образовательного учреждения: ", err)
		}

		orgType.Keywords = keywords
		orgTypesMap[orgType.Key] = orgType

	}

	return orgTypesMap
}
