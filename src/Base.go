package main

import (
	"database/sql"
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"os"
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

func InsertClientBD(client Client) Client {
	db := dbConn()
	qryInsert, err := db.Exec(
		"INSERT INTO regClient( "+
			"	clientName, "+
			"	email,  "+
			"	cpfCnpj "+
			") VALUES ( "+
			"	?, "+
			"	?, "+
			"	? "+
			") ", client.Name, client.Email, client.Cpf)
	if err != nil {
		panic("565656" + err.Error())
	}

	id, err := qryInsert.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	clientReturn := Client{}
	clientReturn.Id = id
	clientReturn.Name = client.Name
	clientReturn.Email = client.Email
	clientReturn.Cpf = client.Cpf

	defer db.Close()
	return clientReturn
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
		err = qryConsult.Scan(&client.Id, &client.Name, &client.Email, &client.Cpf)
		if err != nil {
			panic(err.Error())
		}
		clients = append(clients, client)
	}
	defer db.Close()
	return clients
}
