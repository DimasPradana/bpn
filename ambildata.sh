#!/bin/bash
# $bulan = 5
for number in {1..31}
do
  # reset && go run main.go -tgl="$number/05/2020"
  # go run main.go -tgl="$number/$@/2020"
  go run main.go -tgl="$number/$1/$2"
  # echo "number : $number"
  # echo "masukkan bulan dengan format angka 2 digit: $1"
  # echo "masukkan tahun dengan format angka 4 digit: $2"
done
exit 0
