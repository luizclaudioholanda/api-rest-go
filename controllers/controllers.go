package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/luizclaudioholanda/go-rest-api/database"
	"github.com/luizclaudioholanda/go-rest-api/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func ListaPersonalidades(w http.ResponseWriter, r *http.Request) {

	var listaPersonalidades []models.Personalidade
	database.DB.Find(&listaPersonalidades)

	json.NewEncoder(w).Encode(listaPersonalidades)
}

func RetornaPersonalidadePorId(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade

	database.DB.First(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func CriaPersonalidade(w http.ResponseWriter, r *http.Request) {

	var personalidade models.Personalidade
	json.NewDecoder(r.Body).Decode(&personalidade)

	database.DB.Create(&personalidade)

	json.NewEncoder(w).Encode(personalidade)
}

func DeletaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var personalidade models.Personalidade
	database.DB.Delete(&personalidade, id)

	json.NewEncoder(w).Encode(personalidade)
}

func EditaPersonalidade(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	id := vars["id"]

	var personalidadeDB models.Personalidade

	database.DB.First(&personalidadeDB, id)
	json.NewDecoder(r.Body).Decode(&personalidadeDB)
	database.DB.Save(&personalidadeDB)

	json.NewEncoder(w).Encode(personalidadeDB)

}
