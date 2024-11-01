package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"student/models"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected!")


	rows, err := db.Query("SELECT * FROM students")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()	

	var students []models.Student

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.BirthYear, &student.Group)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s %s, Birth Year: %d, Group: %s\n",
			student.Id, student.FirstName, student.LastName, student.BirthYear, student.Group)
	}
}