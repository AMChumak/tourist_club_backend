package handlers

import (
	"db_backend/dto"
	"db_backend/services"
	"db_backend/utils"
	"encoding/json"
	"log"
	"net/http"
)

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var req dto.PersonCreateRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding person create request:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := services.CreatePerson(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
