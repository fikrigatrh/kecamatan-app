package usecase

import (
	"errors"
	"kecamatan_app/models"
	"kecamatan_app/repo"
)

type BantuanPemerintahUseCaseStruct struct {
	bantuanPemerintah repo.BantuanPemerintahInterface
}

type BantuanPemerintahUsecaseInterface interface {
	CreateBantuan(data *models.BantuanPemerintah) (*models.BantuanPemerintah, error)
	GetDataByID(id int) (*models.BantuanPemerintah, error)
	GetAllData() (*models.ResponseGetAllBantuan, error)
	UpdateData(id int, data *models.BantuanPemerintah) (*models.BantuanPemerintah, error)
	CheckData(namaBantuan string) bool
}

func CreateBantuanUsecase(bantuanPemerintah repo.BantuanPemerintahInterface) BantuanPemerintahUsecaseInterface {
	return &BantuanPemerintahUseCaseStruct{bantuanPemerintah}
}

func (b BantuanPemerintahUseCaseStruct) CreateBantuan(data *models.BantuanPemerintah) (*models.BantuanPemerintah, error) {
	result, err :=b.bantuanPemerintah.CreateBantuan(data)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (b BantuanPemerintahUseCaseStruct) GetDataByID(id int) (*models.BantuanPemerintah, error) {
	data, err := b.bantuanPemerintah.GetDataByID(id)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (b BantuanPemerintahUseCaseStruct) GetAllData() (*models.ResponseGetAllBantuan, error) {
	data, err := b.bantuanPemerintah.GetAllData()
	if err != nil {
		return nil, errors.New("Oppss, something error")
	}

	return data, nil
}

func (b BantuanPemerintahUseCaseStruct) UpdateData(id int, data *models.BantuanPemerintah) (*models.BantuanPemerintah, error) {
	firstData, err := b.bantuanPemerintah.GetDataByID(id)
	if err != nil {
		return nil, errors.New("ID does not exist")
	}

	if data.NamaBantuan == "" {
		data.NamaBantuan = firstData.NamaBantuan
	}

	student, err := b.bantuanPemerintah.UpdateData(id, data)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (b BantuanPemerintahUseCaseStruct) CheckData(namaBantuan string) bool {
	data := b.bantuanPemerintah.CheckData(namaBantuan)
	return data
}
