// {{{ import
package controller

/**
* @author snub on Jum 01 Jan 2021 22:43:24
* /home/snub/skripsi/bphtb/bpn-go/controller/HomeController.go
 */

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"bpn-go/model"
	// "strings"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

// }}}

// {{{ init
func init() {
	// Only log the warning severity or above
	log.SetLevel(log.TraceLevel)
}

// }}}

// {{{AmbilNop
func AmbilNOP() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error ambil .env file")
	}

	EnvUsername := os.Getenv("nama")
	EnvPassword := os.Getenv("password")
	EnvAlamatHost := os.Getenv("alamatHost")

	flagTanggalTransaksi := flag.String("tgl", "26/11/2020", "input tanggal transaksi AKTA")
	flag.Parse()

	//  TODO snub on Jum 01 Jan 2021 22:42:59 : lanjutkan yang belum selesai, hanya untuk request dan respon NOP saja //
	// dump var requirement
	// log.Infof("EnvUsername: %v, EnvPassword: %v, EnvAlamatHost: %v, flagTanggalTransaksi: %v", EnvUsername, EnvPassword, EnvAlamatHost, *flagTanggalTransaksi)

	// ambil struct dari model request
	jsonReq := model.StructReqSingleDOP{
		USERNAME: EnvUsername,
		PASSWORD: EnvPassword, TANGGAL: *flagTanggalTransaksi,
	}
	var jsonRes model.StructResSingleDOP2

	// inisialisasi waktu sekarang
	waktuSekarang := time.Now()

	// request dijadikan json
	jsonValue, _ := json.Marshal(jsonReq)
	response, err := http.Post(EnvAlamatHost, "application/json", bytes.NewBuffer(jsonValue))
	if err != nil {
		log.Fatalf("Http request failed on line 66: %v\n", err)
	} else {
		data, _ := ioutil.ReadAll(response.Body)
		err := json.Unmarshal([]byte(data), &jsonRes)
		if err != nil {
			log.Fatalf("error unmarshal di line 73: %+v\n", err)
		} else {
			if len(jsonRes.Result) == 0 {
				log.Infof("Result kosong untuk tanggal %v\n", *flagTanggalTransaksi)
				durasi := time.Since(waktuSekarang)
				log.WithFields(log.Fields{
					// "latency" : durasi.Nanoseconds() ,
					"latency": durasi.Seconds(),
					"jenis":   "QUERY",
				}).Tracef("waktu yang dibutuhkan : %v detik", durasi.Seconds())
			} else {
				for i := 0; i < len(jsonRes.Result); i++ {
					fmt.Printf("NO : %v, NOP : %v\n", i, jsonRes.Result[i].NOP)
				}
				durasi := time.Since(waktuSekarang)
				log.WithFields(log.Fields{
					// "latency" : durasi.Nanoseconds() ,
					"latency": durasi.Seconds(),
					"jenis":   "QUERY",
				}).Infof("waktu yang dibutuhkan : %v detik", durasi.Seconds())
			}
		}
	}
}

// }}}
