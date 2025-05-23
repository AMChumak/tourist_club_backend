package main

import (
	"db_backend/db"
	"db_backend/handlers"
	"flag"
	"fmt"
	gorillahandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
)

var (
	Logger     = log.New(os.Stdout, "Server:\t", log.LstdFlags)
	listenPort string
)

func main() {
	//flags
	flag.StringVar(&listenPort, "port", "8080", "server's port")
	flag.StringVar(&db.ConnString, "conn", "postgres://", "connection string to postgres")
	flag.Parse()

	//tests

	r := mux.NewRouter()

	corsOptions := []gorillahandlers.CORSOption{
		gorillahandlers.AllowedOrigins([]string{"*", "null"}), // Добавьте "null"
		gorillahandlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		gorillahandlers.AllowedHeaders([]string{"Content-Type", "Authorization", "From"}),
		gorillahandlers.AllowCredentials(),
	}
	h := gorillahandlers.CORS(corsOptions...)(r)

	r.Use(func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			Logger.Printf("%s %s %s", r.RemoteAddr, r.Method, r.URL.Path)
			next.ServeHTTP(w, r)
		})
	})

	r.HandleFunc("/persons/create", handlers.CreatePerson).Methods("POST")
	r.HandleFunc("/tourists/filter", handlers.FindTourists).Methods("GET")
	r.HandleFunc("/trainers/filter", handlers.FindTrainers).Methods("GET")
	r.HandleFunc("/managers/filter", handlers.FindManagers).Methods("GET")
	r.HandleFunc("/championships/filter", handlers.FindChampionships).Methods("GET")
	r.HandleFunc("/trainers/workout-filter", handlers.FindTrainersByWorkouts).Methods("GET")
	r.HandleFunc("/workouts/strain", handlers.GetStrain).Methods("GET")
	r.HandleFunc("/tourists/tour-filter", handlers.FindTouristsByTour).Methods("GET")
	r.HandleFunc("/routes/filter", handlers.FindRoutes).Methods("GET")
	r.HandleFunc("/routes/geofilter", handlers.FindRoutesWithGeo).Methods("GET")
	r.HandleFunc("/instructors/filter", handlers.FindInstructors).Methods("GET")
	r.HandleFunc("/tourists/trainer-instructor", handlers.FindTouristsWithTrainerInstructor).Methods("GET")
	r.HandleFunc("/tourists/completed-all", handlers.FindTouristsCompletedAll).Methods("GET")
	r.HandleFunc("/tourists/completed", handlers.FindTouristsCompletedRoutes).Methods("GET")

	r.HandleFunc("/persons/roles", handlers.GetPersonRole).Methods("GET")
	r.HandleFunc("/persons/roles", handlers.SetPersonRole).Methods("POST")
	r.HandleFunc("/persons/roles", handlers.DeletePersonRole).Methods("DELETE")

	r.HandleFunc("/roles/list", handlers.GetAllRoles).Methods("GET")

	r.HandleFunc("/persons/attribute", handlers.CreatePersonAttribute).Methods("POST")
	r.HandleFunc("/persons/attribute", handlers.GetPersonAttribute).Methods("GET")
	r.HandleFunc("/persons/attribute", handlers.SetPersonAttribute).Methods("PATCH")
	r.HandleFunc("/persons/attribute", handlers.DeletePersonAttribute).Methods("DELETE")

	r.HandleFunc("/groups/group", handlers.CreateGroup).Methods("POST")
	r.HandleFunc("/groups/group", handlers.GetGroup).Methods("GET")
	r.HandleFunc("/groups/group", handlers.UpdateGroup).Methods("PATCH")
	r.HandleFunc("/groups/group", handlers.DeleteGroup).Methods("DELETE")

	r.HandleFunc("/groups/members", handlers.GetGroupMembers).Methods("GET")
	r.HandleFunc("/groups/members/add", handlers.AddGroupMember).Methods("POST")
	r.HandleFunc("/groups/members/remove", handlers.RemoveGroupMember).Methods("DELETE")

	r.HandleFunc("/groups/list", handlers.GetAllGroups).Methods("GET")

	r.HandleFunc("/sections/section", handlers.GetSection).Methods("GET")
	r.HandleFunc("/sections/section", handlers.CreateSection).Methods("POST")
	r.HandleFunc("/sections/section", handlers.UpdateSection).Methods("PATCH")
	r.HandleFunc("/sections/section", handlers.DeleteSection).Methods("DELETE")
	r.HandleFunc("/sections/list", handlers.GetAllSections).Methods("GET")
	r.HandleFunc("/sections/groups", handlers.GetGroupsFromSections).Methods("GET")

	//listen
	addr := fmt.Sprintf(":%s", listenPort)
	if err := http.ListenAndServe(addr, h); err != nil {
		Logger.Fatal(err)
	}
}
