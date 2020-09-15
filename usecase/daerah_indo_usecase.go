package usecase

import (
	"kecamatan_app/models"
	"kecamatan_app/repo"
)

type DaerahIndoUseCaseStruct struct {
	daerahIndoUsecase repo.DaerahRepoInterface
}

type DaerahIndoUsecaseInterface interface {
	GetDataProvinsiByID(id int) (*models.ProvinsiStruct, error)
}

func CreateDaerahUsecase(daerahIndoUsecase repo.DaerahRepoInterface) DaerahIndoUsecaseInterface {
	return &DaerahIndoUseCaseStruct{daerahIndoUsecase}
}

func (d *DaerahIndoUseCaseStruct) GetDataProvinsiByID(id int) (*models.ProvinsiStruct, error)  {
	data, err := d.daerahIndoUsecase.GetDataProvinsiByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}
