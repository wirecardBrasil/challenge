package main

import (
	"database/sql"
	"encoding/json"
	//"fmt"
	"github.com/go-sql-driver/mysql"
	"os"
	//"strconv"
)

func dbConn() (db *sql.DB) {

	file, err := os.Open("config/config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := DataConfiguration{}
	err = decoder.Decode(&configuration)
	if err != nil {
		panic(err.Error())
	}

	dbDriver := "mysql"
	dbUser := configuration.Database.User
	dbPass := configuration.Database.Password
	dbName := configuration.Database.Namedb
	db, err = sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func InsertClientBD(client Client) (Client, error, CustomError) {
	db := dbConn()
	var cError CustomError

	clientReturn := Client{}

	qryInsert, err := db.Exec(
		"INSERT INTO regClient( "+
			"	clientName, "+
			"	email,  "+
			"	cpfCnpj "+
			") VALUES ( "+
			"	?, "+
			"	?, "+
			"	? "+
			") ", client.Name, client.Email, client.CpfCnpj)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			//1062 = SQL code to duplicated key.
			//check duplicated register.
			if driverErr.Number == 1062 {

				cError.Message = "Client already registered."
				cError.TechnicalMessage = err.Error()
				cError.IdMessage = int(1062)
			} else {
				cError.Message = "DB fail registering client."
				cError.TechnicalMessage = err.Error()
				cError.IdMessage = int(driverErr.Number)
			}
		} else {
			cError.Message = "Fail registering client."
			cError.TechnicalMessage = err.Error()
			cError.IdMessage = int(driverErr.Number)

		}
		return clientReturn, err, cError
	}

	id, err := qryInsert.LastInsertId()
	if err != nil {

		cError.Message = "Could'nt get inserted id."
		cError.TechnicalMessage = err.Error()
		cError.IdMessage = int(1230)

		return clientReturn, err, cError
	}

	clientReturn.Id = id
	clientReturn.Name = client.Name
	clientReturn.Email = client.Email
	clientReturn.CpfCnpj = client.CpfCnpj

	defer db.Close()
	return clientReturn, nil, cError
}

func ConsultAllClientsDB() Clients {
	db := dbConn()
	qryConsult, err := db.Query("SELECT " +
		"	id, " +
		"	clientName, " +
		"	email, " +
		"	cpfCnpj " +
		"FROM " +
		"	regCLient " +
		"ORDER BY " +
		"	id desc ")
	if err != nil {
		panic(err.Error())
	}

	client := Client{}
	clients := []Client{}
	for qryConsult.Next() {
		err = qryConsult.Scan(&client.Id, &client.Name, &client.Email, &client.CpfCnpj)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, client)
	}
	defer db.Close()
	return clients
}

func ConsultClientDB(id int64) Clients {
	db := dbConn()
	client := Client{}
	clients := []Client{}
	qryConsult, err := db.Query("SELECT "+
		"	id, "+
		"	clientName, "+
		"	email, "+
		"	cpfCnpj "+
		"FROM "+
		"	regClient "+
		"WHERE "+
		"	id = ? "+
		"ORDER BY "+
		"	id desc ", id)
	if err != nil {
		panic(err.Error())
	}

	for qryConsult.Next() {
		err = qryConsult.Scan(&client.Id, &client.Name, &client.Email, &client.CpfCnpj)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, client)
	}

	defer db.Close()
	return clients

}

func ClientRegistered(id int64) bool {
	db := dbConn()
	var count int
	err := db.QueryRow("SELECT "+
		"	COUNT(id) "+
		"FROM "+
		"	regClient "+
		"WHERE "+
		"	id = ? "+
		"ORDER BY "+
		"	id desc ", id).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
	return count > 0
}

func BuyerRegistered(cpfCnpj string) bool {
	db := dbConn()
	var count int
	err := db.QueryRow("SELECT "+
		"	COUNT(id) "+
		"FROM "+
		"	regBuyer "+
		"WHERE "+
		"	cpfCnpj = ? "+
		"ORDER BY "+
		"	id desc ", cpfCnpj).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	//panic("QUantidade: " + strconv.Itoa(count) + cpfCnpj)

	defer db.Close()
	return count > 0
}

