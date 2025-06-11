package services

import (
	"encoding/json"
	"os"
	"strings"

	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func SeedRegions(districts []models.FederalDistrict) {
	data, err := os.ReadFile("./data/regions.json")

	if err != nil {
		logger.Fatal("Ошибка при чтении файла regions.json: ", err)
	}

	var regions []models.Region

	if err := json.Unmarshal(data, &regions); err != nil {
		logger.Fatal("Ошибка при разборе JSON регионов: ", err)
	}

	// Начинаем транзакцию
	tx, err := database.DB.Begin()
	if err != nil {
		logger.Fatal("Ошибка начала транзакции: ", err)
	}

	defer tx.Rollback()

	QUERY := `
       INSERT INTO regions (    
		name, label, type, type_short, content_type, kladr_id, okato, oktmo, guid, code, iso_3166_2, 
		population, year_founded, area, fullname, name_en, fk_federal_district_id,
		namecase_nominative, namecase_genitive, namecase_dative, namecase_accusative, 
		namecase_ablative, namecase_prepositional, namecase_locative,
		capital_name, capital_label, capital_kladr_id, capital_okato, capital_oktmo, capital_content_type
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17,
			$18, $19, $20, $21, $22, $23, $24,
			$25, $26, $27, $28, $29, $30
		)
		ON CONFLICT (kladr_id) DO NOTHING`

	stmt, err := tx.Prepare(QUERY)
	if err != nil {
		logger.Fatal("Ошибка подготовки запроса: %w", err)
	}
	defer stmt.Close()

	// Создаем маппинг первых слов названий округов на их ID
	districtMap := make(map[string]int)
	for _, district := range districts {
		// Берем первое слово из названия
		firstWord := strings.Fields(district.Name)[0]
		districtMap[firstWord] = district.ID

	}

	totalSuccess := 0

	// Вставляем данные
	for i, region := range regions {
		// Получаем ID федерального округа по первому слову
		federalDistrictID, ok := districtMap[region.FederalDistrictName]
		if !ok {
			logger.Fatal("Федеральный округ не найден для региона ", region.Name, " (округ: ", region.FederalDistrictName, ")")
		}

		_, err := stmt.Exec(
			region.Name,
			region.Label,
			region.Type,
			region.TypeShort,
			region.ContentType,
			region.RegionKladrID,
			region.OKATO,
			region.OKTMO,
			region.GUID,
			region.Code,
			region.ISO3166_2,
			region.Population,
			region.YearFounded,
			region.Area,
			region.Fullname,
			region.NameEn,
			federalDistrictID,
			region.NamecaseNominative,
			region.NamecaseGenitive,
			region.NamecaseDative,
			region.NamecaseAccusative,
			region.NamecaseAblative,
			region.NamecasePrepositional,
			region.NamecaseLocative,
			region.CapitalName,
			region.CapitalLabel,
			region.CapitalKladrID,
			region.CapitalOKATO,
			region.CapitalOKTMO,
			region.CapitalContentType,
		)

		if err != nil {
			logger.Fatal("Ошибка вставки региона ", i, err)
		}

		totalSuccess++
	}

	// Завершаем транзакцию
	if err := tx.Commit(); err != nil {
		logger.Fatal("Ошибка завершения транзакции: %w", err)
	}
	logger.Success("Импорт регонов. Вствлено регионов:  ", totalSuccess)

}
