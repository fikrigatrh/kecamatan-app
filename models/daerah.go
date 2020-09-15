package models

type ProvinsiStruct struct {
	Provinsi []DataProvinsi `json:"provinsi"`
}

type DataProvinsi struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type KotaKabStruct struct {
	KotaKabupaten []DataKotaKab `json:"kota_kabupaten"`
}

type DataKotaKab struct {
	ID         int    `json:"id"`
	IDProvinsi string `json:"id_provinsi"`
	Nama       string `json:"nama"`
}
