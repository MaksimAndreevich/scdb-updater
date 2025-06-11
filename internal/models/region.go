package models

import "time"

type Namecase struct {
	Nominative    string `json:"nominative" db:"namecase_nominative"`
	Genitive      string `json:"genitive" db:"namecase_genitive"`
	Dative        string `json:"dative" db:"namecase_dative"`
	Accusative    string `json:"accusative" db:"namecase_accusative"`
	Ablative      string `json:"ablative" db:"namecase_ablative"`
	Prepositional string `json:"prepositional" db:"namecase_prepositional"`
	Locative      string `json:"locative" db:"namecase_locative"`
}

type Capital struct {
	Name        string `json:"name"`
	Label       string `json:"label"`
	KladrID     string `json:"kladr_id"`
	OKATO       string `json:"okato"`
	OKTMO       string `json:"oktmo"`
	ContentType string `json:"contentType"`
}

type Region struct {
	ID                  int     `json:"-" db:"id"`
	Name                string  `json:"name" db:"name"`
	Label               string  `json:"label" db:"label"`
	Type                string  `json:"type" db:"type"`
	TypeShort           string  `json:"typeShort" db:"type_short"`
	ContentType         string  `json:"contentType" db:"content_type"`
	OKATO               string  `json:"okato" db:"okato"`
	OKTMO               string  `json:"oktmo" db:"oktmo"`
	GUID                string  `json:"guid" db:"guid"`
	Code                string  `json:"code" db:"code"`
	ISO3166_2           string  `json:"iso_3166-2" db:"iso_3166_2"`
	Population          int     `json:"population" db:"population"`
	YearFounded         int     `json:"yearFounded" db:"year_founded"`
	Area                float64 `json:"area" db:"area"`
	Fullname            string  `json:"fullname" db:"fullname"`
	NameEn              string  `json:"name_en" db:"name_en"`
	FederalDistrictName string  `json:"district" db:"district_name"`

	RegionKladrID     string `json:"kladr_id" db:"kladr_id"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id"`
	Cities            []City `json:"cities,omitempty" db:"-"`

	// Падежи названия
	Namecase Namecase `json:"namecase"`

	// Столица региона
	Capital Capital `json:"capital"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegionShortInfo struct {
	ID                int    `json:"id" db:"id"`
	Name              string `json:"name" db:"name"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id"`
}
