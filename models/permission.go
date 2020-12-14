package models

import "time"

type Permission struct {
	ID     uint `gorm:"primaryKey"`
	UserID uint
	User
	OrganizationID uint
	Organization
	DeactivatedAt time.Time
}
