package sqliteInit

import (
	"database/sql"
	"log"
)

type PaidLeave struct {
	EmployeeId string
	Total      int
	Use        int
	Year       string
}

func createPaidLeave(db *sql.DB) {
	paidLeaves := []PaidLeave{
		{
			EmployeeId: "I0001",
			Total:      18,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "I0001",
			Total:      18,
			Use:        18,
			Year:       "2021",
		},
		{
			EmployeeId: "I0001",
			Total:      18,
			Use:        18,
			Year:       "2022",
		},
		{
			EmployeeId: "I0002",
			Total:      16,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "I0003",
			Total:      12,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "I0004",
			Total:      15,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "I0005",
			Total:      14,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "I0006",
			Total:      18,
			Use:        18,
			Year:       "2022",
		},
		{
			EmployeeId: "M0001",
			Total:      18,
			Use:        0,
			Year:       "2022",
		},
		{
			EmployeeId: "M0001",
			Total:      18,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "M0002",
			Total:      12,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "S0001",
			Total:      18,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "S0002",
			Total:      14,
			Use:        0,
			Year:       "2023",
		},
		{
			EmployeeId: "S0003",
			Total:      12,
			Use:        0,
			Year:       "2023",
		},
	}

	for _, pl := range paidLeaves {
		insertPaidLeave(db, pl)
	}
}

func insertPaidLeave(db *sql.DB, pl PaidLeave) {
	log.Println("Inserting department record ...")

	insertStudentSQL := `
	INSERT INTO paid_leave
	(employee_id, total, use, year)
	VALUES
	(?, ?, ?, ?)`

	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		pl.EmployeeId,
		pl.Total,
		pl.Use,
		pl.Year,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
