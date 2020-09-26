package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"github.com/gorilla/mux"
)

//Contato is representing of a person
type Contato struct {
	ID       int    `json:",id"`
	Nome     string `json:",nome"`
	Telefone string `json:",telefone"`
	Email    string `json:",email"`
}

var contatos []Contato

func main() {

	rota := mux.NewRouter()

	contatos = append(contatos,
		Contato{ID: 1, Nome: "Felipe Costa", Telefone: "555555555", Email: "costadefelipe@gmail.com"},
		Contato{ID: 2, Nome: "Henrique", Telefone: "555555555", Email: "gomeslinda@gmail.com"},
		Contato{ID: 3, Nome: "Ruth", Telefone: "549984958594", Email: "diamanteruth@gmai.com"})

	rota.HandleFunc("/contatos", getContatos).Methods("GET")
	rota.HandleFunc("/contato/{id}", getContato).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", rota))

}

func getContatos(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(contatos)
}

func getContato(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])

	if err == nil {
		for _, contato := range contatos {
			if contato.ID == id {
				json.NewEncoder(w).Encode(contato)
			}
		}
	}
}
