package models

import "github.com/jinzhu/gorm"

func InitTable(db *gorm.DB) {

	//db.DropTableIfExists(&BantuanPemerintah{})
	//db.DropTableIfExists(&DataWarga{})
	//db.DropTableIfExists(&User{})

	db.AutoMigrate(&Auth{})
	db.AutoMigrate(&User{})
	db.AutoMigrate(&Role{})
	db.AutoMigrate(&UserRole{})
	db.AutoMigrate(&BantuanPemerintah{})
	db.AutoMigrate(&DataWargaExample{})

	db.Model(&UserRole{}).AddForeignKey("role_id", "tb_role(id)", "CASCADE", "CASCADE")
	db.Model(&UserRole{}).AddForeignKey("user_id", "tb_user(id)", "CASCADE", "CASCADE")

}

