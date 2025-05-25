package handlers

import (
	"db_backend/dto"
	"db_backend/services"
	"db_backend/utils"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
)

func CreateGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.Group
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		log.Println("Error decoding group in create request:", err)
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := services.CreateGroup(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"id": strconv.Itoa(id)})
}

func GetGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	group, err := services.GetGroup(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	var groupResponse = dto.Group{}

	if group == nil {
		utils.RespondWithJSON(w, http.StatusOK, nil)
		return
	}
	groupResponse.Id = group.Id
	groupResponse.GroupNumber = group.GroupNumber
	groupResponse.Section = group.Section
	utils.RespondWithJSON(w, http.StatusOK, groupResponse)
}

func UpdateGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var group dto.Group
	if err := json.NewDecoder(r.Body).Decode(&group); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.UpdateGroup(group)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeleteGroup(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	err := services.DeleteGroup(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func GetGroupMembers(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	members, err := services.GetGroupMembers(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, members)
}

func AddGroupMember(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	group := r.FormValue("group")

	err := services.AddGroupMember(group, person)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func RemoveGroupMember(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	person := r.FormValue("person")
	group := r.FormValue("group")
	err := services.RemoveGroupMember(person, group)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func GetAllGroups(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	groups, err := services.GetAllGroups()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, groups)
}

func CreateSection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var req dto.Section
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	id, err := services.CreateSection(req)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"id": strconv.Itoa(id)})
}

func GetSection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	section, err := services.GetSection(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, section)
}

func UpdateSection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var section dto.Section
	if err := json.NewDecoder(r.Body).Decode(&section); err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	err := services.UpdateSection(section)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func DeleteSection(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	err := services.DeleteSection(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"status": "success"})
}

func GetAllSections(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	sections, err := services.GetAllSections()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, sections)
}

func GetGroupsFromSections(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	id := r.FormValue("id")
	groups, err := services.GetGroupFromSection(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, groups)
}
