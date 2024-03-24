package handlers

import (
	"GC1/config"
	"GC1/entity"
	"encoding/json"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	query := "DELETE FROM employees WHERE id = ?"
	_, err := config.DB.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := entity.Response{
		Code:    http.StatusOK,
		Message: "Success delete employees",
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
