package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"Client",
		"GET",
		"/client/{id}",
		ConsultClient,
	},
	Route{
		"Client",
		"GET",
		"/client",
		ConsultAllClients,
	},
	Route{
		"Client",
		"POST",
		"/client",
		InsertClient,
	},
	Route{
		"Payment",
		"POST",
		"/payment",
		DoPayment,
	},
	Route{
		"Payment",
		"PUT",
		"/payment/{idpayment}/state/{idstate}",
		UpdatePaymentState,
	},
	Route{
		"Payment",
		"GET",
		"/payment/{idpayment}",
		ConsultPaymentById,
	},
	Route{
		"Card",
		"GET",
		"/card/validate/{cardnumber}",
		ValidateCard,
	},
	Route{
		"Card",
		"GET",
		"/card/brand/{cardnumber}",
		ConsultCardBrand,
	},
}
