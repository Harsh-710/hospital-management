package patient

import (
	"encoding/json"
	"net/http"
)

// Patient struct
type Patient struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
}

// GetPatients returns a list of patients
func GetPatients(w http.ResponseWriter, r *http.Request) {
	patients := []Patient{
		{ID: 1, Name: "John Doe", Age: 25},
		{ID: 2, Name: "Jane Doe", Age: 30},
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(patients)
}