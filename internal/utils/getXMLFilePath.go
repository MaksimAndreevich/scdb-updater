package utils

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vinay03/chalk"
	"scdb-updater/internal/logger"
)

func GetXMLFilePath() string {
	var filePath string

	for {
		logger.Info("Введите путь к файлу XML: ")
		n, err := fmt.Scanf("%s", &filePath)

		if err != nil || n != 1 {
			logger.Warning("Ошибка ввода, попробуйте ещё раз")
			continue
		}

		if isXMLFile(filePath) {
			break
		}

		fmt.Println(chalk.Black().BgYellow("Ошибка: файл должен иметь расширение .xml"))
	}

	return filePath

}

func isXMLFile(path string) bool {
	return strings.ToLower(filepath.Ext(path)) == ".xml"
}
