package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	//"github.com/gorilla/mux"
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

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func InsertClient(w http.ResponseWriter, r *http.Request) {
	var client Client
	var ret ReturnClient
	//var retClient Client

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

	retClient := InsertClientBD(client)
	ret.Return.State = 1
	ret.Return.Message = "Client(s) inserted successfully"
	ret.Client = retClient

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(ret); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}

}

func ConsultAllClients(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	if err := json.NewEncoder(w).Encode(ConsultAllClientsDB()); err != nil {
		formatErrorResponse(w, 500, 500, "Response couldn't be parsed.", err.Error())
	}
	w.WriteHeader(http.StatusOK)
}

func hello() string {
	return "Welcome!"

}

func validatePaymentType(pType int) bool {
	if (pType == 1) || (pType == 2) {
		return true
	} else {
		return false
	}
}
