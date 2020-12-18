package models

type User struct {
	ID          uint `gorm:"primaryKey"`
	Email       string
	Username    string
	Permissions []Permission
}

func (u *User) UseSso() bool {
	var count int64

	DB.Debug().Model(&Organization{}).Joins("Permissions").Joins("SamlSetting").
		Where("Permissions.user_id = ? AND Permissions.deactivated_at IS NULL", u.ID).
		Where("SamlSetting.active = 1").Count(&count)

	return count != 0
}
