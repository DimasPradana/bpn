package model

type StructResSingleDOP struct {
	NOMOR_AKTA         string  `json:"nomor_akta"`
	TANGGAL_AKTA       string  `json:"tanggal_akta"`
	NAMA_PPAT          string  `json:"nama_ppat"`
	NOP                string  `json:"nop"`
	NTPD               string  `json:"ntpd"`
	NOMOR_INDUK_BIDANG string  `json:"nomor_induk_bidang"`
	KOORDINAT_X        string  `json:"koordinat_x"`
	KOORDINAT_Y        string  `json:"koordinat_y"`
	NIK                string  `json:"nik"`
	NPWP               string  `json:"npwpd"`
	NAMA_WP            string  `json:"nama_wp"`
	KELURAHAN_OP       string  `json:"kelurahan_op"`
	KECAMATAN_OP       string  `json:"kecamatan_op"`
	KOTA_OP            string  `json:"kota_op"`
	LUASTANAH_OP       float32 `json:"luastanah_op"`
	JENIS_HAK          string  `json:"jenis_hak"`
}

type StructReqSingleDOP struct {
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
	TANGGAL  string `json:"tanggal"`
}

type StructResSingleDOP2 struct {
	Result      []StructResSingleDOP `json:"result"`
	Respon_code string               `json:"respon_code"`
}
