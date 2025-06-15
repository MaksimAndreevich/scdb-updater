package models

import "time"

type EducationOrganization struct {
	ID                       string    `json:"id" db:"id" xml:"Id"`
	FullName                 string    `json:"full_name" db:"full_name" xml:"FullName"`
	ShortName                string    `json:"short_name" db:"short_name" xml:"ShortName"`
	HeadEduOrgID             string    `json:"head_edu_org_id" db:"head_edu_org_id" xml:"HeadEduOrgId"`
	IsBranch                 bool      `json:"is_branch" db:"is_branch" xml:"IsBranch"`
	PostAddress              string    `json:"post_address" db:"post_address" xml:"PostAddress"`
	Phone                    string    `json:"phone" db:"phone" xml:"Phone"`
	Fax                      string    `json:"fax" db:"fax" xml:"Fax"`
	Email                    string    `json:"email" db:"email" xml:"Email"`
	WebSite                  string    `json:"web_site" db:"web_site" xml:"WebSite"`
	OGRN                     string    `json:"ogrn" db:"ogrn" xml:"OGRN"`
	INN                      string    `json:"inn" db:"inn" xml:"INN"`
	KPP                      string    `json:"kpp" db:"kpp" xml:"KPP"`
	HeadPost                 string    `json:"head_post" db:"head_post" xml:"HeadPost"`
	HeadName                 string    `json:"head_name" db:"head_name" xml:"HeadName"`
	FormName                 string    `json:"form_name" db:"form_name" xml:"FormName"`
	KindName                 string    `json:"kind_name" db:"kind_name" xml:"KindName"`
	TypeName                 string    `json:"type_name" db:"type_name" xml:"TypeName"`
	RegionName               string    `json:"region_name" db:"-" xml:"RegionName"`
	FederalDistrictShortName string    `json:"federal_district_short_name" db:"-" xml:"FederalDistrictShortName"`
	FederalDistrictName      string    `json:"federal_district_name" db:"-" xml:"FederalDistrictName"`
	CreatedAt                time.Time `json:"created_at" db:"created_at" xml:"-"`
	UpdatedAt                time.Time `json:"updated_at" db:"updated_at" xml:"-"`

	// Связи
	CityID            string `json:"city_id" db:"fk_city_id" xml:"-"`
	RegionID          int    `json:"region_id" db:"fk_region_id" xml:"-"`
	FederalDistrictID int    `json:"federal_district_id" db:"fk_federal_district_id" xml:"-"`
	EducationTypeKey  string `json:"education_type_key" db:"fk_education_type_key" xml:"-"`
}
