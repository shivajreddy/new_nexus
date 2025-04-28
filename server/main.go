package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var db *sql.DB

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}

func main() {
	var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	psqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("Error pinging database: ", err)
	}

	fmt.Println("Connected to the database!")

	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/employees", employeesHandler)

	port := "8080"
	fmt.Println("Server is running at http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(":::: @ / route ::::")
	fmt.Fprintf(w, "Welcome to the API!")
}

func employeesHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println(":::: @ /employees route ::::")
	rows, err := db.Query("SELECT id, name, position, salary FROM employees")
	if err != nil {
		http.Error(w, "Error querying employees", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var employees []Employee

	for rows.Next() {
		var emp Employee
		err := rows.Scan(&emp.ID, &emp.Name, &emp.Position, &emp.Salary)
		if err != nil {
			http.Error(w, "Error scanning employee", http.StatusInternalServerError)
			return
		}
		employees = append(employees, emp)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

