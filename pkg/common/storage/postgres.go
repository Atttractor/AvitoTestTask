package db

import (
	"dooplik/avitoTestTask/pkg/common/config"
	"dooplik/avitoTestTask/pkg/common/models"
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Storage struct {
	DB *gorm.DB
}

func InitPostgresDB(c *config.Config) Storage{
	url := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", c.DBHost, c.DBUser, c.DBPass, c.DBName, c.DBPort)
	
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Feature{})
	db.AutoMigrate(&models.Tag{})
	db.AutoMigrate(&models.Banner{})

	return Storage {
		DB: db,
	}
}

func (d *Storage) GetBanner(feature_id, tag_id uint) (models.Banner, error) {
	db := d.DB
	banners := []models.Banner{}

	r := db.Where("feature_id = ? AND is_active = true", feature_id).Preload("Tags").Find(&banners)
	if r.Error != nil {
		return models.Banner{}, r.Error
	}

	for _, banner := range(banners) {
		for _, tag := range(banner.Tags) {
			if tag.ID == tag_id {
				return banner, nil
			}
		}
	}

	return models.Banner{}, errors.New("banner not found")
}

func (d Storage) GetBannersFiltred(feature_id, tag_id, limit, offset int) ([]models.Banner, error) {
	db := d.DB
	banners := []models.Banner{}

	if feature_id >= 0 {
		r := db.Limit(limit).Offset(offset).Where("feature_id = ?", feature_id).Preload("Tags").Find(&banners)
		if r.Error != nil {
			return banners, r.Error
		}
	} else {
		r := db.Limit(limit).Offset(offset).Preload("Tags").Find(&banners)
		if r.Error != nil {
			return banners, r.Error
		}
	}

	if tag_id >= 0 {
		res_banners := []models.Banner{}

		for _, banner := range(banners) {
			for _, tag := range(banner.Tags) {
				if int(tag.ID) == tag_id {
					res_banners = append(res_banners, banner)
				}
			}
		}

		return res_banners, nil
	}

	return banners, nil
}

func (d Storage) CreateBanner(tags []uint, feature_id uint, data string, is_active bool) (int, error) {
	db := d.DB
	t := []models.Tag{}
	feature := models.Feature{}

	r := db.Find(&feature, feature_id)
	
	if r.Error != nil || feature.ID != uint(feature_id) {
		return -1, r.Error
	}

	for _, item := range(tags) {
		t = append(t, models.Tag{ID:item})
	}

	banner := models.Banner{
		FeatureId: uint(feature_id),
		IsActive: is_active,
		Data: data,
		Tags: t,
	}


	r = db.Create(&banner)

	if r.Error != nil {
		return -1, r.Error
	}

	return int(banner.ID), nil
}

func (d Storage) UpdateById(id uint, tags []uint, feature_id uint, data string, is_active bool) error {
	db := d.DB
	feature := models.Feature{}
	banner := models.Banner{}

	r := db.Find(&banner, id)
	if r.Error != nil || banner.ID != id {
		return r.Error
	}

	r = db.Find(&feature, feature_id)
	if r.Error != nil || feature.ID != uint(feature_id) {
		return r.Error
	}

	return nil
}

func (d Storage) DeleteBannerById(id uint) error {
	db := d.DB
	
	r := db.Delete(&models.Banner{}, id)

	if r.Error != nil {
		return r.Error
	}

	return nil
}

func (d Storage) InsertTagsAndFeatures() error {
	db := d.DB

	features := []*models.Feature{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}

	tags := []*models.Tag{
		{ID: 1},
		{ID: 2},
		{ID: 3},
	}


	r := db.Create(features)
	if r.Error != nil {
		return r.Error
	}

	r = db.Create(tags)
	if r.Error != nil {
		return r.Error
	}

	return nil
}