func BuyerInfo(cpfCnpj string) (Buyer, error) {
	db := dbConn()
	var buyer Buyer
	err := db.QueryRow("SELECT "+
		"	id, "+
		"	buyerName, "+
		"	email, "+
		"	cpfCnpj "+
		"FROM "+
		"	regBuyer "+
		"WHERE "+
		"	cpfCnpj = ? "+
		"ORDER BY "+
		"	id desc "+
		"LIMIT 1 ", cpfCnpj).Scan(&buyer.Id, &buyer.Name, &buyer.Email, &buyer.CpfCnpj)
	if err != nil {
		return buyer, err
	}
	defer db.Close()
	//panic("Buyer ID " + strconv.FormatInt(buyer.Id, 10))
	return buyer, nil
}

func SaveBuyer(buyer Buyer) (Buyer, error, CustomError) {

	var cError CustomError

	buyerReturn := Buyer{}

	db := dbConn()
	qryInsert, err := db.Exec(
		"INSERT INTO regBuyer( "+
			"	BuyerName, "+
			"	email,  "+
			"	cpfCnpj "+
			") VALUES ( "+
			"	?, "+
			"	?, "+
			"	? "+
			") ", buyer.Name, buyer.Email, buyer.CpfCnpj)
	if err != nil {
		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			//1062 = SQL code to duplicated key.
			//check duplicated register.
			if driverErr.Number == 1062 {

				cError.Message = "Buyer already registered."
				cError.TechnicalMessage = err.Error()
				cError.IdMessage = int(1062)
			} else {
				cError.Message = "DB fail registering buyer."
				cError.TechnicalMessage = err.Error()
				cError.IdMessage = int(driverErr.Number)
			}
		} else {
			cError.Message = "Fail registering buyer."
			cError.TechnicalMessage = err.Error()
			cError.IdMessage = int(driverErr.Number)

		}
		return buyerReturn, err, cError
	}

	id, err := qryInsert.LastInsertId()
	if err != nil {

		cError.Message = "Could'nt get inserted id."
		cError.TechnicalMessage = err.Error()
		cError.IdMessage = int(1230)

		return buyerReturn, err, cError
	}

	buyerReturn.Id = id
	buyerReturn.Name = buyer.Name
	buyerReturn.Email = buyer.Email
	buyerReturn.CpfCnpj = buyer.CpfCnpj

	defer db.Close()
	return buyerReturn, nil, cError
}

func SavePayment(payment Payment) (Payment, error) {
	db := dbConn()
	var paymentRet = Payment{}
	paymentState := 1

	qryInsert, err := db.Exec(
		"INSERT INTO payment "+
			"( "+
			"    idClient, "+
			"    idBuyer, "+
			"    idPaymentType, "+
			"    amount, "+
			"    idPaymentState "+
			") VALUES ( "+
			"	?, "+
			"	?, "+
			"	?, "+
			"	?, "+
			"	? "+
			") ", payment.Client.Id, payment.Buyer.Id, payment.PaymentInfo.PaymentType, payment.PaymentInfo.Amount, paymentState)

	//	panic("Inicio:" + strconv.FormatInt(payment.Client.Id, 10) + " " + strconv.FormatInt(payment.Buyer.Id, 10) + " " + strconv.Itoa(payment.PaymentInfo.PaymentType) + " " + fmt.Sprintf("%f", payment.PaymentInfo.Amount))
	if err != nil {
		return paymentRet, err
	}

	id, err := qryInsert.LastInsertId()
	if err != nil {
		return paymentRet, err
	}

	paymentRet = payment
	paymentRet.PaymentInfo.PaymentID = id
	paymentRet.PaymentInfo.PaymentState = paymentState
	defer db.Close()
	return paymentRet, nil
}

