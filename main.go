package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"

)

type Student struct {
	Id   int
	FirstName string
	BirthDate int
}

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

	for rows.Next() {
		var student Student
		err := rows.Scan(&student.Id, &student.FirstName, &student.BirthDate)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(student)
	}	

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}