package database

import (
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func GetCitiesMap() map[string]models.CityShortInfo {
	rows, err := DB.Query("SELECT city, fias_id, fk_region_id, fk_federal_district_id FROM cities")
	if err != nil {
		logger.Fatal("[GET CITIES] Ошибка при получении городов: ", err)
	}

	var cities []models.CityShortInfo

	for rows.Next() {
		var city models.CityShortInfo

		err := rows.Scan(&city.CityName, &city.FiasID, &city.RegionID, &city.FederalDistrictID)
		if err != nil {
			logger.Fatal("[GET CITIES] Ошибка при сканировнии города: ", err)
		}

		cities = append(cities, city)
	}

	if len(cities) == 0 {
		logger.Error("[GET CITIES] Нет городов в базе")
	}

	citiesMap := make(map[string]models.CityShortInfo)
	for _, city := range cities {
		// Берем название города (Москва, Санкт-Петербург, Екатеринбург) как ключ и сохраняем в мапу
		citiesMap[city.CityName] = models.CityShortInfo{
			CityName:          city.CityName,
			FiasID:            city.FiasID,
			RegionID:          city.RegionID,
			FederalDistrictID: city.FederalDistrictID,
		}
	}

	return citiesMap
}
