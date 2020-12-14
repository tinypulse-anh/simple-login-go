package models

type Organization struct {
	ID          uint `gorm:"primaryKey"`
	Permissions []Permission
	SamlSetting SamlSetting
}
