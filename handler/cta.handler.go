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
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   "Method not allowed. Use POST.",
		})
		return
	}

	var cta models.Cta

	err := json.NewDecoder(r.Body).Decode(&cta)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(Response{Success: false, Error: "Invalid Input Data! Cannot Process Request."})
		return
	}

	err2 := validator.ValidateCta(cta)
	if err2 != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(Response{
			Success: false,
			Error:   err2.Error(),
		})
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(Response{
		Success: true,
		Message: "Request Processed Successfully!",
		Data:    cta,
	})
}
