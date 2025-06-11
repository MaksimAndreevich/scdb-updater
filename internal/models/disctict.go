package models

import "time"

type FederalDistrict struct {
	ID        int       `json:"id" db:"id"`
	ShortName string    `json:"short_name" db:"short_name"` // ЦФО, СЗФО и т.д.
	Name      string    `json:"name" db:"name"`             // Полное название
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`

	Regions []Region `json:"regions,omitempty" db:"-"` // Связь с регионами
}
