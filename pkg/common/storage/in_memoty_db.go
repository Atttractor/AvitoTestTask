package db

import (
	"dooplik/avitoTestTask/pkg/common/models"
	"errors"
	"time"
)

type LocalDB struct {
	db map[uint]models.Banner
}

func InitInMemoryDB(d Storage) (LocalDB, error) {
	m := make(map[uint]models.Banner)

	banners, err := d.GetBannersFiltred(-1, -1, -1, -1)

	if err != nil {
		return LocalDB{}, err
	}

	for _, banner := range(banners) {
		m[banner.ID] = banner
	}

	return LocalDB{db: m}, nil
}

func (l *LocalDB) UpdateInMemotyDB(d Storage) error {
	for {
		banners, err := d.GetBannersFiltred(-1, -1, -1, -1)
		if err != nil {
			return err
		}

		for _, banner := range(banners) {
			l.db[banner.ID] = banner
		}



		time.Sleep(5 * time.Minute)
	}
}

func (l LocalDB) GetBanner(tag_id, feature_id uint) (models.Banner, error) {
	for _, banner := range(l.db) {
		if banner.FeatureId == feature_id {
			for _, tag := range(banner.Tags) {
				if tag.ID == tag_id {
					return banner, nil
				}
			}
		}
	}
	return models.Banner{}, errors.New("not found")
}