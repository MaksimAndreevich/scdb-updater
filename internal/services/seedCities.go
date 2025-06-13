package services

import (
	"encoding/json"
	"os"
	"strings"

	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func SeedCities() {

	data, err := os.ReadFile("./data/cities.json")

	if err != nil {
		logger.Fatal("Ошибка при чтении файла cities.json: ", err)
	}

	var cities []models.City

	if err := json.Unmarshal(data, &cities); err != nil {
		logger.Fatal("Ошибка при разборе JSON городов: ", err)
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
		INSERT INTO cities (
			address, postal_code, country, federal_district_name, region_type,
			region, area_type, area, city_type, city, kladr_id, fias_id, fias_level, capital_marker,
			okato, oktmo, tax_office, timezone, geo_lat, geo_lon,
			population, foundation_year, fk_region_id, fk_federal_district_id
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
			$14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
		)
		ON CONFLICT (fias_id) DO NOTHING`

	stmt, err := tx.Prepare(QUERY)
	if err != nil {
		logger.Fatal("Ошибка подготовки запроса для импорта городов: %w", err)
	}
	defer stmt.Close()

	regionsMap := database.GetRegionsMap()

	// Вставляем данные
	for i, city := range cities {

		// Получаем ID региона по названию
		regionName := strings.Fields(city.RegionName)[0] // Берем только первое слово
		regionInfo, ok := regionsMap[regionName]
		if !ok {
			logger.Fatal("Регион не найден для города ", city.CityName, " (регион: ", city.RegionName, ")")
		}

		_, err = stmt.Exec(
			city.Address,
			city.PostalCode,
			city.Country,
			city.FederalDistrictName,
			city.RegionType,
			city.RegionName,
			city.AreaType,
			city.Area,
			city.CityType,
			city.CityName,
			city.KladrID,
			city.FiasID,
			city.FiasLevel,
			city.CapitalMarker,
			city.OKATO,
			city.OKTMO,
			city.TaxOffice,
			city.Timezone,
			city.GeoLat,
			city.GeoLon,
			city.Population,
			city.FoundationYear,
			regionInfo.ID,
			regionInfo.FederalDistrictID,
		)
		if err != nil {
			logger.Fatal("Ошибка вставки города ", i, err)
		}

	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		logger.Fatal("Ошибка завершения транзакции: %w", err)
	}

	logger.Success("Импорт городов завершен. Всего импортировано: ", len(cities))

}
