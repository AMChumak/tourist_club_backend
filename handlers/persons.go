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

func FindTourists(w http.ResponseWriter, r *http.Request) {
	section := r.FormValue("section")
	group := r.FormValue("group")
	sex := r.FormValue("sex")
	birthYear := r.FormValue("birth_year")
	age := r.FormValue("age")

	data, err := services.GetTouristsWithCondition(section, group, sex, birthYear, age)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTrainers(w http.ResponseWriter, r *http.Request) {
	section := r.FormValue("section")
	sex := r.FormValue("sex")
	age := r.FormValue("age")
	salary := r.FormValue("salary")
	specialization := r.FormValue("specialization")

	data, err := services.GetTrainersWithCondition(section, sex, age, salary, specialization)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindManagers(w http.ResponseWriter, r *http.Request) {
	birthYear := r.FormValue("birth_year")
	beginYear := r.FormValue("begin_year")
	age := r.FormValue("age")
	salary := r.FormValue("salary")

	data, err := services.GetManagersWithCondition(salary, birthYear, age, beginYear)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTrainersByWorkouts(w http.ResponseWriter, r *http.Request) {
	group := r.FormValue("group")
	from := r.FormValue("from_date")
	to := r.FormValue("to_date")
	data, err := services.GetTrainersByWorkout(group, from, to)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}
