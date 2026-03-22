package handler

import (
	"encoding/json"
	"net/http"

	"jatinporiya/models"
	validator "jatinporiya/utils"
)

type Response struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Error   string      `json:"error,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func CreateCTA(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed!", http.StatusMethodNotAllowed)
		return
	}

	var cta models.Cta

	err := json.NewDecoder(r.Body).Decode(&cta)
	if err != nil {
		json.NewEncoder(w).Encode(Response{Success: false, Error: "Invalid Input Data! Cannot Process Request."})
		return
	}

	err2 := validator.ValidateCta(cta)
	if err2 != nil {
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err2.Error(),
		})
		return
	}

	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Request Processed Successfully!",
		Data:    cta,
	})
}
