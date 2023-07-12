package sqliteInit

import (
	"database/sql"
	"log"
	"time"
)

type Leave_requisition struct {
	TypeId     int
	Status     int
	Reason     string
	StartAt    time.Time
	EndAt      time.Time
	DockRate   int
	EmployeeId string
	DirectorId string
}

func initLeaveRequisitionData(db *sql.DB) {
	now := time.Now()

	lr := []Leave_requisition{
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0002",
			DirectorId: "I0001",
			Reason:     "補休",
			StartAt:    now.AddDate(0, 0, 3),
			EndAt:      now.AddDate(0, 0, 4),
			DockRate:   0,
		},
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "上班途中迷了路。",
			StartAt:    now.AddDate(0, 0, 1),
			EndAt:      now.AddDate(0, 0, 2),
			DockRate:   0,
		},
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "FF 16！",
			StartAt:    now.AddDate(0, 0, 3),
			EndAt:      now.AddDate(0, 0, 7),
			DockRate:   0,
		},
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0004",
			DirectorId: "I0001",
			Reason:     "尋找台南美食",
			StartAt:    now.AddDate(0, 0, 11),
			EndAt:      now.AddDate(0, 0, 12),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0002",
			DirectorId: "I0001",
			Reason:     "累了，想休息。",
			StartAt:    now.AddDate(0, 0, 17),
			EndAt:      now.AddDate(0, 0, 19),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "補休",
			StartAt:    now.AddDate(0, 0, 30),
			EndAt:      now.AddDate(0, 0, 31),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0004",
			DirectorId: "I0001",
			Reason:     "補休",
			StartAt:    now.AddDate(0, 0, 29),
			EndAt:      now.AddDate(0, 0, 30),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0002",
			DirectorId: "I0001",
			Reason:     "旅遊",
			StartAt:    now.AddDate(0, 0, 40),
			EndAt:      now.AddDate(0, 0, 50),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "回家探親",
			StartAt:    now.AddDate(0, 0, 41),
			EndAt:      now.AddDate(0, 0, 45),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0004",
			DirectorId: "I0001",
			Reason:     "颱風假，我覺得會是颱風天",
			StartAt:    now.AddDate(0, 0, 22),
			EndAt:      now.AddDate(0, 0, 24),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0002",
			DirectorId: "I0001",
			Reason:     "思考人生",
			StartAt:    now.AddDate(0, 0, 12),
			EndAt:      now.AddDate(0, 0, 13),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "旅行",
			StartAt:    now.AddDate(0, 0, 13),
			EndAt:      now.AddDate(0, 0, 14),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0004",
			DirectorId: "I0001",
			Reason:     "家裡有事請。",
			StartAt:    now.AddDate(0, 0, 8),
			EndAt:      now.AddDate(0, 0, 9),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0005",
			DirectorId: "I0001",
			Reason:     "私事",
			StartAt:    now.AddDate(0, 0, 2),
			EndAt:      now.AddDate(0, 0, 3),
			DockRate:   0,
		},
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0005",
			DirectorId: "I0001",
			Reason:     "私事",
			StartAt:    now.AddDate(0, 0, 6),
			EndAt:      now.AddDate(0, 0, 9),
			DockRate:   0,
		},
		{
			Status:     0,
			TypeId:     1,
			EmployeeId: "I0005",
			DirectorId: "I0001",
			Reason:     "私事",
			StartAt:    now.AddDate(0, 0, 2),
			EndAt:      now.AddDate(0, 0, 3),
			DockRate:   0,
		},
		{
			Status:     1,
			TypeId:     1,
			EmployeeId: "I0003",
			DirectorId: "I0001",
			Reason:     "私事",
			StartAt:    now.AddDate(0, 0, 4),
			EndAt:      now.AddDate(0, 0, 6),
			DockRate:   0,
		},
	}

	for _, l := range lr {
		insertLeaveRequisition(db, l)
	}
}

type PublicVacation struct {
	Title string
	DataS string
	DataE string
}

func insertLeaveRequisition(db *sql.DB, l Leave_requisition) {
	log.Println("Inserting Leave_requisition record ...")

	insertStudentSQL := `
	INSERT INTO leave_requisition
	(status, type_id, employee_id, director_id, reason, start_at, end_at, dock_rate)
	VALUES
	(?, ?, ?, ?, ?, ?, ?, ?)
	RETURNING id`

	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	var id int
	StartAt := l.StartAt.Format("2006-01-02 15:04:05")
	EndAt := l.EndAt.Format("2006-01-02 15:04:05")
	err = statement.QueryRow(
		l.Status,
		l.TypeId,
		l.EmployeeId,
		l.DirectorId,
		l.Reason,
		StartAt,
		EndAt,
		l.DockRate,
	).Scan(&id)

	if err != nil {
		log.Fatal(err)
	}
	if l.Status == 1 {
		createLeave(db, id, StartAt, EndAt)
	}
}

func initPublicVacation(db *sql.DB) {
	pv := []PublicVacation{
		{
			Title: "元旦",
			DataS: "2023-01-01 00:00:00",
			DataE: "2023-01-02 23:59:59",
		},
		{
			Title: "農曆春節",
			DataS: "2023-02-25 00:00:00",
			DataE: "2023-01-20 23:59:59",
		},
		{
			Title: "228紀念日",
			DataS: "2023-02-25 00:00:00",
			DataE: "2023-02-28 23:59:59",
		},
		{
			Title: "兒童節及清明節",
			DataS: "2023-04-01 00:00:00",
			DataE: "2023-04-05 23:59:59",
		},
		{
			Title: "五一勞動節",
			DataS: "2023-04-29 00:00:00",
			DataE: "2023-05-01 23:59:59",
		},
		{
			Title: "端午節",
			DataS: "2023-06-22 00:00:00",
			DataE: "2023-06-25 23:59:59",
		},
		{
			Title: "中秋節",
			DataS: "2023-09-29 00:00:00",
			DataE: "2023-10-01 23:59:59",
		},
		{
			Title: "國慶日",
			DataS: "2023-10-07 00:00:00",
			DataE: "2023-10-10 23:59:59",
		},
	}

	for _, v := range pv {
		insertStudentSQL := `
		INSERT INTO leave_requisition
		(status, type_id, employee_id, director_id, reason, start_at, end_at, dock_rate)
		VALUES
		(?, ?, ?, ?, ?, ?, ?, ?)
		RETURNING id`

		statement, err := db.Prepare(insertStudentSQL)
		if err != nil {
			log.Fatalln(err.Error())
		}
		var id int
		err = statement.QueryRow(
			1,
			5,
			nil,
			nil,
			v.Title,
			v.DataS,
			v.DataE,
			0,
		).Scan(&id)

		if err != nil {
			log.Fatal(err)
		}

		createLeave(db, id, v.DataS, v.DataE)
	}
}
