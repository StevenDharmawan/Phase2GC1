package handlers

import (
	"GC1/config"
	"GC1/entity"
	"encoding/json"
	"net/http"
	"strconv"
)

func Update(w http.ResponseWriter, r *http.Request) {
	var employee entity.Employee
	err := json.NewDecoder(r.Body).Decode(&employee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := r.PathValue("id")
	query := "UPDATE employees SET name = ?, email = ?, phone = ?, updated_at = CURRENT_TIMESTAMP() WHERE id = ?"
	result, err := config.DB.Exec(query, employee.Name, employee.Email, employee.Phone, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "No rows affected, data not found or no change made", http.StatusNotFound)
		return
	}
	idInt, _ := strconv.Atoi(id)
	employee.Id = idInt
	response := entity.Response{
		Code:     http.StatusOK,
		Message:  "Success update employees",
		Employee: employee,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
