package utils

import (
	"os"

	"gitlab.com/scdb/updater/internal/logger"
)

// TODO: добавить проверку на sql
func LoadSQLFile(filePath string) string {
	content, err := os.ReadFile(filePath)
	if err != nil {
		logger.Fatal("Ошибка чтения SQL файла: ", filePath, err)
	}

	return string(content)
}
