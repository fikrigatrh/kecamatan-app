package repo

import (
	"encoding/json"
	"fmt"
	"gopkg.in/resty.v1"
	"kecamatan_app/models"
)

type DaerahRepo struct {
}

type DaerahRepoInterface interface {
	GetDataProvinsiByID(id int) (*models.ProvinsiStruct, error)
	GetDataKotaKab(id int) (*models.KotaKabStruct, error)
	GetDataKecamatan(id int) (*models.KecStruct, error)
	GetDataKelurahan(id int) (*models.KelStruct, error)
}

func CreateDaerahIndoRepo() DaerahRepoInterface {
	return &DaerahRepo{}
}

func (s *DaerahRepo) GetDataProvinsiByID(id int) (*models.ProvinsiStruct, error) {
	resty.SetDebug(true)
	resty.SetContentLength(true)


	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/provinsi/%v", id)
	data := &models.ProvinsiStruct{}

	resp, err := resty.R().
		SetResult(data).
		Get(url)

	if err != nil {
		return nil, err
	}

	result := resp.String()
	if err := json.Unmarshal([]byte(result), data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *DaerahRepo) GetDataKotaKab(id int) (*models.KotaKabStruct, error) {
	resty.SetDebug(true)
	resty.SetContentLength(true)

	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kota?id_provinsi=%v", id)
	data := &models.KotaKabStruct{}

	resp, err := resty.R().
		SetResult(data).
		Get(url)

	if err != nil {
		return nil, err
	}

	result := resp.String()
	if err := json.Unmarshal([]byte(result), data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *DaerahRepo) GetDataKecamatan(id int) (*models.KecStruct, error) {
	resty.SetDebug(true)
	resty.SetContentLength(true)

	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kecamatan?id_kota=%v", id)
	data := &models.KecStruct{}

	resp, err := resty.R().
		SetResult(data).
		Get(url)

	if err != nil {
		return nil, err
	}

	result := resp.String()
	if err := json.Unmarshal([]byte(result), data); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *DaerahRepo) GetDataKelurahan(id int) (*models.KelStruct, error) {
	resty.SetDebug(true)
	resty.SetContentLength(true)

	url := fmt.Sprintf("https://dev.farizdotid.com/api/daerahindonesia/kelurahan?id_kecamatan=%v", id)
	data := &models.KelStruct{}

	resp, err := resty.R().
		SetResult(data).
		Get(url)

	if err != nil {
		return nil, err
	}

	result := resp.String()
	if err := json.Unmarshal([]byte(result), data); err != nil {
		return nil, err
	}

	return data, nil
}