func SaveCardPayment(payment Payment) (Payment, error) {
	var paymentRet = Payment{}
	db := dbConn()

	//stores only last 4 characters of card number
	payment.PaymentInfo.Card.Number = string(payment.PaymentInfo.Card.Number[len(payment.PaymentInfo.Card.Number)-4:])

	_, err := db.Exec(
		"INSERT INTO cardPayment( "+
			"    idPayment, "+
			"    holderName, "+
			"    cardFinalNumber, "+
			"    expirationDate "+
			") VALUES( "+
			"	?, "+
			"    ?, "+
			"    ?, "+
			"    STR_TO_DATE(? , '%m/%y') "+
			") ", payment.PaymentInfo.PaymentID, payment.PaymentInfo.Card.HolderName, payment.PaymentInfo.Card.Number, payment.PaymentInfo.Card.ExpirationDate)
	if err != nil {
		return paymentRet, err
	}

	/*
		id, err := qryInsert.LastInsertId()
		if err != nil {
			return paymentRet, err
		}
	*/
	paymentRet = payment
	defer db.Close()
	return paymentRet, nil
}

func SaveBoletoPayment(payment Payment) (Payment, error) {
	var paymentRet = Payment{}
	db := dbConn()

	_, err := db.Exec(
		"INSERT INTO boletoPayment( "+
			"    idPayment, "+
			"    boletoNumber "+
			") VALUES( "+
			"	?, "+
			"    ? "+
			") ", payment.PaymentInfo.PaymentID, payment.PaymentInfo.Boleto.Number)
	if err != nil {

		return paymentRet, err
	}

	paymentRet = payment
	defer db.Close()
	return paymentRet, nil
}

func AlterPaymentStateDB(idPayment int64, idState int) (bool, string) {
	db := dbConn()

	qryUpdate, err := db.Exec(
		"UPDATE "+
			"	payment "+
			"SET "+
			"	idPaymentState = ? "+
			"WHERE "+
			"	id = ? ", idState, idPayment)
	if err != nil {
		return false, err.Error()
	}

	lines, err := qryUpdate.RowsAffected()
	if err != nil {
		return false, err.Error()
	}
	if lines < 1 {
		return false, "Couldn't update register. Check the payment id and state id."
	}
	return true, ""

}

/*

func PaymentConsultReturn(rows *rows) PaymentConsult {
	for rows.Next() {
		err = rows.Scan(&client.Id, &client.Name, &client.Email, &client.CpfCnpj)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, client)
	}
}
*/

func ConsultPaymentByIdBD(idPayment int64) (Payments, error) {
	db := dbConn()
	payment := Payment{}
	payments := []Payment{}

	qryConsult, err := db.Query("SELECT  "+
		"	 payment.id, "+
		"    payment.idClient, "+
		"    regBuyer.buyerName, "+
		"    regBuyer.email, "+
		"    regBuyer.CpfCnpj, "+
		"    payment.amount, "+
		"	 IFNULL(cardPayment.holderName, '') as holderName, "+
		"	 IFNULL(cardPayment.cardFinalNumber, '') as cardFinalNumber, "+
		"    IFNULL(DATE_FORMAT(cardPayment.expirationDate, '%m/%y'), '') as expirationDate, "+
		"    idPaymentType, "+
		"    idPaymentState, "+
		"    IFNULL(boletoPayment.boletoNumber, '') as boletoNumber "+
		"FROM "+
		"	payment "+
		"	LEFT JOIN regBuyer "+
		"		on payment.idBuyer = regBuyer.id "+
		"	LEFT JOIN cardPayment "+
		"		on cardPayment.idPayment = payment.id "+
		"	LEFT JOIN boletoPayment "+
		"		on boletoPayment.idPayment = payment.id "+
		"WHERE  "+
		"	payment.id = ? ", idPayment)
	if err != nil {
		return payments, err
	}

	for qryConsult.Next() {
		err = qryConsult.Scan(&payment.PaymentInfo.PaymentID, &payment.Client.Id, &payment.Buyer.Name, &payment.Buyer.Email, &payment.Buyer.CpfCnpj,
			&payment.PaymentInfo.Amount, &payment.PaymentInfo.Card.HolderName, &payment.PaymentInfo.Card.Number, &payment.PaymentInfo.Card.ExpirationDate,
			&payment.PaymentInfo.PaymentType, &payment.PaymentInfo.PaymentState, &payment.PaymentInfo.Boleto.Number)
		if err != nil {
			return payments, err
		}
		payments = append(payments, payment)
	}

	defer db.Close()
	return payments, nil

}
