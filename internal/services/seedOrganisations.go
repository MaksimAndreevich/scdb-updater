package services

import (
	"time"

	"github.com/lib/pq"
	"gitlab.com/scdb/updater/internal/config"
	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
)

// Размер пакета для вставки данных
// Оптимальное значение зависит от размера данных и доступной памяти
// 1000 - хороший компромисс между скоростью и использованием памяти
const BatchSize = 1000

func SeedOrganisations() {
	start := time.Now()
	config.LoadConfig()

	db, _ := database.Connect()

	data := GetDataParsedXML()

	// Начинаем транзакцию
	// Транзакция обеспечивает атомарность операций и лучшую производительность
	// Все операции будут выполнены как единое целое
	tx, err := db.Begin()
	if err != nil {
		logger.Fatal("Ошибка при начале транзакции:", err)
	}

	// Отключаем триггеры перед массовой вставкой
	// Это значительно ускоряет процесс вставки, так как триггеры не будут выполняться
	// для каждой строки. Триггеры могут выполнять дополнительные проверки или
	// обновления других таблиц, что замедляет процесс
	_, err = tx.Exec("ALTER TABLE education_organizations DISABLE TRIGGER ALL;")
	if err != nil {
		logger.Fatal("Ошибка при отключении триггеров:", err)
	}

	// Создаем временную таблицу для COPY
	// Временная таблица будет автоматически удалена после завершения транзакции
	// (ON COMMIT DROP). Это позволяет:
	// 1. Быстро загрузить данные без проверки ограничений
	// 2. Избежать конфликтов с существующими данными
	// 3. Оптимизировать процесс вставки
	_, err = tx.Exec(`
		CREATE TEMP TABLE temp_education_organizations (
			id text,
			full_name text,
			short_name text,
			head_edu_org_id text,
			is_branch boolean,
			post_address text,
			phone text,
			fax text,
			email text,
			web_site text,
			ogrn text,
			inn text,
			kpp text,
			head_post text,
			head_name text,
			form_name text,
			kind_name text,
			type_name text,
			region_name text,
			federal_district_short_name text,
			federal_district_name text
		) ON COMMIT DROP;
	`)
	if err != nil {
		logger.Fatal("Ошибка при создании временной таблицы:", err)
	}

	// Подготавливаем COPY операцию
	// pq.CopyIn создает специальный statement для быстрой загрузки данных
	// Это намного быстрее обычных INSERT запросов, так как:
	// 1. Использует бинарный протокол PostgreSQL
	// 2. Минимизирует количество сетевых запросов
	// 3. Оптимизирует использование памяти
	stmt, err := tx.Prepare(pq.CopyIn("temp_education_organizations",
		"id", "full_name", "short_name", "head_edu_org_id", "is_branch",
		"post_address", "phone", "fax", "email", "web_site",
		"ogrn", "inn", "kpp", "head_post", "head_name",
		"form_name", "kind_name", "type_name", "region_name",
		"federal_district_short_name", "federal_district_name"))

	if err != nil {
		logger.Fatal("Ошибка при подготовке COPY:", err)
	}

	for i, cert := range data.Certificates {
		org := cert.ActualEducationOrganization
		isBranch := org.IsBranch == "1"

		logger.Info("Обработка сертификата №", i+1, ":", org.ShortName)

		// Добавляем строку в COPY
		// Каждый Exec добавляет одну строку в буфер COPY
		// Данные не отправляются в базу сразу, а накапливаются в буфере
		_, err = stmt.Exec(
			org.ID,
			org.FullName,
			org.ShortName,
			org.HeadEduOrgId,
			isBranch,
			org.PostAddress,
			org.Phone,
			org.Fax,
			org.Email,
			org.WebSite,
			org.OGRN,
			org.INN,
			org.KPP,
			org.HeadPost,
			org.HeadName,
			org.FormName,
			org.KindName,
			org.TypeName,
			org.RegionName,
			org.FederalDistrictShortName,
			org.FederalDistrictName,
		)
		if err != nil {
			logger.Error("Ошибка при добавлении строки в COPY:", err)
			continue
		}

	}

	// Завершаем COPY операцию
	// Пустой Exec() сигнализирует о завершении COPY
	// В этот момент все накопленные данные отправляются в базу
	_, err = stmt.Exec()
	if err != nil {
		logger.Fatal("Ошибка при завершении COPY:", err)
	}
	stmt.Close()

	// Вставляем данные из временной таблицы в основную
	// Используем DISTINCT ON для исключения дубликатов по id
	// ON CONFLICT DO NOTHING пропускает записи, которые уже существуют
	_, err = tx.Exec(`
		INSERT INTO education_organizations
		SELECT DISTINCT ON (id) *
		FROM temp_education_organizations
		ON CONFLICT (id) DO NOTHING;
	`)
	if err != nil {
		logger.Fatal("Ошибка при вставке из временной таблицы:", err)
	}

	// Включаем триггеры обратно
	// Важно включить триггеры после завершения массовой вставки
	// для обеспечения целостности данных
	_, err = tx.Exec("ALTER TABLE education_organizations ENABLE TRIGGER ALL;")
	if err != nil {
		logger.Fatal("Ошибка при включении триггеров:", err)
	}

	// Коммитим транзакцию
	// Все изменения становятся видимыми другим транзакциям
	// Временная таблица автоматически удаляется
	if err := tx.Commit(); err != nil {
		logger.Fatal("Ошибка при коммите транзакции:", err)
	}

	// Выводим статистику выполнения
	logger.Info("Обработано сертификатов ", len(data.Certificates))

	// Выводим общее время выполнения
	spendedTime := time.Since(start).Truncate(time.Second)
	logger.Info("Время выполнения: ", spendedTime)

	// Закрываем соединение с базой данных
	defer func() {
		if err := db.Close(); err != nil {
			logger.Error("Ошибка при закрытии базы данных:", err)
		}
	}()
}
