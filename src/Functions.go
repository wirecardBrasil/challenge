package main

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
