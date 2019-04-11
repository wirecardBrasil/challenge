package main

/*
import (
	"database/sql"
	//"fmt"
	//"time"
)
*/

type Person struct {
	Id      int64  `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Email   string `json:"email,omitempty"`
	CpfCnpj string `json:"cpfCnpj,omitempty"`
}

type Client Person

type Clients []Client

type Buyer Person

type ReturnStruct struct {
	State            int    `json:"state"`
	Message          string `json:"message,omitempty"`
	TechnicalMessage string `json:"technicalMessage,omitempty"`
}

type ReturnClients struct {
	Return  ReturnStruct `json:"return"`
	Clients Clients      `json:clients`
}

type CustomError struct {
	Message          string
	TechnicalMessage string
	IdMessage        int
}

type CardInfos struct {
	HolderName     string `json:"holderName,omitempty"`
	Number         string `json:"number,omitempty"`
	ExpirationDate string `json:"expirationDate,omitempty"`
	Cvv            string `json:"cvv,omitempty"`
}

type PaymentInfo struct {
	PaymentID    int64        `json:"paymentId,omitepmpty"`
	Amount       float64      `json:"amount,omitempty"`
	PaymentType  int          `json:"type,omitempty"`
	Card         CardInfos    `json:"card,omitempty"`
	Boleto       BoletoReturn `json:"boleto,omitempty"`
	PaymentState int          `json:"paymentState,omitempty"`
}

type Payment struct {
	Client      Client      `json:"client,omitempty"`
	Buyer       Buyer       `json:"buyer,omitempty"`
	PaymentInfo PaymentInfo `json:"payment,omitempty"`
}

type BoletoReturn struct {
	Number string `json:"number,omitempty"`
}

type CardReturn struct {
	Successful bool `json:"successful"`
}

type PaymentDataReturn struct {
	PaymentId int64        `json:"paymentId,omitempty"`
	StateId   int          `json:"paymentState,omitempty"`
	Card      CardReturn   `json:"card,omitempty"`
	Boleto    BoletoReturn `json:"boleto,omitempty"`
}

type PaymentReturn struct {
	Return ReturnStruct `json:"return,omitempty"`
	//Card    CardReturn   `json:"card,omitempty"`
	//Boleto  BoletoReturn `json:"boleto,omitempty"`
	Payment PaymentDataReturn `json:"payment,omitempty"`
}

type Payments []Payment

type PaymentConsult struct {
	Return   ReturnStruct `json:"return,omitempty"`
	Payments Payments     `json:"payments,omitempty"`
}

type BrandData struct {
	Code Brand  `json:"code"`
	Name string `json:"name,omitempty"`
}
