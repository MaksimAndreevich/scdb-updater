package models

import "time"

type City struct {
	FiasID              string    `json:"fias_id" db:"fias_id"`
	Address             string    `json:"address" db:"address"`
	PostalCode          int       `json:"postal_code" db:"postal_code"`
	Country             string    `json:"country" db:"country"`
	RegionName          string    `json:"region" db:"region"`
	RegionType          string    `json:"region_type" db:"region_type"`
	AreaType            string    `json:"area_type" db:"area_type"`
	Area                string    `json:"area" db:"area"`
	CityType            string    `json:"city_type" db:"city_type"`
	CityName            string    `json:"city" db:"city"`
	KladrID             int       `json:"kladr_id" db:"kladr_id"`
	FederalDistrictName string    `json:"federal_district_name" db:"federal_district_name"`
	FiasLevel           int       `json:"fias_level" db:"fias_level"`
	CapitalMarker       int       `json:"capital_marker" db:"capital_marker"`
	OKATO               int       `json:"okato" db:"okato"`
	OKTMO               int       `json:"oktmo" db:"oktmo"`
	TaxOffice           int       `json:"tax_office" db:"tax_office"`
	Timezone            string    `json:"timezone" db:"timezone"`
	GeoLat              float64   `json:"geo_lat" db:"geo_lat"`
	GeoLon              float64   `json:"geo_lon" db:"geo_lon"`
	Population          int       `json:"population" db:"population"`
	FoundationYear      int       `json:"foundation_year" db:"foundation_year"`
	CreatedAt           time.Time `json:"created_at" db:"created_at"`
	UpdatedAt           time.Time `json:"updated_at" db:"updated_at"`

	RegionID          int `json:"region_id" db:"fk_region_id"`
	FederalDistrictID int `json:"federal_district_id" db:"fk_federal_district_id"`

	// Связи
	Organizations []EducationOrganization `json:"organizations,omitempty" db:"-"`
}

type CityShortInfo struct {
	CityName          string `json:"city" db:"city"`
	FiasID            string `json:"fias_id" db:"fias_id"`
	RegionID          int    `json:"region_id" db:"fk_region_id"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id"`
}
