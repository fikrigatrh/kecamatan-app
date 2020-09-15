package usecase

import (
	"kecamatan_app/models"
	"kecamatan_app/repo"
	"log"
)

type DataWargaUsecaseImpl struct {
	dataWargaUsecase repo.DataWargaInterface
	daerahIndoUsecase repo.DaerahRepoInterface
}

type DataWargaUsecaseInterface interface {
	CreateDataWarga(data *models.DataWargaExample) (*models.DataWargaExample, error)
}

// CreateBantuanUsecaseImpl ...
func CreateDataWargaUsecaseImpl(dataWargaUsecase repo.DataWargaInterface, daerahIndoUsecase repo.DaerahRepoInterface) DataWargaUsecaseInterface {
	return &DataWargaUsecaseImpl{dataWargaUsecase, daerahIndoUsecase}
}

func (d *DataWargaUsecaseImpl) CreateDataWarga(data *models.DataWargaExample) (*models.DataWargaExample, error) {
	dataKotaKab, errs := d.daerahIndoUsecase.GetDataKotaKab(data.ProvinsiID)
	if errs != nil {
		return nil, errs
	}
	var test models.DataWargaExample

	length := len(dataKotaKab.KotaKabupaten)

	for i := 0; i < length ; i++  {
		if dataKotaKab.KotaKabupaten[i].ID == data.KotaKabID {
			test.NamaKotaKab = dataKotaKab.KotaKabupaten[i].Nama
		}
	}

	data.NamaKotaKab = test.NamaKotaKab
	result, err := d.dataWargaUsecase.CreateDataWarga(data)
	if err != nil {
		log.Println("ERROR WHEN GET DATA FROM REPO DATA WARGA")
		return nil, err
	}
	result.NamaKotaKab = test.NamaKotaKab

	return result, nil
}
