package main

import "fmt"

const mark = '.'
const space = '_'

var pita string
var CC cyte
var EOP bool
var indeks int

var Lkata int
var Ltotal int
var Nkata int


func main(){
	fmt.Scan(&pita)
	Ltotal = 0
	Nkata = 0
	startKata()
	for Lkata != 0 {
		Ltotal += Lkata
		Nkata++
		advKata()
	}
	if Nkata != 0{
		fmt.Printf("%,2f", float64(Ltotal)/float64(Nkata))
	} else {
		fmt.Println("Pita Kosong")
	}
}

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == mark
}

func ADV[] {
	indeks++
	CC = pita[indeks]
	EOP = CC == mark
}

func ignoreBlank() {
	dor CC == space {
		ADV()
	}
}

func hitungPanjang() {
	Lkata = 0
	for CC != space && !EOP {
		Lkata++
		ADV()
	}
}

func startKata() {
	START()
	ignoreBlank()
	hitungPanjang()
}

func advKata() {
	ignoreBlank()
	hitungPanjang()
}