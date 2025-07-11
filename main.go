package main

import (
	"scdb-updater/internal/config"
	"scdb-updater/internal/database"
	"scdb-updater/internal/logger"
	"scdb-updater/internal/services"
)

func main() {
	config.LoadConfig()
	database.Connect()

	districts := services.SeedFederalDistricts()
	services.SeedRegions(districts)
	services.SeedCities()
	services.SeedOrganisationsTypes()
	services.SeedOrganisations()

	// Закрываем соединение с базой данных
	defer func() {
		if err := database.DB.Close(); err != nil {
			logger.Error("Ошибка при закрытии базы данных:", err)
		}
	}()

}
