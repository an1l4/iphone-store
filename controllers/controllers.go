package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/an1l4/iphoneshop/models"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type controller struct {
	DB *gorm.DB
}

func New(db *gorm.DB) controller {
	return controller{db}
}

func (c controller) GetAllIphones(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var iphone []models.Iphone

	if result := c.DB.Find(&iphone); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(iphone)

}

func (c controller) GetIphoneById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var iphone models.Iphone

	if result := c.DB.First(&iphone, id); result.Error != nil {
		fmt.Println(result.Error)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(iphone)
}

func (c controller) CreateNewIphone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal("unable to read json body")
	}

	var iphone models.Iphone
	json.Unmarshal(body, &iphone)

	if result := c.DB.Create(&iphone); result.Error != nil {
		fmt.Println(result.Error)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode("Created")

}

func (c controller) UpdateIphone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var updatedIphone models.Iphone
	json.Unmarshal(body, &updatedIphone)

	var iphone models.Iphone

	if result := c.DB.First(&iphone, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	iphone.Name = updatedIphone.Name
	iphone.Model = updatedIphone.Model
	iphone.Feature = updatedIphone.Feature

	c.DB.Save(&iphone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Updated")

}

func (c controller) DeleteIphone(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	id, _ := strconv.Atoi(params["id"])

	var iphone models.Iphone

	if result := c.DB.First(&iphone, id); result.Error != nil {
		fmt.Println(result.Error)
	}

	c.DB.Delete(&iphone)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("Deleted")

}
