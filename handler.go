package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func addExoplanetHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var planet Exoplanet
		if err := json.NewDecoder(r.Body).Decode(&planet); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := validateExoplanet(planet); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		planet.ID = generateID() // Assume generateID generates a unique ID
		storage.AddExoplanet(planet)
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(planet)
	}
}

func listExoplanetsHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		planets := storage.GetExoplanets()
		json.NewEncoder(w).Encode(planets)
	}
}

func getExoplanetByIDHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		planet, err := storage.GetExoplanetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(planet)
	}
}

func updateExoplanetHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		var updated Exoplanet
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := validateExoplanet(updated); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated.ID = id
		if err := storage.UpdateExoplanet(id, updated); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		json.NewEncoder(w).Encode(updated)
	}
}

func deleteExoplanetHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		if err := storage.DeleteExoplanet(id); err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		w.WriteHeader(http.StatusNoContent)
	}
}

func calculateFuelHandler(storage *Storage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		planet, err := storage.GetExoplanetByID(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		crewCapacity := r.URL.Query().Get("crew_capacity")
		if crewCapacity == "" {
			http.Error(w, "crew_capacity query parameter is required", http.StatusBadRequest)
			return
		}
		c, err := strconv.Atoi(crewCapacity)
		if err != nil {
			http.Error(w, "invalid crew_capacity", http.StatusBadRequest)
			return
		}
		gravity := calculateGravity(planet)
		fuelCost := float64(planet.Distance) / (gravity * gravity) * float64(c)
		json.NewEncoder(w).Encode(map[string]float64{"fuel_cost": fuelCost})
	}
}

func calculateGravity(planet Exoplanet) float64 {
	if planet.Type == GasGiant {
		return 0.5 / (planet.Radius * planet.Radius)
	} else if planet.Type == Terrestrial {
		return planet.Mass / (planet.Radius * planet.Radius)
	}
	return 0.0
}
