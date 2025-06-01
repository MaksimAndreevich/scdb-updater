package internal

import (
	"time"

	"gitlab.com/scdb/core/config"

	logger "gitlab.com/scdb/core/logger"
	database "gitlab.com/scdb/database/services"
)

func Seed() {
	start := time.Now()
	data := GetDataParsedXML()

	config.LoadConfig()
	db, _ := database.Connect()
	// insertOrganisatingQuery := utils.LoadSQLFile("internal/database/sql/insert/educationOrganisation.sql")

	// Начинаем транзакцию — все последующие Exec будут частью одной атомарной операции
	tx, err := db.Begin()
	if err != nil {
		logger.Fatal("Ошибка при начале транзакции:", err)
	}

	// Откладываем коммит транзакции на конец функции:
	// когда функция завершится, попробуем зафиксировать изменения
	defer func() {
		if err := tx.Commit(); err != nil {
			logger.Fatal("Ошибка при коммите транзакции:", err)
		}
	}()

	insertQuery := `INSERT INTO education_organizations (
			id, full_name, short_name, head_edu_org_id, is_branch,
			post_address, phone, fax, email, web_site,
			ogrn, inn, kpp, head_post, head_name,
			form_name, kind_name, type_name, region_name,
			federal_district_short_name, federal_district_name
		) VALUES (
			$1, $2, $3, $4, $5,
			$6, $7, $8, $9, $10,
			$11, $12, $13, $14, $15,
			$16, $17, $18, $19,
			$20, $21
		) ON CONFLICT (id) DO NOTHING;
`
	// Подготавливаем запрос один раз — база его распарсит и создаст план выполнения
	stmt, err := tx.Prepare(insertQuery)

	if err != nil {
		logger.Fatal("Ошибка при подготовке запроса:", err)
	}

	// Закроем подготовленный запрос, когда выйдем из функции
	defer stmt.Close()

	var totalSuccess int = 0

	for i, cert := range data.Certificates {
		org := cert.ActualEducationOrganization
		isBranch := org.IsBranch == "1"

		logger.Info("Обработка сертификата №", i+1, ":", org.ShortName)

		result, err := stmt.Exec(
			org.ID, org.FullName, org.ShortName, org.HeadEduOrgId, isBranch,
			org.PostAddress, org.Phone, org.Fax, org.Email, org.WebSite,
			org.OGRN, org.INN, org.KPP, org.HeadPost, org.HeadName,
			org.FormName, org.KindName, org.TypeName, org.RegionName,
			org.FederalDistrictShortName, org.FederalDistrictName,
		)
		if err != nil {
			logger.Error("Ошибка при вставке организации:", err)
			continue
		}

		rowsAffected, err := result.RowsAffected()
		if err != nil {
			logger.Error("Ошибка при получении количества затронутых строк: ", err)
			continue
		}

		if rowsAffected > 0 {
			totalSuccess++
		}
	}

	logger.Info("Обработано сертификатов ", len(data.Certificates))
	logger.Success("Колличество вставленных организаций в таблицу ", totalSuccess)

	spendedTime := time.Since(start).Truncate(time.Second) // считаем прошедшее время

	logger.Info("Время выполнения: ", spendedTime)

	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Ошибка при закрытии базы данных:", err)
		}
	}()
}
