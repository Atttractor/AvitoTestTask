package models

import "gorm.io/gorm"

type Tag struct {
	ID        	uint `gorm:"primarykey"`
	gorm.Model
}

type Banner struct {
	gorm.Model
	IsActive  bool   `gorm:"not null"`
	Data      string `gorm:"not null"`
	Tags      []Tag  `gorm:"many2many:banners_tags;"`
	FeatureId uint
}

type Feature struct {
	ID        	uint `gorm:"primarykey"`
	gorm.Model
	Banners 	[]Banner
}