package main

/*
import (
	"strconv"
)
*/

func FormatClientConsult(cl Clients, state int, msg string) ReturnClients {
	var retClient ReturnClients
	if cl == nil {
		retClient.Return.State = 0
		retClient.Return.Message = "Client(s) not found."
	} else {
		retClient.Return.State = state
		retClient.Return.Message = msg
		retClient.Clients = cl
	}
	return retClient
}

func BoletoPayment() string {
	//Mocked value
	const boletoNumber = "23790504004199033014836008109203478470000019900"
	return boletoNumber
}

func CardPayment() bool {
	//Mocked value
	const successful = true
	return successful
}

func validPaymentType(paymentType int) bool {
	return (paymentType == 1 || paymentType == 2)
}

func PaymentMethod(payInfo2 Payment) PaymentReturn {

	payInfo := payInfo2
	var payReturn = PaymentReturn{}

	//Check if client exists
	if !(ClientRegistered(payInfo.Client.Id)) {
		payReturn.Return.State = 0
		payReturn.Return.Message = "We've coundn't found this client."
		payReturn.Return.TechnicalMessage = "Client not found."
		return payReturn
	}

	//var retBuyer = Buyer{}
	var retBuyer = Buyer{}
	//var customError = CustomError{}
	//Check if buyer exists. If doesn't exists, save
	if BuyerRegistered(payInfo.Buyer.Cpf) {
		localRetBuyer, err := BuyerInfo(payInfo.Buyer.Cpf)

		if err != nil {
			payReturn.Return.State = 0
			payReturn.Return.TechnicalMessage = err.Error()
			return payReturn
		}

		retBuyer = localRetBuyer
	} else {

		//panic("RET3 " + strconv.FormatInt(retBuyer.Id, 10))
		localRetBuyer, err, customError := SaveBuyer(payInfo.Buyer)
		if err != nil {
			//avoid empty technical message
			tmessage := customError.TechnicalMessage
			if tmessage != "" {
				tmessage = err.Error()
			}
			payReturn.Return.State = customError.IdMessage
			payReturn.Return.Message = customError.Message
			payReturn.Return.TechnicalMessage = tmessage
			return payReturn
		}
		retBuyer = localRetBuyer
	}

	//panic("RET4 " + strconv.FormatInt(retBuyer.Id, 10))

	//panic("RET2 " + strconv.FormatInt(retBuyer.Id, 10))
	payInfo.Buyer = retBuyer
	//panic("ALO" + payInfo.Buyer.Cpf + " = " + retBuyer.Cpf)
	//Save general payment info
	payInfo, err := SavePayment(payInfo)
	if err != nil {
		payReturn.Return.State = 0
		payReturn.Return.Message = "Coudln't save payment."
		payReturn.Return.TechnicalMessage = err.Error()
		return payReturn
	}

	//Boleto
	if payInfo.PaymentInfo.PaymentType == 1 {
		//Save at database
		payReturn.Boleto.Number = BoletoPayment()

	} else {
		if payInfo.PaymentInfo.PaymentType == 2 {
			payInfo, err = SaveCardPayment(payInfo)
			if err != nil {
				payReturn.Return.State = 0
				payReturn.Return.Message = "Coudln't save card payment."
				payReturn.Return.TechnicalMessage = err.Error()
				return payReturn
			}

		}
		payReturn.Card.Successful = CardPayment()
	}
	payReturn.Payment.PaymentId = payInfo.PaymentInfo.PaymentID
	payReturn.Payment.StateId = payInfo.PaymentInfo.PaymentState

	payReturn.Return.State = 1
	payReturn.Return.Message = "Ok"
	return payReturn

}
