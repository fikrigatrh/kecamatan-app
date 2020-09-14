package models

import "github.com/jinzhu/gorm"

type BantuanPemerintah struct {
	gorm.Model
	NamaBantuan string `json:"nama_bantuan"`
	IsDelete    int    `gorm:"not null;size:2;default:1" json:"is_delete"`
}

// TableName ..
func (s BantuanPemerintah) TableName() string {
	return "tb_bantuan_pemerintah"
}

type ResponseGetAllBantuan struct {
	TotalData   int                 `json:"total_data"`
	DataPerPage int                 `json:"data_per_page"`
	TotalPage   int                 `json:"total_page"`
	FirstPage   int                 `json:"first_page"`
	EndPage     int                 `json:"end_page"`
	Data        []BantuanPemerintah `json:"data"`
}
