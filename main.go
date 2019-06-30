package main

import (
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type User struct {
	gorm.Model
	Name  string
	Email string
}

func allUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "usuario:contraseña@/nombrebasededatos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("No se puede conectar con la base de datos.")
	}
	defer db.Close()

	var users []User
	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func newUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "usuario:contraseña@/nombrebasededatos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("No se puede conectar con la base de datos.")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	db.Create(&User{Name: name, Email: email})
	fmt.Fprintf(w, "Nuevo usuario creado correctamente.")
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "usuario:contraseña@/nombrebasededatos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("No se puede conectar con la base de datos.")
	}
	defer db.Close()

	vars := mux.Vars(r)
	email := vars["email"]

	var user User
	db.Where("email = ?", email).Find(&user)
	db.Delete(&user)

	fmt.Fprintf(w, "Usuario eliminado correctamente.")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("mysql", "usuario:contraseña@/nombrebasededatos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("No se puede conectar con la base de datos.")
	}
	defer db.Close()

	vars := mux.Vars(r)
	name := vars["name"]
	email := vars["email"]

	var user User
	db.Where("email = ?", email).Find(&user)

	user.Name = name

	db.Save(&user)
	fmt.Fprintf(w, "Usuario actualizado correctamente.")
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", allUsers).Methods("GET")
	myRouter.HandleFunc("/user/{name}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/{name}/{email}", updateUser).Methods("PUT")
	myRouter.HandleFunc("/user/{name}/{email}", newUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":8081", myRouter))
}


func initialMigration() {
	db, err := gorm.Open("mysql", "usuario:contraseña@/nombrebasededatos?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
		panic("No se puede conectar con la base de datos.")
	}
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&User{})
}

func main() {
	fmt.Println("Tutorial Go ORM ")

	initialMigration()

	handleRequests()
}