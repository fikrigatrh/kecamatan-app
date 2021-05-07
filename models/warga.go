package models

import "github.com/jinzhu/gorm"

type ResponseGetAllDataWarga struct {
	TotalData   int         `json:"total_data"`
	DataPerPage int         `json:"data_per_page"`
	TotalPage   int         `json:"total_page"`
	FirstPage   int         `json:"first_page"`
	EndPage     int         `json:"end_page"`
	Duration    string      `json:"duration"`
	Data        []DataWarga `json:"data"`
}

type DataWarga struct {
	gorm.Model
	Nik                 int    `json:"nik"`
	NoKk                int    `json:"no_kk"`
	SudahMemilikiKTP    string `json:"sudah_memiliki_ktp"`
	NamaLengkap         string `json:"nama_lengkap"`
	JenisKelamin        string `json:"jenis_kelamin"`
	Pendidikan          string `json:"pendidikan"`
	TempatLahir         string `json:"tempat_lahir"`
	TanggalLahir        string `json:"tanggal_lahir"`
	Agama               string `json:"agama"`
	GolDarah            string `json:"gol_darah"`
	ProvinsiID          int    `sql:"-" json:"provinsi_id,omitempty"`
	KotaKabID           int    `sql:"-" json:"kota_kab_id,omitempty"`
	KecamatanID         int    `sql:"-" json:"kecamatan_id,omitempty"`
	KelurahanID         int    `sql:"-" json:"kelurahan_id,omitempty"`
	Alamat              string `sql:"-" json:"alamat,omitempty"`
	RW                  string `sql:"-" json:"rw,omitempty"`
	RT                  string `sql:"-" json:"rt,omitempty"`
	DetailAlamat        string `sql:"type:JSONB NOT NULL DEFAULT '{}'::JSONB" json:"detail_alamat"`
	StatusPerkawinan    string `json:"status_perkawinan"`
	NoBukuNikah         int    `json:"no_buku_nikah"`
	JenisPekerjaan      string `json:"jenis_pekerjaan"`
	SHOK                string `json:"shok"`
	Kewarganegaraan     string `json:"kewarganegaraan"`
	Disabilitas         string `json:"disabilitas"`
	NoPaspor            int    `json:"no_paspor"`
	NoKitasKitap        int    `json:"no_kitas_kitap"`
	NamaAyah            string `json:"nama_ayah"`
	NamaIbu             string `json:"nama_ibu"`
	AlamatSesuaiKK      string `json:"alamat_sesuai_kk"`
	StatusTempatTinggal string `json:"status_tempat_tinggal"`
	NoShmHcb            int    `json:"no_shm_hcb"`
	NoAjb               int    `json:"no_ajb"`
	NoSpptPbb           int    `json:"no_sppt_pbb"`
	BantuanPemerintahID int    `json:"bantuan_pemerintah_id"`
	JenPend             string `json:"jen_pend"`
}

// TableName ..
func (s DataWarga) TableName() string {
	return "tb_data_warga"
}

type DetailAlamat struct {
	Provinsi      string `json:"provinsi"`
	KotaKabupaten string `json:"kota_kabupaten"`
	Kecamatan     string `json:"kecamatan"`
	Kelurahan     string `json:"kelurahan"`
	Alamat        string `json:"alamat"`
	RW            string `json:"rw"`
	RT            string `json:"rt"`
}
