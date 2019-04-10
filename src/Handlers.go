package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
)

func formatErrorResponse(w http.ResponseWriter, statusCode int, internalState int, message string, technicalMessage string) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)

	var ret ReturnStruct
	ret.State = internalState
	ret.Message = message
	ret.TechnicalMessage = technicalMessage

	if err := json.NewEncoder(w).Encode(ret); err != nil {
		http.Error(w, "Json couldn't be parsed. E: "+err.Error(), 500)
	}
}

func formatJsonReturn(w http.ResponseWriter, statusCode int, jsonObj interface{}) {
	jsonData, err := json.Marshal(jsonObj)
	if err != nil {
		// write your error to w, then return
		formatErrorResponse(w, 500, 500, "Couldn't parse json response.", "We've couldn't get response.")
	}
	w.WriteHeader(statusCode)
	w.Write(jsonData)
}

func formatSuccessfulResponse(w http.ResponseWriter, statusCode int) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(statusCode)
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func InsertClient(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	var client Client

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		formatErrorResponse(w, 500, 500, "Json request couldn't be read.", err.Error())
		return
	}

	if err := r.Body.Close(); err != nil {
		formatErrorResponse(w, 500, 500, "Request body couldn't be closed.", err.Error())
		return
	}
	if err := json.Unmarshal(body, &client); err != nil {
		formatErrorResponse(w, 422, 422, "Request should be a valid json.", err.Error())
		return
	}

	if (client.Name == "") && (client.Cpf == "") && (client.Email == "") {
		formatErrorResponse(w, 500, 500, "You should inform name, cpf or email to save client.", "")
		return
	}

	var retClients Clients
	retClient, err, customError := InsertClientBD(client)
	if err != nil {
		//avoid empty technical message
		tmessage := customError.TechnicalMessage
		if tmessage != "" {
			tmessage = err.Error()
		}
		formatErrorResponse(w, 500, customError.IdMessage, customError.Message, tmessage)
		return
	}
	retClients = append(retClients, retClient)

	formatJsonReturn(w, http.StatusCreated, retClients)

}

func ConsultAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	formatJsonReturn(w, http.StatusOK, FormatClientConsult(ConsultAllClientsDB(), 1, ""))
}

func ConsultClient(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {

		formatErrorResponse(w, 422, 422, "Id to be consulted not found.", err.Error())
		return
	}

	formatJsonReturn(w, http.StatusOK, FormatClientConsult(ConsultClientDB(id), 1, ""))
}

func DoPayment(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	var payInfo Payment

	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		formatErrorResponse(w, 500, 500, "Json request couldn't be read.", err.Error())
		return
	}
	if err := r.Body.Close(); err != nil {
		formatErrorResponse(w, 500, 500, "Request body couldn't be closed.", err.Error())
		return
	}
	if err := json.Unmarshal(body, &payInfo); err != nil {
		formatErrorResponse(w, 422, 422, "Request should be a valid json.", err.Error())
		return
	}

	//var retPayment PaymentReturn

	retPayment := PaymentMethod(payInfo)

	formatJsonReturn(w, http.StatusCreated, retPayment)

}

func UpdatePaymentState(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	//var ret ReturnClients
	vars := mux.Vars(r)
	idpayment, err := strconv.ParseInt(vars["idpayment"], 10, 64)
	if err != nil {

		formatErrorResponse(w, 422, 422, "Id from payment to be altered not found.", err.Error())
		return
	}
	idstate, err := strconv.Atoi(vars["idstate"])
	if err != nil {

		formatErrorResponse(w, 422, 422, "Id from state to be altered not found.", err.Error())
		return
	}

	formatJsonReturn(w, http.StatusOK, AlterPaymentState(idpayment, idstate))
}

func ClientPaymen(w http.ResponseWriter, r *http.Request) {
	param1 := r.URL.Query().Get("idclient")
	//param1s := r.URL.Query()["idclient"]
	formatErrorResponse(w, 422, 422, param1, "")

}

func ConsultPaymentById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	idpayment, err := strconv.ParseInt(vars["idpayment"], 10, 64)
	if err != nil {
		formatErrorResponse(w, 422, 422, "Id from payment to be consulted not found.", err.Error())
		return
	}

	payments := ConsultPaymentID(idpayment)

	formatJsonReturn(w, http.StatusOK, payments)
}

func ValidateCard(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	cardNumber := vars["cardnumber"] //strconv.ParseInt(vars["cardnumber"], 10, 64)
	/*
		if err != nil {
			formatErrorResponse(w, 422, 422, "Card number to be validated not found.", err.Error())
			return
		}*/

	var cardReturn CardReturn
	cardReturn.Successful = ValidCardNumber(cardNumber)

	formatJsonReturn(w, http.StatusOK, cardReturn)

}

func ConsultCardBrand(w http.ResponseWriter, r *http.Request) {
	//func GetCarBrand(cardNumber string) BrandData
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	vars := mux.Vars(r)
	cardNumber := vars["cardnumber"] //strconv.ParseInt(vars["cardnumber"], 10, 64)
	/*
		if err != nil {
			formatErrorResponse(w, 422, 422, "Card number to be validated not found.", err.Error())
			return
		}
	*/

	brand := GetCarBrand(cardNumber)

	formatJsonReturn(w, http.StatusOK, brand)
}
