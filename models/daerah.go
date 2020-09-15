package models

type ProvinsiStruct struct {
	ID   int    `json:"id"`
	Nama string `json:"nama"`
}

type DataProvinsi struct {
}

type KotaKabStruct struct {
	KotaKabupaten []DataKotaKab `json:"kota_kabupaten"`
}

type DataKotaKab struct {
	ID         int    `json:"id"`
	IDProvinsi string `json:"id_provinsi"`
	Nama       string `json:"nama"`
}

type KecStruct struct {
	Kecamatan []DataKec `json:"kecamatan"`
}

type DataKec struct {
	ID     int    `json:"id"`
	IDKota string `json:"id_kota"`
	Nama   string `json:"nama"`
}

type KelStruct struct {
	Kelurahan []DataKel `json:"kelurahan"`
}

type DataKel struct {
	ID          int    `json:"id"`
	IDKecamatan string `json:"id_kecamatan"`
	Nama        string `json:"nama"`
}
