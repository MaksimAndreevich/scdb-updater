package models

import "time"

type EducationOrganization struct {
	ID           string    `json:"id" db:"id"`
	FullName     string    `json:"full_name" db:"full_name"`
	ShortName    string    `json:"short_name" db:"short_name"`
	HeadEduOrgID string    `json:"head_edu_org_id" db:"head_edu_org_id"`
	IsBranch     bool      `json:"is_branch" db:"is_branch"`
	PostAddress  string    `json:"post_address" db:"post_address"`
	Phone        string    `json:"phone" db:"phone"`
	Fax          string    `json:"fax" db:"fax"`
	Email        string    `json:"email" db:"email"`
	WebSite      string    `json:"web_site" db:"web_site"`
	OGRN         string    `json:"ogrn" db:"ogrn"`
	INN          string    `json:"inn" db:"inn"`
	KPP          string    `json:"kpp" db:"kpp"`
	HeadPost     string    `json:"head_post" db:"head_post"`
	HeadName     string    `json:"head_name" db:"head_name"`
	FormName     string    `json:"form_name" db:"form_name"`
	KindName     string    `json:"kind_name" db:"kind_name"`
	TypeName     string    `json:"type_name" db:"type_name"`
	CreatedAt    time.Time `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time `json:"updated_at" db:"updated_at"`

	// Связи
	CityID            string `json:"city_id" db:"fk_city_id"`
	RegionID          int    `json:"region_id" db:"fk_region_id"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id"`
}
