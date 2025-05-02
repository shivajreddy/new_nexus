package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Employee struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Position string `json:"position"`
	Salary   int    `json:"salary"`
}

func main() {
	t := Task{}
	fmt.Println("t", t)

	var err error
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Build DSN (Data Source Name) for PostgreSQL
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", dbHost, dbPort, dbUser, dbPassword, dbName)

	// Initialize GORM DB connection
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to database: ", err)
	}

	// Auto migrate the Employee struct
	err = db.AutoMigrate(&Employee{})
	if err != nil {
		log.Fatal("Error migrating database: ", err)
	}

	fmt.Println("Connected to the database!")

	// Set up HTTP routes
	http.HandleFunc("/", homeHandler)
	http.HandleFunc("/employees", employeesHandler)
	http.HandleFunc("/new", newEmployeeHandler)

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
	var employees []Employee
	// Retrieve all employees from the database using GORM
	result := db.Find(&employees)
	if result.Error != nil {
		http.Error(w, "Error querying employees", http.StatusInternalServerError)
		return
	}

	// Return employees as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(employees)
}

func newEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	// Get the last employee from the database using GORM
	var lastEmp Employee
	result := db.Order("id desc").First(&lastEmp)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			http.Error(w, "No employees found", http.StatusNotFound)
		} else {
			http.Error(w, "Error fetching last employee", http.StatusInternalServerError)
		}
		return
	}

	// Create the new employee in the database using GORM
	newEmployee := Employee{
		ID:       lastEmp.ID + 1,
		Name:     lastEmp.Name,
		Position: lastEmp.Position,
		Salary:   int(time.Now().Local().Unix()), // Current time in seconds
	}

	result = db.Create(&newEmployee)
	if result.Error != nil {
		http.Error(w, "Error inserting new employee", http.StatusInternalServerError)
		return
	}

	// Return the newly created employee as a JSON response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newEmployee)
}
