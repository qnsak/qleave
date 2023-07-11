package sqliteInit

import (
	"database/sql"
	"log"
	"time"
)

type Account struct {
	Id            string
	Name          string
	Email         string
	Password      string
	Department_id int32
	Salary        int32
	Director_id   string
	Entry_at      time.Time
}

func createAccount(db *sql.DB) {
	users := []Account{
		{
			Id:            "I0001",
			Name:          "主管一號",
			Email:         "I0001@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        100000,
			Director_id:   "",
			Entry_at:      time.Now(),
		},
		{
			Id:            "I0002",
			Name:          "工程師A",
			Email:         "I0002@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        90000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "I0003",
			Name:          "工程師B",
			Email:         "I0003@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        80000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "I0004",
			Name:          "工程師C",
			Email:         "I0004@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        60000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "I0005",
			Name:          "測試A",
			Email:         "I0005@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        60000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "I0006",
			Name:          "測試B",
			Email:         "I0006@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 1,
			Salary:        60000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "M0001",
			Name:          "人事A",
			Email:         "M0001@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 2,
			Salary:        70000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "M0002",
			Name:          "人事B",
			Email:         "M0002@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 2,
			Salary:        65000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "S0001",
			Name:          "業務1",
			Email:         "S0001@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 3,
			Salary:        50000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "S0002",
			Name:          "業務2",
			Email:         "S0002@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 3,
			Salary:        40000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
		{
			Id:            "S0003",
			Name:          "業務3",
			Email:         "S0003@mail.com",
			Password:      "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
			Department_id: 3,
			Salary:        35000,
			Director_id:   "I0001",
			Entry_at:      time.Now(),
		},
	}

	for _, user := range users {
		insertAccount(db, user)
	}
}

func insertAccount(db *sql.DB, user Account) {
	log.Println("Inserting department record ...")

	insertStudentSQL := `
	INSERT INTO employee
	(id, name, email, password, department_id, salary, director_id, entry_at)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?)`

	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		user.Id,
		user.Name,
		user.Email,
		user.Password,
		user.Department_id,
		user.Salary,
		user.Director_id,
		user.Entry_at,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
