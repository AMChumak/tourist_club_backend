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
	defer r.Body.Close()
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

func GetPersonRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	section := r.FormValue("section")

	role, err := services.GetPersonRole(person, section)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	utils.RespondWithJSON(w, http.StatusOK, role)
}

func SetPersonRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	section := r.FormValue("section")
	role := r.FormValue("role")
	err := services.SetPersonRole(person, section, role)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	section := r.FormValue("section")

	err := services.DeletePersonRole(person, section)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func CreatePersonAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.PersonAttribute
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding person create request:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	err := services.CreatePersonAttribute(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func GetPersonAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")

	attr, err := services.GetPersonAttribute(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, attr)
}

func SetPersonAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.PersonAttribute
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding person create request:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
	}
	err := services.SetPersonAttribute(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	attr := r.FormValue("id")
	err := services.DeletePersonAttribute(attr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func FindTourists(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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
	defer r.Body.Close()
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
	defer r.Body.Close()
	birthYear := r.FormValue("birth_year")
	beginYear := r.FormValue("begin_year")
	age := r.FormValue("age")
	salary := r.FormValue("salary")
	sex := r.FormValue("sex")

	data, err := services.GetManagersWithCondition(salary, birthYear, age, beginYear, sex)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTrainersByWorkouts(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
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

func GetAllRoles(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	roles, err := services.GetAllRoles()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, roles)
}

func GetAllPersonAttributes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	attributes, err := services.GetAllAttributes()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, attributes)
}

func GetPersonIntAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	val, err := services.GetPersonIntAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, val)
}

func GetPersonFloatAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	val, err := services.GetPersonFloatAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, val)
}

func GetPersonStringAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	val, err := services.GetPersonStringAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, val)
}

func GetPersonDateAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	val, err := services.GetPersonDateAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, val)
}

func SetPersonIntAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var attr dto.PersonIntAttribute
	if err := json.NewDecoder(r.Body).Decode(&attr); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.SetPersonIntAttribute(attr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func SetPersonFloatAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var attr dto.PersonFloatAttribute
	if err := json.NewDecoder(r.Body).Decode(&attr); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.SetPersonFloatAttribute(attr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func SetPersonStringAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var attr dto.PersonStringAttribute
	if err := json.NewDecoder(r.Body).Decode(&attr); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.SetPersonStringAttribute(attr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func SetPersonDateAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var attr dto.PersonDateAttribute
	if err := json.NewDecoder(r.Body).Decode(&attr); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.SetPersonDateAttribute(attr)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonIntAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	err := services.DeletePersonIntAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonFloatAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	err := services.DeletePersonFloatAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonStringAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	err := services.DeletePersonStringAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeletePersonDateAttribute(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	attribute := r.FormValue("attribute")

	err := services.DeletePersonDateAttribute(person, attribute)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}
