package repo

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"kecamatan_app/models"
)

type BantuanPemerintahRepoImpl struct {
	db *gorm.DB
}

type BantuanPemerintahInterface interface {
	CreateBantuan(data *models.BantuanPemerintah) (*models.BantuanPemerintah, error)
	GetDataByID(id int) (*models.BantuanPemerintah, error)
	GetAllData() (*models.ResponseGetAllBantuan, error)
	UpdateData(id int, data *models.BantuanPemerintah) (*models.BantuanPemerintah, error)
	CheckData(namaBantuan string) bool
}

// CreateBantuanRepoImpl ...
func CreateBantuanRepoImpl(db *gorm.DB) BantuanPemerintahInterface {
	return &BantuanPemerintahRepoImpl{db}
}

func (b *BantuanPemerintahRepoImpl) CreateBantuan(data *models.BantuanPemerintah) (*models.BantuanPemerintah, error) {
	tx := b.db.Begin()
	if err := tx.Debug().Create(&data).Error; err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[RegisterAdminRepoImpl.Insert] Error when query save data with : %w", err)
	}
	tx.Commit()

	return data, nil
}

// GetDataByID ...
func (b *BantuanPemerintahRepoImpl) GetDataByID(id int) (*models.BantuanPemerintah, error) {
	data := models.BantuanPemerintah{}
	err := b.db.Debug().Where("id = ? and is_delete=?", id, 1).Find(&data).Error
	if err != nil {
		return nil, fmt.Errorf("[StudentRepo.GetTabunganByID] Error when query get by id with error: %w", err)
	}

	return &data, nil
}

func (b *BantuanPemerintahRepoImpl) GetAllData() (*models.ResponseGetAllBantuan, error) {
	tx := b.db.Begin()

	var result models.ResponseGetAllBantuan
	var data models.BantuanPemerintah
	var dataCurrent []models.BantuanPemerintah

	rows, err := b.db.Debug().Raw("SELECT * from tb_bantuan_pemerintah;").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&data.ID, &data.CreatedAt, &data.UpdatedAt, &data.DeletedAt, &data.NamaBantuan, &data.IsDelete)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("[StudentRepo.GetByID] Error when query GetByID data with error: %w", err)
		}
		dataCurrent = append(dataCurrent, data)
	}

	result.Data = dataCurrent

	tx.Commit()

	return &result, nil
}

// UpdateData ...
func (b *BantuanPemerintahRepoImpl) UpdateData(id int, data *models.BantuanPemerintah) (*models.BantuanPemerintah, error) {
	tx := b.db.Begin()

	err := b.db.Debug().Model(&data).Where("id = ? AND is_delete=?", id, 1).Update(data).Error
	if err != nil {
		tx.Rollback()
		return nil, fmt.Errorf("[StudentRepo.Update] Error when query update data with error: %w", err)
	}

	tx.Commit()
	return data, nil
}

// CheckData ...
func (b *BantuanPemerintahRepoImpl) CheckData(namaBantuan string) bool {
	var total int

	b.db.Debug().Table("tb_bantuan_pemerintah").Where("nama_bantuan = ? and is_delete = ?", namaBantuan, 1).Count(&total)
	if total > 0 {
		return true
	}
	return false
}
