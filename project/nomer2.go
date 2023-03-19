package main

import "fmt"

var pita string
var CC, CC2 byte
var EOP, EOP2 bool
var indeks, angka, angka1 int
var a, b, hasil bool
var tes string

func main() {
	fmt.Scanln(&pita)
	START()
	if !EOP {
		tes = pemisah()
		a = tes == "IF" || tes == "SE" || tes == "IT"
	}
	for !EOP {
		tes = pemisah()
		tes = tes + "."
		angka = 0
		angka1 = 0
		START2()
		for !EOP2 {
			angka1++
			if DIGIT2() {
				angka++
			}
			ADV2()
		}
		b = angka == angka1 && angka > 0
		pemisah()
	}
	hasil = a && b
	fmt.Println(hasil)
}

func START() {
	indeks = 0
	CC = pita[0]
	EOP = CC == '.'
}
func START2() {
	indeks = 0
	CC2 = tes[indeks]
	EOP2 = CC2 == '.'
}

func ADV() {
	indeks = indeks + 1
	CC = pita[indeks]
	EOP = CC == '.'
}

func ADV2() {
	indeks = indeks + 1
	CC2 = tes[indeks]
	EOP2 = CC2 == '.'
}

func DIGIT() bool {
	return CC >= 48 && CC <= 57
}
func DIGIT2() bool {
	return CC2 >= 48 && CC2 <= 57
}

func pemisah() string {
	var bagian string
	for CC == 45 {
		bagian = bagian + string(CC)
		ADV()
	}
	bagian = string(CC)
	if DIGIT() {
		ADV()
		for DIGIT() {
			bagian = bagian + string(CC)
			ADV()
		}
	} else if CC >= 65 && CC <= 90 || CC >= 97 && CC <= 122 {
		ADV()
		for DIGIT() || CC >= 65 && CC <= 90 || CC >= 97 && CC <= 122 {
			bagian = bagian + string(CC)
			ADV()
		}
	}
	return bagian
}
