package services

import (
	"encoding/xml"
	"io"
	"os"

	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/models"
)

func GetDataParsedXML() models.OpenData {
	filePath := GetXMLFilePath()

	xmlFile, err := os.Open(filePath)

	if err != nil {
		logger.Fatal("Ошибка при получении файла по пути: ", err)
	}

	byteValue, err := io.ReadAll(xmlFile)

	if err != nil {
		logger.Fatal("Ошибка при разбитии файла на байты ", err)
	}

	defer func() {
		if err := xmlFile.Close(); err != nil {
			logger.Error("Ошибка при закрытии XML файла:", err)
		}
	}()

	var data models.OpenData

	err = xml.Unmarshal(byteValue, &data)
	if err != nil {
		logger.Fatal("Ошибка при парсинге XML:", err)
	}

	logger.Info("Найдено сертификатов ", len(data.Certificates))

	return data
}
