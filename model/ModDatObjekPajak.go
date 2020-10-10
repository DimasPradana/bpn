package model

/** {{{ TODO
TODO snub on Sel 06 Okt 2020 07:51:15  : alamat servis yang baru
- alamatHost = https://services.atrbpn.go.id/BPNApiService/Api/BPHTB/getDataATRBPN
- json body = {"USERNAME":"bapendakabsitubondo", "PASSWORD":"b3kn4s4p4", "TANGGAL":"08/01/2020"}
response{
	"NOMOR_AKTA": "320/2018",
	"TANGGAL_AKTA": "16/03/2018",
	"NAMA_PPAT": "DIVI IKA RAHMAWATI, SH. M.KN",
	"NOP": "351207000602200310",
	"NTPD": "1530",
	"NOMOR_INDUK_BIDANG": "12350606.01304",
	"KOORDINAT_X": "113.996256671961",
	"KOORDINAT_Y": "-7.69111791059957",
	"NIK": "3512086207870003",
	"NPWP": null,
	"NAMA_WP": "RINA ANDRIANI",
	"KELURAHAN_OP": "ALAS MALANG",
	"KECAMATAN_OP": "PANARUKAN",
	"KOTA_OP": "Kabupaten Situbondo",
	"LUASTANAH_OP": 89.0,
	"JENIS_HAK": "Hak Guna Bangunan"
}
 }}} */

// {{{ StructResSingleDOP
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

/**type StructResSingleDOP struct {
	AKTAID string `json:"aktaid"`
	TGL_AKTA string `json:"tgl_akta"`
	NOP string `json:"nop"`
	NIB string `json:"nib"`
	NIK string `json:"nik"`
	NPWP string `json:"npwp"`
	NAMA_WP string `json:"nama_wp"`
	ALAMAT_OP string `json:"alamat_op"`
	KELURAHAN_OP string `json:"kelurahan_op"`
	KECAMATAN_OP string `json:"kecamatan_op"`
	KOTA_OP string `json:"kota_op"`
	LUASTANAH_OP int `json:"luastanah_op"`
	LUASBANGUNAN_OP int `json:"luasbangunan_op"`
	PPAT string `json:"ppat"`
	NO_SERTIPIKAT string `json:"no_sertipikat"`
	NO_AKTA string `json:"no_akta"`
}*/
// }}}

// {{{ StructReqSingleDOP
type StructReqSingleDOP struct {
	USERNAME string `json:"username"`
	PASSWORD string `json:"password"`
	TANGGAL  string `json:"tanggal"`
}

// }}}

// {{{ StructResSingleDOP2
type StructResSingleDOP2 struct {
	Result      []StructResSingleDOP `json:"result"`
	Respon_code string               `json:"respon_code"`
}

// }}}
