package handlers

import (
	"GoProjects/models"
	"GoProjects/services"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

type HouseHandlers struct {
	HouseService *services.HouseService
}

func NewHouseHandlers(houseService *services.HouseService) *HouseHandlers {
	return &HouseHandlers{HouseService: houseService}
}

func (hh *HouseHandlers) GetAllHouses(w http.ResponseWriter, r *http.Request) {
	houses, err := hh.HouseService.GetAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(houses)
}

func (hh *HouseHandlers) GetOneHouseById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	house, err := hh.HouseService.GetById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(house)
}

func (hh *HouseHandlers) CreateOneHouse(w http.ResponseWriter, r *http.Request) {
	var house models.House
	if err := json.NewDecoder(r.Body).Decode(&house); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createdHouse, err := hh.HouseService.Create(&house)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(createdHouse)
}

func (hh *HouseHandlers) UpdateOneHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var house models.House
	if err := json.NewDecoder(r.Body).Decode(&house); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	house.ID = id
	if err := hh.HouseService.Update(&house); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (hh *HouseHandlers) DeleteOneHouse(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	if err := hh.HouseService.Delete(id); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
