package main

import (
	"fmt"
	"log"
	"os"
	"context"

	"student/models"

	"github.com/joho/godotenv"
	"github.com/jackc/pgx/v5"

)
func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	var students []models.Student

	rows, err := conn.Query(context.Background(), 
	"SELECT id, first_name, last_name, birth_year, group_name FROM students")

	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {
		var student models.Student
		err := rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.BirthYear, &student.Group)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, student)
	}
	
	for _, student := range students {
		fmt.Printf("ID: %d, Name: %s %s, Birth Year: %d, Group: %s\n",
			student.Id, student.FirstName, student.LastName, student.BirthYear, student.Group)
	}
}