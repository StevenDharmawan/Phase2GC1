package main

import (
	"GC1/config"
	"GC1/handlers"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kelseyhightower/envconfig"
	"net/http"
	"os"
)

func main() {
	var dbEnv config.DatabaseEnv
	err := godotenv.Load(".env")
	if err != nil {
		panic(err)
	}
	err = envconfig.Process("DATABASE", &dbEnv)
	database := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbEnv.Username, dbEnv.Password, dbEnv.Hostname, dbEnv.Port, dbEnv.DBName)
	config.ConnectDB(database)
	defer config.DB.Close()
	mux := http.NewServeMux()
	mux.HandleFunc("GET /employees", handlers.GetAll)
	mux.HandleFunc("POST /employees", handlers.Create)
	mux.HandleFunc("GET /employees/{id}", handlers.GetById)
	mux.HandleFunc("PUT /employees/{id}", handlers.Update)
	mux.HandleFunc("DELETE /employees/{id}", handlers.Delete)
	fmt.Println("Running Server on Port:8081")
	err = http.ListenAndServe(":"+os.Getenv("PORT"), mux)
	if err != nil {
		fmt.Println("Error Starting Server:", err.Error())
	}
}
