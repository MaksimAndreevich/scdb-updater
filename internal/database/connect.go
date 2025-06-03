package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"gitlab.com/scdb/updater/internal/config"
	"gitlab.com/scdb/updater/internal/logger"
)

func Connect() (*sql.DB, error) {
	// Формируем строку для подключения к pq с оптимизированными параметрами
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable "+
			// Таймаут на установление соединения - 10 секунд
			"connect_timeout=10 "+
			// Убираем ограничение на время выполнения запроса (0 = без ограничений)
			"statement_timeout=0 "+
			// Убираем ограничение на время простоя в транзакции (0 = без ограничений)
			"idle_in_transaction_session_timeout=0 "+
			// Убираем ограничение на время ожидания блокировки (0 = без ограничений)
			"lock_timeout=0 "+
			// Имя приложения для идентификации в логах PostgreSQL
			"application_name=updater",
		config.AppConfig.DBHost,
		config.AppConfig.DBPort,
		config.AppConfig.DBUser,
		config.AppConfig.DBPassword,
		config.AppConfig.DBName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		logger.Fatal("Ошибка подключения к базе данных:", err)
	}

	// Устанавливаем оптимальные параметры пула соединений
	// Максимальное количество одновременных соединений с базой данных
	// 25 - оптимальное значение для большинства случаев
	db.SetMaxOpenConns(25)
	// Максимальное количество простаивающих соединений в пуле
	// Должно быть равно MaxOpenConns для оптимальной производительности
	db.SetMaxIdleConns(25)
	// Время жизни соединения (0 = без ограничений)
	// Важно для длительных операций, чтобы соединения не закрывались
	db.SetConnMaxLifetime(0)
	// Время простоя соединения (0 = без ограничений)
	// Важно для длительных операций, чтобы соединения не закрывались
	db.SetConnMaxIdleTime(0)

	if err := db.Ping(); err != nil {
		logger.Fatal(`База данных не отвечает: `, err)
	}

	logger.Success("Успешное подключение к базе данных!")

	return db, nil
}
