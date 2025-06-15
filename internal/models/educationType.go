package models

// EducationType представляет тип образовательной организации
type EducationType struct {
	Key            string   `json:"key" db:"key"`                         // Уникальный ключ типа (например, "kindergarten")
	Title          string   `json:"title" db:"title"`                     // Название типа (например, "Детский сад")
	Level          string   `json:"level" db:"level"`                     // Уровень образования (например, "Дошкольное")
	OwnershipForms []string `json:"ownership_forms" db:"ownership_forms"` // Формы собственности
	Keywords       []string `json:"keywords" db:"keywords"`               // Ключевые слова для поиска
}

// EducationLevel представляет уровень образования
type EducationLevel string

const (
	EducationLevelPreschool    EducationLevel = "Дошкольное"
	EducationLevelGeneral      EducationLevel = "Общее"
	EducationLevelSpecial      EducationLevel = "Специальное"
	EducationLevelProfessional EducationLevel = "Среднее профессиональное"
	EducationLevelHigher       EducationLevel = "Высшее"
)

// EducationTypeKey — тип для ключа типа образовательного учреждения.
type EducationTypeKey string

const (
	EducationTypeKindergarten       EducationTypeKey = "kindergarten"        // Детский сад
	EducationTypeSchool             EducationTypeKey = "school"              // Школа
	EducationTypeGymnasium          EducationTypeKey = "gymnasium"           // Гимназия
	EducationTypeLyceum             EducationTypeKey = "lyceum"              // Лицей
	EducationTypeBoardingSchool     EducationTypeKey = "boarding_school"     // Школа-интернат
	EducationTypeCorrectionalSchool EducationTypeKey = "correctional_school" // Коррекционная школа
	EducationTypeCollege            EducationTypeKey = "college"             // Колледж
	EducationTypeTechnicalSchool    EducationTypeKey = "technical_school"    // Техникум
	EducationTypeUniversity         EducationTypeKey = "university"          // Университет
	EducationTypeInstitute          EducationTypeKey = "institute"           // Институт
	EducationTypeAcademy            EducationTypeKey = "academy"             // Академия
)

type EducationTypeShortInfo struct {
	Key      string   `json:"key" db:"key"`
	Keywords []string `json:"keywords" db:"keywords" type:"text[]"`
}
