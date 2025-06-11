package models

import "time"

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
	NamecaseNominative    string `json:"namecase.nominative" db:"namecase_nominative"`
	NamecaseGenitive      string `json:"namecase.genitive" db:"namecase_genitive"`
	NamecaseDative        string `json:"namecase.dative" db:"namecase_dative"`
	NamecaseAccusative    string `json:"namecase.accusative" db:"namecase_accusative"`
	NamecaseAblative      string `json:"namecase.ablative" db:"namecase_ablative"`
	NamecasePrepositional string `json:"namecase.prepositional" db:"namecase_prepositional"`
	NamecaseLocative      string `json:"namecase.locative" db:"namecase_locative"`

	// Столица региона
	CapitalName        string `json:"capital.name" db:"capital_name"`
	CapitalLabel       string `json:"capital.label" db:"capital_label"`
	CapitalKladrID     string `json:"capital.kladr_id" db:"capital_kladr_id"`
	CapitalOKATO       string `json:"capital.okato" db:"capital_okato"`
	CapitalOKTMO       string `json:"capital.oktmo" db:"capital_oktmo"`
	CapitalContentType string `json:"capital.contentType" db:"capital_content_type"`

	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

type RegionShortInfo struct {
	ID                int    `json:"-" db:"id"`
	Name              string `json:"name" db:"name"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id"`
}
