package utils

import (
	"strings"

	"gitlab.com/scdb/updater/internal/models"
)

type OrgTypeMatch struct {
	Key   string
	Found bool
}

func findType(fullName string, orgTypesMap map[string]models.EducationTypeShortInfo) OrgTypeMatch {
	fullName = strings.ToLower(fullName)

	for key, value := range orgTypesMap {
		for _, keyword := range value.Keywords {
			if strings.Contains(fullName, strings.ToLower(keyword)) {
				return OrgTypeMatch{
					Key:   key,
					Found: true,
				}
			}
		}
	}

	return OrgTypeMatch{
		Key:   "other",
		Found: false,
	}

}
