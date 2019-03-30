package main

type Person struct {
	Id    int64  `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Cpf   string `json:"cpf"`
}

type Client Person

type Clients []Client

type Buyer Person

type ReturnStruct struct {
	State            int    `json:"state"`
	Message          string `json:"message"`
	TechnicalMessage string `json:"technicalMessage"`
}

type ReturnClient struct {
	Return ReturnStruct `json:"return"`
	Client Client       `json:client`
}
