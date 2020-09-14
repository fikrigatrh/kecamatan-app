package models

import "github.com/jinzhu/gorm"

type BantuanPemerintah struct {
	gorm.Model
	NamaBantuan string `json:"nama_bantuan"`
}

// TableName ..
func (s BantuanPemerintah) TableName() string {
	return "tb_bantuan_pemerintah"
}
