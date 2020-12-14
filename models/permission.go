package models

import "time"

type Permission struct {
	ID             uint `gorm:"primaryKey"`
	UserID         uint
	User           User
	OrganizationID uint
	Organization   Organization
	DeactivatedAt  time.Time
}
