package sqliteInit

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func insertDepartment(db *sql.DB, name string) {
	log.Println("Inserting department record ...")
	insertStudentSQL := `INSERT INTO department(name) VALUES (?)`
	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(name)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
