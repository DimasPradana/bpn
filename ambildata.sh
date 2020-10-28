#!/bin/bash
# $bulan = 5
for number in {1..31}
do
  # reset && go run main.go -tgl="$number/05/2020"
  go run main.go -tgl="$number/$@/2020"
done
exit 0
