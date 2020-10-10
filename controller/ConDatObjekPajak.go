package controller

// {{{ import
import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"kantor/bpn/model"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// }}}

// {{{ GetSingleDOP
func GetSingleDOP() {
	/**
	TODO snub on Sel 06 Okt 2020 07:51:15  : alamat servis yang baru
	› alamatHost = https://services.atrbpn.go.id/BPNApiService/Api/BPHTB/getDataATRBPN
	› json body = {"USERNAME":"bapendakabsitubondo", "PASSWORD":"b3kn4s4p4", "TANGGAL":"08/01/2020"}
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
	*/

	// ambil dari env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	EnvUsername := os.Getenv("nama")
	EnvPassword := os.Getenv("password")

	// kasih flag untuk filter tanggal transaksi
	flagTanggalTransaksi := flag.String("tgl", "17/06/2020",
		"input tanggal transaksi AKTA")
	flag.Parse()

	// alamat server
	/** alamatHost := "http://103.49.37.84:8080/BPNApiService/Api/BPHTB/getDataBPN" */
	alamatHost := "https://services.atrbpn.go.id/BPNApiService/Api/BPHTB/getDataATRBPN"

	// tampung struct request dari model
	jsonDataReq := model.StructReqSingleDOP{USERNAME: EnvUsername,
		PASSWORD: EnvPassword, TANGGAL: *flagTanggalTransaksi}
	// variable untuk format jsonnya
	var jsonDataRes model.StructResSingleDOP2

	// men-json kan requestnya
	jsonValue, _ := json.Marshal(jsonDataReq)
	response, err := http.Post(alamatHost, "application/json",
		bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Http request failed on line 68: %v\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal([]byte(data), &jsonDataRes)
		if err != nil {
			log.Fatalf("error unmarshal di line 73: %+v\n", err)
		} else {
			/*
				TODO dimas on Rab 24 Jun 2020 03:08:40
				› coba koneksi ke database (oracle, mysql),
				› insert data balikan dari BPN ke database, (buat/pake controller)
			*/
			// log.Infof("%+v\n\n", jsonDataRes.Result[0]) // contoh parsing pake slice
			if len(jsonDataRes.Result) == 0 {
				log.Infof("Result kosong untuk tanggal: %+v\n",
					*flagTanggalTransaksi)
			} else {
				for i := 0; i < len(jsonDataRes.Result); i++ {
					// kd_kecamatan := jsonDataRes.Result[i].NOP[4:7]
					// kd_kelurahan := jsonDataRes.Result[i].NOP[7:10]
					// kd_blok := jsonDataRes.Result[i].NOP[10:13]
					// no_urut := jsonDataRes.Result[i].NOP[13:17]
					// fmt.Printf("[%v]: kd_kecamatan: %+v, kd_kelurahan: %+v, kd_blok: %+v, no_urut: %+v\n",
					// 	i, kd_kecamatan, kd_kelurahan, kd_blok, no_urut)
					fmt.Printf("[%v]: kd_kecamatan: %+v, kd_kelurahan: %+v, kd_blok: %+v, no_urut: %+v\n",
						i, jsonDataRes.Result[i].NOP[4:7], jsonDataRes.Result[i].NOP[7:10],
						jsonDataRes.Result[i].NOP[10:13], jsonDataRes.Result[i].NOP[13:17])
					/**
					fmt.Printf("[%v]: NomorAkta : %+v\n", i, jsonDataRes.Result[i].NOMOR_AKTA)
					fmt.Printf("[%v]: TanggalAkta : %+v\n", i, jsonDataRes.Result[i].TANGGAL_AKTA)
					fmt.Printf("[%v]: NamaPPAT : %+v\n", i, jsonDataRes.Result[i].NAMA_PPAT)
					fmt.Printf("[%v]: NOP : %+v\n", i, jsonDataRes.Result[i].NOP)
					fmt.Printf("[%v]: NTPD: %+v\n", i, jsonDataRes.Result[i].NTPD)
					fmt.Printf("[%v]: NIB: %+v\n", i, jsonDataRes.Result[i].NOMOR_INDUK_BIDANG)
					fmt.Printf("[%v]: KoordinatX: %+v\n", i, jsonDataRes.Result[i].KOORDINAT_X)
					fmt.Printf("[%v]: KoordinatY: %+v\n", i, jsonDataRes.Result[i].KOORDINAT_Y)
					fmt.Printf("[%v]: NIK: %+v\n", i, jsonDataRes.Result[i].NIK)
					fmt.Printf("[%v]: NPWP: %+v\n", i, jsonDataRes.Result[i].NPWP)
					fmt.Printf("[%v]: NamaWP: %+v\n", i, jsonDataRes.Result[i].NAMA_WP)
					fmt.Printf("[%v]: KelurahanOP: %+v\n", i, jsonDataRes.Result[i].KELURAHAN_OP)
					fmt.Printf("[%v]: KecamatanOP: %+v\n", i, jsonDataRes.Result[i].KECAMATAN_OP)
					fmt.Printf("[%v]: KotaOP: %+v\n", i, jsonDataRes.Result[i].KOTA_OP)
					fmt.Printf("[%v]: LuasTanah_OP: %+v\n", i, jsonDataRes.Result[i].LUASTANAH_OP)
					fmt.Printf("[%v]: JenisHak: %+v\n\n", i, jsonDataRes.Result[i].JENIS_HAK)
					*/
					InsertDataBPN(jsonDataRes.Result[i].NOMOR_AKTA, jsonDataRes.Result[i].TANGGAL_AKTA,
						jsonDataRes.Result[i].NAMA_PPAT, jsonDataRes.Result[i].NOP,
						jsonDataRes.Result[i].NTPD, jsonDataRes.Result[i].NOMOR_INDUK_BIDANG,
						jsonDataRes.Result[i].KOORDINAT_X, jsonDataRes.Result[i].KOORDINAT_Y,
						jsonDataRes.Result[i].NIK, jsonDataRes.Result[i].NPWP,
						jsonDataRes.Result[i].NAMA_WP, jsonDataRes.Result[i].KELURAHAN_OP,
						jsonDataRes.Result[i].KECAMATAN_OP, jsonDataRes.Result[i].KOTA_OP,
						jsonDataRes.Result[i].JENIS_HAK, jsonDataRes.Result[i].LUASTANAH_OP)
				}
			}
		}
	}
}

// }}}
// {{{ InsertDataBPN
func InsertDataBPN(NomorAkta, TanggalAkta, NamaPPAT, NOP, NTPD, NomorIndukBidang, KoordinatX, KoordinatY, NIK, NPWP,
	NamaWP, KelurahanOP, KecamatanOP, KotaOP, JenisHak string, LuasTanahOP float32) {
	/*
		TODO snub on Sel 06 Okt 2020 09:37:02  : insert ke database di function ini
	*/
	fmt.Printf("NIB: %+v, NOP: %+v, LuasTanah: %+v, NTPD: %+v\n", NomorIndukBidang, NOP, LuasTanahOP, NTPD)
}

// }}}
