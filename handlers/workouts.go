package handlers

import (
	"db_backend/services"
	"db_backend/utils"
	"net/http"
)

func GetStrain(w http.ResponseWriter, r *http.Request) {
	trainer := r.FormValue("trainer")
	fromDate := r.FormValue("from_date")
	toDate := r.FormValue("to_date")

	data, err := services.GetStrainForTrainer(trainer, fromDate, toDate)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}
