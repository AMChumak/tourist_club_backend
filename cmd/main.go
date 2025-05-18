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
	//listen
	addr := fmt.Sprintf(":%s", listenPort)
	if err := http.ListenAndServe(addr, h); err != nil {
		Logger.Fatal(err)
	}
}
