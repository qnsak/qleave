package sqliteInit

import (
	"database/sql"
	"log"
)

type TypeLeave struct {
	Name     string
	Status   int
	DockRate int
}

func createTypeLeave(db *sql.DB) {
	typeLeaves := []TypeLeave{
		{
			Name:     "公假",
			Status:   1,
			DockRate: 0,
		},
		{
			Name:     "特休",
			Status:   1,
			DockRate: 0,
		},
		{
			Name:     "病假",
			Status:   1,
			DockRate: 50,
		},
		{
			Name:     "榮譽假",
			Status:   1,
			DockRate: 0,
		},
		{
			Name:     "國訂假日",
			Status:   1,
			DockRate: 0,
		},
	}

	for _, tl := range typeLeaves {
		insertTypeLeave(db, tl)
	}
}

func insertTypeLeave(db *sql.DB, tl TypeLeave) {
	log.Println("Inserting department record ...")

	insertStudentSQL := `
	INSERT INTO leave_type
	(name, status, dock_rate)
	VALUES
	(?, ?, ?)`

	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		tl.Name,
		tl.Status,
		tl.DockRate,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
