package internal

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/vinay03/chalk"
	logger "gitlab.com/scdb/core/logger"
)

func GetXMLFilePath() string {
	var filePath string

	for {
		fmt.Println(chalk.Black().BgBlue("Введите путь к файлу XML: "))
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
