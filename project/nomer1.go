package main

import "fmt"

// deklarasi variabel global
var pita string
var CC byte
var EOP bool
var indeks int

func START() {
	indeks = 0
	CC = pita[indeks]
	EOP = CC == '.'
}
func ADV() {
	indeks++
	CC = pita[indeks]
	EOP = CC == '.'
}
func DIGIT() bool {
	return CC >= '0' && CC <= '9'
}
func main() {
	var valid bool = false
	fmt.Scan(&pita)
	START()
	if CC == 'I' {
		ADV()
		if CC == 'T' {
			ADV()
			if CC == '-' {
				ADV()
				if DIGIT() {
					for DIGIT() {
						ADV()
					}
					valid = EOP
				}
			}
		}
	}
	fmt.Println(valid)
}
