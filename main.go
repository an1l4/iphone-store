package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/an1l4/iphoneshop/controllers"
	"github.com/an1l4/iphoneshop/db"
	"github.com/gorilla/mux"
)

func main() {
	DB := db.Init()
	c := controllers.New(DB)
	router := mux.NewRouter()

	router.HandleFunc("/iphones", c.GetAllIphones).Methods("GET", "OPTIONS")
	router.HandleFunc("/iphone/{id}", c.GetIphoneById).Methods("GET", "OPTIONS")
	router.HandleFunc("/newiphone", c.CreateNewIphone).Methods("POST", "OPTIONS")
	router.HandleFunc("/iphone/{id}", c.UpdateIphone).Methods("PUT", "OPTIONS")
	router.HandleFunc("/iphone/{id}", c.DeleteIphone).Methods("DELETE", "OPTIONS")

	fmt.Println("server running at 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))

}
