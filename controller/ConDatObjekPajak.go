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
					fmt.Printf("[%v]: kd_kecamatan: %+v, kd_kelurahan: %+v, kd_blok: %+v, no_urut: %+v\n",
						i, jsonDataRes.Result[i].NOP[4:7], jsonDataRes.Result[i].NOP[7:10],
						jsonDataRes.Result[i].NOP[10:13], jsonDataRes.Result[i].NOP[13:17])
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
