package handlers

import (
	"GC1/config"
	"GC1/entity"
	"encoding/json"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	var employees []entity.Employee
	query := "SELECT * FROM employees"
	rows, err := config.DB.Query(query)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	defer rows.Close()
	for rows.Next() {
		var employee entity.Employee
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Phone, &employee.CreatedAt, &employee.UpdatedAt)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		employees = append(employees, employee)
	}
	response := entity.Response{
		Code:     http.StatusOK,
		Message:  "Success get all employees",
		Employee: employees,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetById(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	var employee entity.Employee
	query := "SELECT * FROM employees WHERE id = ?"
	err := config.DB.QueryRow(query, id).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Phone, &employee.CreatedAt, &employee.UpdatedAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	response := entity.Response{
		Code:     http.StatusOK,
		Message:  "Success get employees by id",
		Employee: employee,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
