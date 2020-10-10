package controller

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"kantor/bpn/database"
	"kantor/bpn/model"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func GetSingleDOP() {
	// ambil dari env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	EnvUsername := os.Getenv("nama")
	EnvPassword := os.Getenv("password")
	alamatHost := os.Getenv("alamathost")

	// kasih flag untuk filter tanggal transaksi
	flagTanggalTransaksi := flag.String("tgl", "17/06/2020",
		"input tanggal transaksi AKTA")
	flag.Parse()

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
						jsonDataRes.Result[i].JENIS_HAK, *flagTanggalTransaksi, jsonDataRes.Result[i].LUASTANAH_OP)
				}
			}
		}
	}
}

func InsertDataBPN(NomorAkta, TanggalAkta, NamaPPAT, NOP, NTPD, NomorIndukBidang, KoordinatX, KoordinatY, NIK, NPWP,
	NamaWP, KelurahanOP, KecamatanOP, KotaOP, JenisHak, TanggalGet string, LuasTanahOP float32) {
	/*
		TODO snub on Sel 06 Okt 2020 09:37:02  : insert ke database di function ini
		- jika nomor akta sudah ada maka tidak bisa di insert v
		- usahakan nomer harus urut v
		- belum bisa parsing dari nama wp yg ada karakter ' nya
	*/

	var vNopSertifikatID *uint64
	var vNomorUrut uint64
	fmt.Printf("NIB: %+v, NOP: %+v, LuasTanah: %+v, NTPD: %+v\n", NomorIndukBidang, NOP, LuasTanahOP, NTPD)

	/*
		TODO snub on Sab 04 Jul 2020 11:11:28  : ambil config dari env file
	*/
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	envUser := os.Getenv("userpbb")
	envPass := os.Getenv("passwordpbb")
	envAddr := os.Getenv("addressserverpbb")
	envPort := os.Getenv("portserverpbb")
	envSN := os.Getenv("servicenamepbb")

	// log.Infof("userpbb: %v, passPBB: %v, addrPBB: %v, portPBB: %v, serviceNamePBB: %v", envUser, envPass, envAddr, envPort, envSN)

	kon, _ := database.KonekOracle(envUser, envPass, envAddr, envPort, envSN)

	getCounterNIB := fmt.Sprintf("select count(nop_sertifikat_id) from nop_sertifikat "+
		" where nomor_akta = '%s' and nomor_induk_bidang = '%s'", NomorAkta, NomorIndukBidang)
	hasilGetCounterNIB, err := kon.Query(getCounterNIB)
	if err != nil {
		log.Fatalf("errornya di baris 112: %v\n", err.Error())
	} else {
		// log.Infof("Sukses di baris 111: %v", hasilGetCounterNIB)
		for hasilGetCounterNIB.Next() {
			if err := hasilGetCounterNIB.Scan(&vNopSertifikatID); err != nil {
				log.Infof("errornya di baris 117: %v", err.Error())
			} /*else {
				log.Infof("sukses baris 118: %v", *vNopSertifikatID)
			}*/
		}
	}

	if *vNopSertifikatID == 0 {
		log.Infof("nomor akta atau nomor induk bidang belum ada")

		hasilGetNomorUrut, err := kon.Query("select ns.nop_sertifikat_id from nop_sertifikat ns where rownum = 1 order by ns.nop_sertifikat_id desc")
		if err != nil {
			log.Fatalf("errornya di baris 129: %v\n", err.Error())
		} else {
			log.Infof("Sukses di baris 131: %v", hasilGetNomorUrut)
			for hasilGetNomorUrut.Next() {
				if err := hasilGetNomorUrut.Scan(&vNomorUrut); err != nil {
					log.Infof("errornya di baris 134: %v", err.Error())
				}
				vNomorUrut = vNomorUrut + 1
			}
		}
		log.Infof("Nop sertifikat = %v, nomor urut = %v", *vNopSertifikatID, vNomorUrut)
		doInsert := fmt.Sprintf("insert into nop_sertifikat(nop_sertifikat_id, "+
			" nomor_akta, tanggal_akta, nama_ppat, nop, ntpd, nomor_induk_bidang, "+
			" koordinat_x, koordinat_y, nik, npwp, nama_wp, kelurahan_op, kecamatan_op, "+
			" kota_op, luastanah_op, jenis_hak, tanggal_get) values (%v, "+
			" `%s`, `%s`, `%s`, `%s`, `%s`, `%s`, "+
			" `%s`, `%s`, `%s`, `%s`, `%s`, `%s`, `%s`, "+
			" `%s`, %v, `%s`, `%s` )", vNomorUrut, NomorAkta, TanggalAkta, NamaPPAT, NOP, NTPD, NomorIndukBidang, KoordinatX, KoordinatY, NIK, NPWP,
			NamaWP, KelurahanOP, KecamatanOP, KotaOP, LuasTanahOP, JenisHak, TanggalGet)
		hasilDoInsert, err := kon.Exec(doInsert)
		if err != nil {
			log.Fatalf("errornya di baris 150: %v\n", err.Error())
		} else {
			log.Infof("Insert Sukses (%v)", hasilDoInsert)
		}
	} else {
		log.Infof("nomor akta atau nomor induk bidang sudah ada")
	}

	defer kon.Close()
}
