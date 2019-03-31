package main

type Person struct {
	Id    int64  `json:"id"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Cpf   string `json:"cpf,omitempty"`
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
