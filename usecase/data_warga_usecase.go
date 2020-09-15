package usecase

import (
	"encoding/json"
	"fmt"
	"kecamatan_app/models"
	"kecamatan_app/repo"
	"log"
)

type DataWargaUsecaseImpl struct {
	dataWargaUsecase repo.DataWargaInterface
	daerahIndoUsecase repo.DaerahRepoInterface
}

type DataWargaUsecaseInterface interface {
	CreateDataWarga(data *models.DataWarga) (*models.DataWarga, error)
}

// CreateBantuanUsecaseImpl ...
func CreateDataWargaUsecaseImpl(dataWargaUsecase repo.DataWargaInterface, daerahIndoUsecase repo.DaerahRepoInterface) DataWargaUsecaseInterface {
	return &DataWargaUsecaseImpl{dataWargaUsecase, daerahIndoUsecase}
}

func (d *DataWargaUsecaseImpl) CreateDataWarga(data *models.DataWarga) (*models.DataWarga, error) {
	provinsi, _ := d.daerahIndoUsecase.GetDataProvinsiByID(data.ProvinsiID)

	dataKotaKab, errs := d.daerahIndoUsecase.GetDataKotaKab(data.ProvinsiID)
	if errs != nil {
		return nil, errs
	}

	dataKec, _ := d.daerahIndoUsecase.GetDataKecamatan(data.KotaKabID)
	dataKel, _ := d.daerahIndoUsecase.GetDataKelurahan(data.KecamatanID)

	var tost models.DetailAlamat
	tost.Alamat = data.Alamat
	tost.RT = data.RT
	tost.RW = data.RW
	tost.Provinsi = provinsi.Nama

	length := len(dataKotaKab.KotaKabupaten)
	lengthKec := len(dataKec.Kecamatan)
	lengthKel := len(dataKel.Kelurahan)

	for i := 0; i < length ; i++  {
		if dataKotaKab.KotaKabupaten[i].ID == data.KotaKabID {
			tost.KotaKabupaten = dataKotaKab.KotaKabupaten[i].Nama
		}
	}

	for i := 0; i < lengthKec ; i++  {
		if dataKec.Kecamatan[i].ID == data.KecamatanID {
			tost.Kecamatan = dataKec.Kecamatan[i].Nama
		}
	}

	for i := 0; i < lengthKel ; i++  {
		if dataKel.Kelurahan[i].ID == data.KelurahanID {
			tost.Kelurahan = dataKel.Kelurahan[i].Nama
		}
	}

	e, err := json.Marshal(tost)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	data.DetailAlamat = string(e)
	result, err := d.dataWargaUsecase.CreateDataWarga(data)
	if err != nil {
		log.Println("ERROR WHEN GET DATA FROM REPO DATA WARGA")
		return nil, err
	}

	return result, nil
}
