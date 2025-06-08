package services

import (
	"encoding/json"
	"os"
	"strconv"

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
		logger.Fatal("Ошибка при разборе JSON: ", err)
	}

	db, err := database.Connect()

	if err != nil {
		logger.Fatal("Ошибка при подключении к базе данных: ", err)
	}

	// Начинаем транзакцию
	tx, err := db.Begin()
	if err != nil {
		logger.Fatal("Ошибка начала транзакции: ", err)
	}

	defer tx.Rollback()

	QUERY := `
		INSERT INTO cities (
			address, postal_code, country, federal_district, region_type,
			region, area_type, area, city_type, city, settlement_type,
			settlement, kladr_id, fias_id, fias_level, capital_marker,
			okato, oktmo, tax_office, timezone, geo_lat, geo_lon,
			population, foundation_year
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13,
			$14, $15, $16, $17, $18, $19, $20, $21, $22, $23, $24
		)
		ON CONFLICT (fias_id) DO NOTHING`

	stmt, err := tx.Prepare(QUERY)
	if err != nil {
		logger.Fatal("Ошибка подготовки запроса: %w", err)
	}
	defer stmt.Close()

	// Вставляем данные
	for i, city := range cities {

		_, err = stmt.Exec(
			city.Address,
			strconv.Itoa(city.PostalCode),
			city.Country,
			city.FederalDistrict,
			city.RegionType,
			city.Region,
			city.AreaType,
			city.Area,
			city.CityType,
			city.City,
			city.SettlementType,
			city.Settlement,
			strconv.FormatInt(city.KladrID, 10),
			city.FiasID,
			city.FiasLevel,
			city.CapitalMarker,
			strconv.FormatInt(city.Okato, 10),
			strconv.FormatInt(city.Oktmo, 10),
			strconv.Itoa(city.TaxOffice),
			city.Timezone,
			city.GeoLat,
			city.GeoLon,
			city.Population,
			city.FoundationYear,
		)
		if err != nil {
			logger.Error("Ошибка вставки города ", i, err)
		}

	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		logger.Fatal("Ошибка завершения транзакции: %w", err)
	}

	logger.Success("Импорт завершен. Всего импортировано городов: ", len(cities))

}
