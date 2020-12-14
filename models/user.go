package models

type User struct {
	ID          uint `gorm:"primaryKey"`
	Email       string
	Username    string
	Permissions []Permission
}

func (u *User) UseSso() bool {
	var count int64

	DB.Model(&Organization{}).Joins("Permissions").Joins("SamlSetting").
		Where("permissions.user_id = ? AND permissions.deactivated_at IS NULL").
		Where("saml_settings.active = 1", u.ID).Count(&count)

	return count == 0
}
