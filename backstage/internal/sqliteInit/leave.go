package sqliteInit

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

type Leave struct {
	RequisitionId int
	StartAt       string
	EndAt         string
}

func GetBetweenDates(sdate, edate string) []string {
	d := []string{}
	timeFormatTpl := "2006-01-02 15:04:05"
	if len(timeFormatTpl) != len(sdate) {
		timeFormatTpl = timeFormatTpl[0:len(sdate)]
	}
	date, err := time.Parse(timeFormatTpl, sdate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	date2, err := time.Parse(timeFormatTpl, edate)
	if err != nil {
		// 时间解析，异常
		return d
	}
	if date2.Before(date) {
		// 如果结束时间小于开始时间，异常
		return d
	}
	// 输出日期格式固定
	timeFormatTpl = "2006-01-02"
	date2Str := date2.Format(timeFormatTpl)
	d = append(d, date.Format(timeFormatTpl))
	for {
		date = date.AddDate(0, 0, 1)
		dateStr := date.Format(timeFormatTpl)
		d = append(d, dateStr)
		if dateStr == date2Str {
			break
		}
	}
	return d
}

func createLeave(db *sql.DB, id int, dataS string, dataE string) {
	dates := GetBetweenDates(dataS, dataE)

	for _, d := range dates {
		startAt := fmt.Sprintf("%s 00:00:00", d)
		endAt := fmt.Sprintf("%s 23:59:59", d)
		leaves := Leave{
			RequisitionId: id,
			StartAt:       startAt,
			EndAt:         endAt,
		}

		insertLeave(db, leaves)

	}
	log.Println(dates)
}

func insertLeave(db *sql.DB, l Leave) {
	log.Println("Inserting leave record ...")

	insertStudentSQL := `
	INSERT INTO leave
	(requisition_id, start_at, end_at)
	VALUES
	(?, ?, ?)`

	statement, err := db.Prepare(insertStudentSQL)
	if err != nil {
		log.Fatalln(err.Error())
	}
	_, err = statement.Exec(
		l.RequisitionId,
		l.StartAt,
		l.EndAt,
	)

	if err != nil {
		log.Fatalln(err.Error())
	}
}
