package models

// City представляет информацию о городе России
type City struct {
	Address         string  `json:"address"`          // Адрес одной строкой
	PostalCode      int     `json:"postal_code"`      // Почтовый индекс
	Country         string  `json:"country"`          // Страна
	FederalDistrict string  `json:"federal_district"` // Федеральный округ
	RegionType      string  `json:"region_type"`      // Тип региона
	Region          string  `json:"region"`           // Регион
	AreaType        string  `json:"area_type"`        // Тип района
	Area            string  `json:"area"`             // Район
	CityType        string  `json:"city_type"`        // Тип города
	City            string  `json:"city"`             // Город
	SettlementType  string  `json:"settlement_type"`  // Тип населенного пункта
	Settlement      string  `json:"settlement"`       // Населенный пункт
	KladrID         int64   `json:"kladr_id"`         // КЛАДР-код
	FiasID          string  `json:"fias_id"`          // ФИАС-код
	FiasLevel       int     `json:"fias_level"`       // Уровень по ФИАС
	CapitalMarker   int     `json:"capital_marker"`   // Признак центра региона или района
	Okato           int64   `json:"okato"`            // Код ОКАТО
	Oktmo           int64   `json:"oktmo"`            // Код ОКТМО
	TaxOffice       int     `json:"tax_office"`       // Код ИФНС
	Timezone        string  `json:"timezone"`         // Часовой пояс
	GeoLat          float64 `json:"geo_lat"`          // Широта
	GeoLon          float64 `json:"geo_lon"`          // Долгота
	Population      int     `json:"population"`       // Население
	FoundationYear  int     `json:"foundation_year"`  // Год основания
}
