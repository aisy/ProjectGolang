package main

import (
	"encoding/json"
	"log"
	"net/http"
)

var users UsersFormat
var arrUser []UsersFormat
var response ResponseUsers

// get all
func returnAllUsers(w http.ResponseWriter, r *http.Request) {

	db := connectDatabase()
	defer db.Close()

	rows, err := db.Query("SELECT * from users")
	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&users.IDUSers, &users.NamaLengkap, &users.NoTelp); err != nil {
			log.Fatal(err.Error())

		} else {
			arrUser = append(arrUser, users)
		}
	}

	response.Status = 1
	response.Message = "Success"
	response.Data = arrUser

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

// insert
func insertUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var response ResponseUsers

	db := connectDatabase()
	defer db.Close()

	// err := r.ParseMultipartForm(4096) : code ini kalau gak mau make form data
	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	NamaLengkap := r.Form.Get("namalengkap")
	NoTelp := r.Form.Get("notlp")

	_, err = db.Exec("INSERT INTO users (nama_lengkap, no_tlp) values (?,?)",
		NamaLengkap,
		NoTelp,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Add"
	log.Print("Insert data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func updateUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var response ResponseUsers

	db := connectDatabase()
	defer db.Close()

	err := r.ParseForm()
	if err != nil {
		panic(err)
	}

	id := r.Form.Get("user_id")
	namalengkap := r.Form.Get("namalengkap")
	notTelp := r.Form.Get("notlp")

	_, err = db.Exec("UPDATE users set nama_lengkap = ?, no_tlp = ? where id_users = ?",
		namalengkap,
		notTelp,
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Update Data"
	log.Print("Update data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}

func deleteUsersMultipart(w http.ResponseWriter, r *http.Request) {
	// var response ResponseUsers
	db := connectDatabase()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	id := r.FormValue("user_id")

	_, err = db.Exec("DELETE from users where id_users = ?",
		id,
	)

	if err != nil {
		log.Print(err)
	}

	response.Status = 1
	response.Message = "Success Delete Data"
	log.Print("Delete data to database")

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)

}
