# Payment API

This API is intended to support payment workflow.
It uses Golang as programming language and MySQL as database.

## Methods:
- POST /client
- GET /client
- GET /client/:idclient
- POST /payment
- PUT /payment/:idpayment/state/:idstate
- GET /payment/:idpayment
- GET /card/validate/:cardnumber
- GET /card/brand/:cardnumber

###  POST /client
Method is intended to insert a new client.

>URL: http://localhost:8080/client

**Header:**  
"Content-Type": "application/json"

**Body:**
```
{
    "name": "string",
    "email": "string",
    "cpfCnpj": "999999999"
}
```
Where,  
**name** -> Client's name.  
**email** -> Client's email.  
**cpfCnpj** -> Client's CPF or CNPJ  

**Return**
```
[
    {
        id: 9,
        "name": "string",
        "email": "string",
        "cpfCnpj": "999999999"
    }
]
```
Where,  
**id** -> Id of inserted client.  
**name** -> Client's name.  
**email** -> Client's email.  
**cpfCnpj** -> Client's CPF or CNPJ  

****
###  GET /client
Method is intended to return all clients.

>URL: http://localhost:8080/client

**Return**
```
{
    "return": {
        "state": 9
    },
    "Clients": [
        {
            "id": 9,
            "name": "string",
            "email": "string",
            "cpfCnpj": "99999999999"
        }
    ]
}
```

Where,  
**state** -> State of return. 0 = false, 1 = true.  
**id** -> Client's id.  
**name** -> Client's name.  
**email** -> Client's email.  
**cpfCnpj** -> Client's CPF or CNPJ  

****
###  GET /client/:idclient
Method is intended to return an especific client.

>URL:http://localhost:8080/client/:idclient

Where,  
**:idclient** -> Is the client's id to be consulted.  

Example:
>http://localhost:8080/client/10  
This will return the informations of client with id 10.

**Return**
```
{
    "return": {
        "state": 9
    },
    "Clients": [
        {
            "id": 9,
            "name": "string",
            "email": "string",
            "cpfCnpj": "99999999999"
        }
    ]
}
```

Where,  
**state** -> State to indicates success of opperaion. 0 = false, 1 = true.  
**id** -> Client's id.  
**name** -> Client's name.  
**email** -> Client's email.  
**cpfCnpj** -> Client's CPF or CNPJ  

****
###  POST /payment
Method is intended to create new payment.

>URL: http://localhost:8080/payment

**Header**  
"Content-Type": "application/json"

**Body**
```
{
	"client":{
		"id": 1
	},
	"buyer":{
		"name": "string",
		"email": "string",
		"cpfCnpj": "99999999999"	
	},
	"payment":{
		"amount":99.99,
		"type":9,
		"card":{
			"holderName":"string",
			"number":"9999999999999",
			"expirationDate":"01/01",
			"cvv":"999"
		}
	}
}
```
Where,  
**client.id** -> Client's id.  
**buyer.name** -> Buyer's name.  
**buyer.email** -> Buyer's email.  
**buyer.cpfCnpj** -> Buyer's CPF or CNPJ.  
**payment.amount** -> Payment's amount.  
**payment.type** -> Type of payment. 1 = boleto, 2 = credit card.  
**card.holderName** -> Name of card holder.  
**card.number** -> Card's number.  
**card.expirationDate** -> Card's expiration date.  
**card.cvv** -> Card's verification value.  
*Note: card infos should only send when payment type was 2 (credit card).*  

**Return**
```
{
    "return": {
        "state": 9,
        "message": "string."
    },
    "payment": {
        "paymentId": 9,
        "paymentState": 9,
        "card": {
            "successful": true
        },
        "boleto": {
            "number": "99999999999999999999999999999999999999999999999"
        }
    }
}

```

**state** -> State to indicates success of opperaion. 0 = false, 1 = true.  
**message** -> Message of opperation.  
**payment.paymentId** -> Payment's id.  
**payment.paymentState** -> Payment's state. 1 - "Pending", 2 - "Paid", 3 - "Canceled", 4 - "Refused".  
**card.successful** -> Informs if card's payment was successful or not. Returning true or false, respectively.  
**boleto.number** -> Returns the boleto's number to payment.  

