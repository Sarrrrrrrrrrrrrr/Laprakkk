package main

import "fmt"

func main() {
	var Nama, Nim, kelas, region string
	fmt.Scan(&Nama, &Nim, &kelas, &region)

	teks1 := "Perkenalkan nama saya " + Nama + ", "
	teks2 := " NIM saya " + Nim + ", "
	teks3 := "Saya dari kelas " + kelas + ", "
	teks4 := "asal saya dari " + region + ". "

	fmt.Print(teks1, teks2, teks3, teks4)

}
