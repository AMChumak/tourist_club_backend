package handlers

import (
	"db_backend/dto"
	"db_backend/services"
	"db_backend/utils"
	"encoding/json"
	"net/http"
)

func FindTouristsByTour(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	section := r.FormValue("section")
	group := r.FormValue("group")
	cntTours := r.FormValue("cnt_tours")
	tourId := r.FormValue("tour_id")
	tourTime := r.FormValue("tour_time")
	routeId := r.FormValue("route_id")
	placeId := r.FormValue("place_id")

	data, err := services.GetTouristsByTour(section, group, cntTours, tourId, tourTime, routeId, placeId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindRoutes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	section := r.FormValue("section")
	dateFrom := r.FormValue("date_from")
	dateTo := r.FormValue("date_to")
	instructor := r.FormValue("instructor")
	groupCnt := r.FormValue("group_cnt")

	data, err := services.GetRoutesWithConditions(section, dateFrom, dateTo, instructor, groupCnt)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindRoutesWithGeo(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	place := r.FormValue("place")
	length := r.FormValue("length")
	difficulty := r.FormValue("difficulty")

	data, err := services.GetRoutesWithGeoCond(place, length, difficulty)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindInstructors(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	role := r.FormValue("role")
	routeType := r.FormValue("type")
	routeDifficulty := r.FormValue("difficulty")
	cntTours := r.FormValue("cnt_tours")
	tourId := r.FormValue("tour_id")
	placeId := r.FormValue("place_id")

	data, err := services.GetInstructorsWithCondition(role, routeType, routeDifficulty, cntTours, tourId, placeId)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTouristsWithTrainerInstructor(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	section := r.FormValue("section")
	group := r.FormValue("group")

	data, err := services.GetTouristsWithTrainerInstructor(section, group)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTouristsCompletedAll(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	section := r.FormValue("section")
	group := r.FormValue("group")

	data, err := services.GetTouristsCompletedALl(section, group)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func FindTouristsCompletedRoutes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	section := r.FormValue("section")
	group := r.FormValue("group")

	var requestBody = dto.CompletedRoutesRequest{}
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Invalid JSON: "+err.Error(), http.StatusBadRequest)
		return
	}

	data, err := services.GetTouristsCompletedRoutes(section, group, requestBody)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func GetAllRouteTypes(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	data, err := services.GetAllRouteTypes()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}

func GetTouristsByTour(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	routeType := r.FormValue("type_id")

	difficulty := r.FormValue("difficulty")

	data, err := services.GetSuitablePersonsByRoute(routeType, difficulty)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, data)
}