*Notes:*  
- *card.successful is only returned when is a card payment (type 2).*  
- *boleto.number is only returned when is a boleto payment (type 1).*  

****
###  PUT /payment/:idpayment/state/:idstate
Method is intended to update payment's state.

>URL: http://localhost:8080/payment/:idpayment/state/:idstate

Where,  
**:idpayment** -> Payment's id.  
**:idstate** -> State that will be associated to payment. 1 - "Pending", 2 - "Paid", 3 - "Canceled", 4 - "Refused".  

Example:
>http://localhost:8080/payment/10/state/2  
This will alter the state of payment 10 to "Paid".

**Return**
```
{
    "state": 9,
    "message": "string",
    "technicalMessage": "string"
}
```

Where,   
**state** -> State to indicates success of opperaion. 0 = false, 1 = true.  
**message** -> Opperation's message.  
**technicalMessage** -> Technical message, describes detailed erros if occurs.  

****
###  GET /payment/:idpayment
Method is intended consult an especific payment

>URL: http://localhost:8080/payment/:idpayment

Where,  
**:idpayment** -> Payment's id to be consulted.  

Example:
>http://localhost:8080/payment/10  
This will return infos fo payment 10.

**Return**
```
{
    "return": {
        "state": 9
    },
    "payments": [
        {
            "client": {
                "id": 9
            },
            "buyer": {
                "name": "string",
                "email": "string",
                "cpfCnpj": "99999999999"
            },
            "payment": {
                "paymentId": 9,
                "amount": 99.9,
                "type": 9,
                "card": {
                    "holderName": "string",
                    "number": "9999",
                    "expirationDate": "01/01"
                },
                "boleto": {
                    "number": "99999999999999999999999999999999999999999999999"
                },
                "paymentState": 9
            }
        }
    ]
}
```
Where,   
**state** -> State to indicates success of opperaion. 0 = false, 1 = true.  
**client.id** -> Client's id.  
**buyer.name** -> Buyer's name.  
**buyer.email** -> Buyer's email.  
**buyer.cpfCnpj** -> Buyer's CPF or CNPJ.  
**payment.paymentId** -> Payment's id.  
**payment.amount** -> Payment's amount.  
**payment.type** -> Type of payment. 1 = boleto, 2 = credit card.  
**card.holderName** -> Name of card holder.  
**card.number** -> Card final number.Only returns last 4 digits.  
**card.expirationDate** -> Card's expiration date.  
**boleto.number** -> Boleto's number.  
**payment.paymentState** -> Payment's state. 1 - "Pending", 2 - "Paid", 3 - "Canceled", 4 - "Refused".  

*Notes:*  
*-card infos should only send when payment type was 2 (credit card).*  
*-For safety purpose, the card number only returns for last digits and the method doesn   't returns cvv.*  
*-boleto infos should only send when payment type was 1 (boleto).*  


****
###  GET /card/validate/:cardnumber
Method is intended to check if card number is valid.

>URL: http://localhost:8080/card/validate/:cardnumber

Where,  
**:cardnumber** -> Card's number.  

Example:
>http://localhost:8080/card/validate/9999999999999  
This will return if card's number "9999999999999" is a valid one .

**Return**
```
{
    "successful": true
}
```

Where,  
**successful** -> Indicates it's a valid card number. Returns true or false.  

****
###  GET /card/brand/:cardnumber
Method is intended to return de card's brand

>URL: http://localhost:8080/card/brand/:cardnumber

Where,  
**:cardnumber** -> Card's number.  

Example:
>http://localhost:8080/card/brand/9999999999999  
This will return the card's brand with number "9999999999999".

**Returns**
```
{
    "code": 9,
    "name": "string"
}
```

Where,  
**code** -> Brand code. *Brands supported: 1 - Amex, 2 - Diners, 3 - Elo, 4 - Hipercard, 5 - Hiper, 6 - Master, 7 - Visa*  
**name** -> Brand name.  

