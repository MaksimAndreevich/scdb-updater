package main

import (
	"gitlab.com/scdb/updater/internal/config"
	"gitlab.com/scdb/updater/internal/services"
)

func main() {

	config.LoadConfig()

	// services.SeedOrganisations()
	services.SeedCities()
}
