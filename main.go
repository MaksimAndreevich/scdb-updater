package main

import (
	"gitlab.com/scdb/updater/internal/config"
	"gitlab.com/scdb/updater/internal/database"
	"gitlab.com/scdb/updater/internal/logger"
	"gitlab.com/scdb/updater/internal/services"
)

func main() {
	config.LoadConfig()
	database.Connect()

	districts := services.SeedFederalDistricts()
	services.SeedRegions(districts)

	services.SeedCities()
	services.SeedOrganisations()

	// Закрываем соединение с базой данных
	defer func() {
		if err := database.DB.Close(); err != nil {
			logger.Error("Ошибка при закрытии базы данных:", err)
		}
	}()

}
