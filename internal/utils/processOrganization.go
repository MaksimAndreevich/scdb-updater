package utils

import (
	"gitlab.com/scdb/updater/internal/models"
)

// ProcessOrganization обрабатывает одну организацию
func ProcessOrganization(org models.EducationOrganization,
	citiesMap map[string]models.CityShortInfo,
	regionsMap map[string]models.RegionShortInfo,
	noLocationCount *int,
	orgTypesMap map[string]models.EducationTypeShortInfo,
	noTypeOrganisationsCount *int) (string, int, int, string) {

	location := findLocation(org.PostAddress, org.RegionName, citiesMap, regionsMap)
	orgType := findType(org.FullName, orgTypesMap)

	if !location.Found {
		// Организация не привязана к городу или региону. Назначаем дефолтный регион 'Другое'
		*noLocationCount++
	}

	if !orgType.Found {
		// Организация не привязана к типу. Назначаем дефолтный тип 'Другое'
		*noTypeOrganisationsCount++
	}

	return location.CityID, location.RegionID, location.FederalDistrictID, orgType.Key

}
