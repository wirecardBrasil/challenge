package main

import (
	"regexp"
)

type Brand int

const (
	Unknown   Brand = 0
	Amex      Brand = 1
	Diners    Brand = 2
	Elo       Brand = 3
	Hipercard Brand = 4
	Hiper     Brand = 5
	Master    Brand = 6
	Visa      Brand = 7
)

func ValidCreditCard(cardNumber string) bool {
	match, _ := regexp.MatchString("\\d{13,19}", cardNumber)
	return match
}

func AmexBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("^3[47]\\d{13}$", cardNumber)
	return match
}

func DinersBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("^3(0[0-5]|[68]\\d)\\d{11}$", cardNumber)
	return match
}

func EloBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("^((((636368)|(438935)|(504175)|(451416)|(636297))\\d{0,10})|((5067)|(4576)|(4011))\\d{0,12})$", cardNumber)
	return match
}

func HipercardBrand(cardNumber string) bool {
	var match bool
	match, _ = regexp.MatchString("^606282[0-9]{10}$", cardNumber)
	if !match {
		match, _ = regexp.MatchString("^3841(0|4|6)0[0-9]{13}$", cardNumber)
	}
	return match
}

func HiperBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("^(((637095)|(637612)|(637599)|(637609)|(637568))\\d{0,10})$", cardNumber)
	return match
}

func MasterBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("^(5[1-5]\\d{4}|677189)\\d{10}$", cardNumber)
	return match
}

func VisaBrand(cardNumber string) bool {
	match, _ := regexp.MatchString("4[0-9]{15}", cardNumber)
	return match
}

func GetBrand(cardNumber string) Brand {
	if AmexBrand(cardNumber) {
		return Amex
	}
	if DinersBrand(cardNumber) {
		return Diners
	}
	if EloBrand(cardNumber) {
		return Elo
	}
	if HipercardBrand(cardNumber) {
		return Hipercard
	}
	if HiperBrand(cardNumber) {
		return Hiper
	}
	if MasterBrand(cardNumber) {
		return Master
	}
	if VisaBrand(cardNumber) {
		return Visa
	}
	return Unknown
}
