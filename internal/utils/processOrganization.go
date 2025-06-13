package utils

import (
	"gitlab.com/scdb/updater/internal/models"
)

// LocationMatch представляет результат поиска местоположения
type LocationMatch struct {
	CityID            string
	RegionID          int
	FederalDistrictID int
	Found             bool
}

// findLocation ищет местоположение по адресу и названию региона
func findLocation(address string, regionName string, citiesMap map[string]models.CityShortInfo, regionsMap map[string]models.RegionShortInfo) LocationMatch {

	// Сначала ищем город
	if city, found := FindInText(address, citiesMap); found {
		return LocationMatch{
			CityID:            city.FiasID,
			RegionID:          city.RegionID,
			FederalDistrictID: city.FederalDistrictID,
			Found:             true,
		}
	}

	// Если город не нашелся - ищем регион
	if region, found := FindInText(regionName, regionsMap); found {
		return LocationMatch{
			RegionID:          region.ID,
			FederalDistrictID: region.FederalDistrictID,
			Found:             true,
		}
	}

	// Если не нашелся город и регион, используем регион 'Другое'
	otherRegion := regionsMap["Другое"]
	return LocationMatch{
		RegionID:          otherRegion.ID,
		FederalDistrictID: otherRegion.FederalDistrictID,
		Found:             false,
	}
}

// ProcessOrganization обрабатывает одну организацию
func ProcessOrganization(org models.EducationOrganization, citiesMap map[string]models.CityShortInfo, regionsMap map[string]models.RegionShortInfo, noLocationCount *int) (string, int, int, string) {

	location := findLocation(org.PostAddress, org.RegionName, citiesMap, regionsMap)

	if !location.Found {
		// Организация не привязана к городу или региону. Назначаем дефолтный регион 'Другое'
		*noLocationCount++
	}

	// TODO: определять тип организации
	orgType := "school"

	return location.CityID, location.RegionID, location.FederalDistrictID, orgType

}
