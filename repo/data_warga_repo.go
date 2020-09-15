package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"kecamatan_app/models"
)

type DataWargaRepoImpl struct {
	db *gorm.DB
}

type DataWargaInterface interface {
	CreateDataWarga(data *models.DataWarga) (*models.DataWarga, error)
}

// CreateBantuanRepoImpl ...
func CreateDataWargaRepoImpl(db *gorm.DB) DataWargaInterface {
	return &DataWargaRepoImpl{db}
}

func (b *DataWargaRepoImpl) CreateDataWarga(data *models.DataWarga) (*models.DataWarga, error) {
	tx := b.db.Begin()
	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
	}
	tx.Commit()

	return data, nil
}
