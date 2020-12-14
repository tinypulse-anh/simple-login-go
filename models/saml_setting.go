package models

type SamlSetting struct {
	ID             uint `gorm:"primaryKey"`
	OrganizationID uint
	Organization
	Active bool
}
