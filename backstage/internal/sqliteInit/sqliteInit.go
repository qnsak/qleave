package sqliteInit

import (
	"database/sql"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	msqlite "github.com/golang-migrate/migrate/v4/database/sqlite3"
	"github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
)

func createDb() {
	os.Remove("sqlite-database.db")

	log.Println("Creating sqlite-database.db...")
	newdb, err := os.Create("sqlitedb.db")
	if err != nil {
		log.Fatal(err.Error())
	}
	newdb.Close()

	db, err := sql.Open("sqlite3", "./sqlitedb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	instance, err := msqlite.WithInstance(db, &msqlite.Config{})
	if err != nil {
		log.Fatal(err)
	}

	fSrc, err := (&file.File{}).Open("./database/migrations")
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.NewWithInstance("file", fSrc, "sqlite3", instance)
	if err != nil {
		log.Fatal(err)
	}

	if err := m.Up(); err != nil {
		log.Fatal(err)
	}

}

func InitDatabase() {
	createDb()
	sqliteDatabase, _ := sql.Open("sqlite3", "./sqlitedb.db")
	defer sqliteDatabase.Close()
	insertDepartment(sqliteDatabase, "資訊部")
	insertDepartment(sqliteDatabase, "人資部")
	insertDepartment(sqliteDatabase, "業務部")
	createAccount(sqliteDatabase)
	createTypeLeave(sqliteDatabase)
	createPaidLeave(sqliteDatabase)
	initLeaveRequisitionData(sqliteDatabase)
	log.Println("init database complete.")
}
