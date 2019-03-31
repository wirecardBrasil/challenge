package main

import (
	"database/sql"
	"encoding/json"
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
			") ", client.Name, client.Email, client.Cpf)
	if err != nil {
		/*
			me, ok := err.(*mysql.MySQLError)
			if !ok {
				return err
			}
			if me.Number == 1062 {
				return errors.New("It already exists in a database.")
			}
			return err*/

		if driverErr, ok := err.(*mysql.MySQLError); ok { // Now the error number is accessible directly
			if driverErr.Number == 1062 {

				cError.Message = err.Error()
				cError.TechnicalMessage = "Client already registered."

				//id, err := strconv.ParseInt(vars["id"], 10, 64)
				cError.IdMessage = int(1062)
			} else {
				cError.Message = err.Error()
				cError.TechnicalMessage = "DB fail registering client."
				cError.IdMessage = int(driverErr.Number)
			}
		} else {
			cError.Message = err.Error()
			cError.TechnicalMessage = "Fail registering client."
			cError.IdMessage = int(driverErr.Number)

		}
		return clientReturn, err, cError
	}

	id, err := qryInsert.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	clientReturn.Id = id
	clientReturn.Name = client.Name
	clientReturn.Email = client.Email
	clientReturn.Cpf = client.Cpf

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
		err = qryConsult.Scan(&client.Id, &client.Name, &client.Email, &client.Cpf)
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
		"	regCLient "+
		"WHERE "+
		"	id = ? "+
		"ORDER BY "+
		"	id desc ", id)
	if err != nil {
		panic(err.Error())
	}

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
