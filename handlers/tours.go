package handlers

import (
	"db_backend/services"
	"db_backend/utils"
	"net/http"
)

func FindTouristsByTour(w http.ResponseWriter, r *http.Request) {
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
