package handlers

import (
	"GC1/config"
	"GC1/entity"
	"encoding/json"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var employee entity.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "INSERT INTO employees(name, email, phone) VALUES (?, ?, ?)"
	result, err := config.DB.Exec(query, employee.Name, employee.Email, employee.Phone)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	resultId, err := result.LastInsertId()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	employee.Id = int(resultId)
	response := entity.Response{
		Code:     http.StatusOK,
		Message:  "Success create employees",
		Employee: employee,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
