package models

import "encoding/xml"

// Описание полей можно посмотреть в scructure-example.xml

type OpenData struct {
	XMLName      xml.Name      `xml:"OpenData"`
	Certificates []Certificate `xml:"Certificates>Certificate"`
}

type Certificate struct {
	XMLName                          xml.Name              `xml:"Certificate"`
	ID                               string                `xml:"Id"`
	IsFederal                        string                `xml:"IsFederal"` // 1 - федеральное, 0 - региональное
	StatusName                       string                `xml:"StatusName"`
	TypeName                         string                `xml:"TypeName"`
	RegionName                       string                `xml:"RegionName"`
	RegionCode                       string                `xml:"RegionCode"`
	FederalDistrictName              string                `xml:"FederalDistrictName"`
	FederalDistrictShortName         string                `xml:"FederalDistrictShortName"`
	RegNumber                        string                `xml:"RegNumber"`
	SerialNumber                     string                `xml:"SerialNumber"`
	FormNumber                       string                `xml:"FormNumber"`
	IssueDate                        string                `xml:"IssueDate"`
	EndDate                          string                `xml:"EndDate"`
	ControlOrgan                     string                `xml:"ControlOrgan"`
	PostAddress                      string                `xml:"PostAddress"`
	EduOrgFullName                   string                `xml:"EduOrgFullName"`
	EduOrgShortName                  string                `xml:"EduOrgShortName"`
	EduOrgINN                        string                `xml:"EduOrgINN"`
	EduOrgKPP                        string                `xml:"EduOrgKPP"`
	EduOrgOGRN                       string                `xml:"EduOrgOGRN"`
	IndividualEntrepreneurLastName   string                `xml:"IndividualEntrepreneurLastName"`
	IndividualEntrepreneurFirstName  string                `xml:"IndividualEntrepreneurFirstName"`
	IndividualEntrepreneurMiddleName string                `xml:"IndividualEntrepreneurMiddleName"`
	IndividualEntrepreneurAddress    string                `xml:"IndividualEntrepreneurAddress"`
	IndividualEntrepreneurEGRIP      string                `xml:"IndividualEntrepreneurEGRIP"`
	IndividualEntrepreneurINN        string                `xml:"IndividualEntrepreneurINN"`
	ActualEducationOrganization      EducationOrganization `xml:"ActualEducationOrganization"`
	Supplements                      []Supplement          `xml:"Supplements>Supplement"`
	Decisions                        []Decision            `xml:"Decisions>Decision"`
}

type Supplement struct {
	XMLName                     xml.Name              `xml:"Supplement"`
	ID                          string                `xml:"Id"`
	StatusName                  string                `xml:"StatusName"`
	StatusCode                  string                `xml:"StatusCode"`
	Number                      string                `xml:"Number"`
	SerialNumber                string                `xml:"SerialNumber"`
	FormNumber                  string                `xml:"FormNumber"`
	IssueDate                   string                `xml:"IssueDate"`
	IsForBranch                 string                `xml:"IsForBranch"` // 1 - филиалу, 0 - головной
	Note                        string                `xml:"Note"`
	EduOrgFullName              string                `xml:"EduOrgFullName"`
	EduOrgShortName             string                `xml:"EduOrgShortName"`
	EduOrgAddress               string                `xml:"EduOrgAddress"`
	EduOrgKPP                   string                `xml:"EduOrgKPP"`
	ActualEducationOrganization EducationOrganization `xml:"ActualEducationOrganization"`
	EducationalPrograms         []EducationalProgram  `xml:"EducationalPrograms>EducationalProgram"`
}

type EducationalProgram struct {
	XMLName            xml.Name `xml:"EducationalProgram"`
	ID                 string   `xml:"Id"`
	TypeName           string   `xml:"TypeName"`
	EduLevelName       string   `xml:"EduLevelName"`
	ProgrammName       string   `xml:"ProgrammName"`
	ProgrammCode       string   `xml:"ProgrammCode"`
	UGSName            string   `xml:"UGSName"`
	UGSCode            string   `xml:"UGSCode"`
	EduNormativePeriod string   `xml:"EduNormativePeriod"`
	Qualification      string   `xml:"Qualification"`
	IsAccredited       string   `xml:"IsAccredited"` // 0 - аккредитована, 1 - отказ
	IsCanceled         string   `xml:"IsCanceled"`
	IsSuspended        string   `xml:"IsSuspended"`
}

type Decision struct {
	XMLName             xml.Name `xml:"Decision"`
	ID                  string   `xml:"Id"`
	DecisionTypeName    string   `xml:"DecisionTypeName"`
	OrderDocumentNumber string   `xml:"OrderDocumentNumber"`
	OrderDocumentKind   string   `xml:"OrderDocumentKind"`
	DecisionDate        string   `xml:"DecisionDate"`
}
