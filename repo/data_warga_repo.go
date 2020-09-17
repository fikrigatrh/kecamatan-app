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
	GetAllData() (*models.ResponseGetAllDataWarga, error)
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

func (b *DataWargaRepoImpl) GetAllData() (*models.ResponseGetAllDataWarga, error) {
	tx := b.db.Begin()

	var dataWarga models.DataWarga
	var result models.ResponseGetAllDataWarga
	var dataCurrent []models.DataWarga

	rows, err := b.db.Debug().Raw("SELECT * from tb_data_warga;").Rows()
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&dataWarga.ID, &dataWarga.CreatedAt, &dataWarga.UpdatedAt, &dataWarga.DeletedAt, &dataWarga.Nik, &dataWarga.NoKk, &dataWarga.SudahMemilikiKTP, &dataWarga.NamaLengkap, &dataWarga.JenisKelamin, &dataWarga.Pendidikan, &dataWarga.TempatLahir, &dataWarga.TanggalLahir,
			&dataWarga.Agama, &dataWarga.GolDarah, &dataWarga.DetailAlamat, &dataWarga.StatusPerkawinan, &dataWarga.NoBukuNikah, &dataWarga.JenisPekerjaan,&dataWarga.SHOK, &dataWarga.Kewarganegaraan, &dataWarga.Disabilitas, &dataWarga.NoPaspor, &dataWarga.NoKitasKitap, &dataWarga.NamaAyah,
			&dataWarga.NamaIbu, &dataWarga.AlamatSesuaiKK, &dataWarga.StatusTempatTinggal, &dataWarga.NoShmHcb, &dataWarga.NoAjb, &dataWarga.NoSpptPbb, &dataWarga.BantuanPemerintahID)
		if err != nil {
			tx.Rollback()
			return nil, fmt.Errorf("[StudentRepo.GetByID] Error when query GetByID data with error: %w", err)
		}
		dataCurrent = append(dataCurrent, dataWarga)
	}

	result.Data = dataCurrent

	tx.Commit()

	return &result, nil
}
