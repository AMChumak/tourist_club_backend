package handlers

import (
	"db_backend/services"
	"db_backend/utils"
	"net/http"
)

func FindChampionships(w http.ResponseWriter, r *http.Request) {
	section := r.FormValue("section")

	data, err := services.GetChampionshipsWithCondition(section)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}
