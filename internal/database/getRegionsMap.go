package database

import (
	"strings"

	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func GetRegionsMap() map[string]models.RegionShortInfo {
	// Получаем регионы для дальнейшего матчинга с городами
	rows, err := DB.Query("SELECT id, name, fk_federal_district_id FROM regions")
	if err != nil {
		logger.Fatal("[GET REGIONS] Ошибка при получении регионов во время вставки городов: ", err)
	}

	var regions []models.RegionShortInfo

	for rows.Next() {
		var region models.RegionShortInfo

		err := rows.Scan(&region.ID, &region.Name, &region.FederalDistrictID)
		if err != nil {
			logger.Fatal("[GET REGIONS] Ошибка при сканировнии региона во время вставки городов: ", err)
		}

		regions = append(regions, region)
	}

	if len(regions) == 0 {
		logger.Error("[GET REGIONS] Нет регионов в базе")
	}

	regionsMap := make(map[string]models.RegionShortInfo)
	for _, region := range regions {
		// Берем название региона (Алтайский, Дагестан, Владимирская) как ключ и сохраняем в мапу
		regionsMap[strings.Fields(region.Name)[0]] = models.RegionShortInfo{
			ID:                region.ID,
			Name:              region.Name,
			FederalDistrictID: region.FederalDistrictID,
		}
	}

	return regionsMap
}
