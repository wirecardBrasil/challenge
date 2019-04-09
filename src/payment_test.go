package main

import "testing"

func TestValidPaymentType(t *testing.T) {
	expected := true
	if observed := validatePaymentType(1); observed != expected {
		t.Fatalf("validPaymentType() = %v, want %v", observed, expected)
	}
}

func TestInvalidPaymentsType(t *testing.T) {
	expected := false
	if observed := validatePaymentType(5); observed != expected {
		t.Fatalf("validPaymentType() = %v, want %v", observed, expected)
	}
}

func TestGenerateBoleto(t *testing.T) {
	expected := "23790504004199033014836008109203478470000019900"
	if observed := BoletoPayment(); observed != expected {
		t.Fatalf("BoletoPayment() = %v, want %v", observed, expected)
	}
}

func TestCardPayment(t *testing.T) {
	expected := true
	if observed := CardPayment(); observed != expected {
		t.Fatalf("CardPayment() = %v, want %v", observed, expected)
	}

}

func TestValidCreditCard(t *testing.T) {
	expected := true
	if observed := ValidCreditCard("3542449243588413"); observed != expected {
		t.Fatalf("ValidCreditCard() = %v, want %v", observed, expected)
	}
}

func TestInvalidCreditCard(t *testing.T) {
	expected := false
	if observed := ValidCreditCard("129738465"); observed != expected {
		t.Fatalf("ValidCreditCard() = %v, want %v", observed, expected)
	}
}

func TestAmex(t *testing.T) {
	expected := Amex
	if observed := GetBrand("341111111111111"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}

func TestDiners(t *testing.T) {
	expected := Diners
	if observed := GetBrand("30111122223331"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestElo(t *testing.T) {
	expected := Elo
	if observed := GetBrand("4514160123456789"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestHipercard(t *testing.T) {
	expected := Hipercard
	if observed := GetBrand("3841001111222233334"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestHiper(t *testing.T) {
	expected := Hiper
	if observed := GetBrand("6370950000000005"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestMaster(t *testing.T) {
	expected := Master
	if observed := GetBrand("5555666677778884"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestVisa(t *testing.T) {
	expected := Visa
	if observed := GetBrand("4111111111111111"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
func TestUnknown(t *testing.T) {
	expected := Unknown
	if observed := GetBrand("1"); observed != expected {
		t.Fatalf("GetBrand() = %v, want %v", observed, expected)
	}
}
