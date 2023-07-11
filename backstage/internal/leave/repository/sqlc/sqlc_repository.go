package sqlcRepo

import (
	"context"
	"database/sql"
	"log"
	"time"

	"leaveManager/domain"
	sqlRepository "leaveManager/infrastructure/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

type leaveRepository struct {
	SqlRepository *sqlRepository.Queries
}

func NewSqliteRepository(sqlRepository *sqlRepository.Queries) *leaveRepository {
	return &leaveRepository{
		SqlRepository: sqlRepository,
	}
}

func (ur *leaveRepository) GetLeaveByUserId(userId string, page int64) (domain.GetLeavePage, error) {
	ctx := context.Background()

	employeeID := sql.NullString{
		String: userId,
		Valid:  true,
	}

	leaveTotal, leaveTotalErr := ur.SqlRepository.GetTotalLeaveListByEmployeeId(ctx, employeeID)

	if leaveTotalErr != nil {
		return domain.GetLeavePage{}, leaveTotalErr
	}

	var pageLimit int64 = 10
	statrAt := page * pageLimit

	leavedata, err := ur.SqlRepository.GetLeaveListByUserId(ctx, sqlRepository.GetLeaveListByUserIdParams{
		EmployeeID: employeeID,
		Limit:      pageLimit,
		Offset:     statrAt,
	})

	var leaveList []domain.Leave

	for _, l := range leavedata {
		ld := domain.Leave{
			ID:      l.ID,
			Status:  l.Status,
			TypeID:  l.TypeID,
			Reason:  l.Reason,
			StartAt: l.StartAt.Time,
			EndAt:   l.EndAt.Time,
		}
		leaveList = append(leaveList, ld)
	}

	if err != nil {
		return domain.GetLeavePage{}, err
	}

	return domain.GetLeavePage{
		Data:  leaveList,
		Total: leaveTotal,
	}, nil

}

// 取得核假資料
func (ur *leaveRepository) GetLeaveByDirectorId(directorID string, page int64) (domain.GetLeavePage, error) {
	ctx := context.Background()

	id := sql.NullString{
		String: directorID,
		Valid:  true,
	}

	leavesTotal, getLeavetotalErr := ur.SqlRepository.GetTotalLeaveListByDirectorId(ctx, id)

	if getLeavetotalErr != nil {
		return domain.GetLeavePage{}, getLeavetotalErr
	}

	var pageLimit int64 = 10
	statrAt := page * pageLimit

	leavedata, err := ur.SqlRepository.GetLeaveListByDirectorId(ctx, sqlRepository.GetLeaveListByDirectorIdParams{
		DirectorID: id,
		Limit:      pageLimit,
		Offset:     statrAt,
	})

	var leaveList []domain.Leave

	for _, l := range leavedata {
		var eID string = ""
		if l.EmployeeID.Valid {
			eID = l.EmployeeID.String
		}

		var dID string = ""
		if l.DirectorID.Valid {
			dID = l.DirectorID.String
		}

		ld := domain.Leave{
			ID:         l.ID,
			Status:     l.Status,
			TypeID:     l.TypeID,
			LeaveTitle: l.LeaveTitle,
			EmployeeID: eID,
			DirectorID: dID,
			Reason:     l.Reason,
			Name:       l.Name,
			StartAt:    l.StartAt.Time,
			EndAt:      l.EndAt.Time,
		}
		leaveList = append(leaveList, ld)
	}

	if err != nil {
		return domain.GetLeavePage{}, err
	}

	return domain.GetLeavePage{
		Data:  leaveList,
		Total: leavesTotal,
	}, nil
}

func (ur *leaveRepository) CreateLeave(l domain.Leave) error {
	ctx := context.Background()

	eID := sql.NullString{
		String: l.EmployeeID,
		Valid:  true,
	}

	dID := sql.NullString{
		String: l.DirectorID,
		Valid:  true,
	}

	_, err := ur.SqlRepository.CreateLeave(ctx, sqlRepository.CreateLeaveParams{
		TypeID:     l.TypeID,
		Status:     0,
		EmployeeID: eID,
		DirectorID: dID,
		Reason:     l.Reason,
		StartAt:    sql.NullTime{Time: l.StartAt, Valid: true},
		EndAt:      sql.NullTime{Time: l.EndAt, Valid: true},
		DockRate:   sql.NullInt64{Int64: 10, Valid: true},
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *leaveRepository) UpdateLeave(leaveId int64) error {
	ctx := context.Background()

	err := ur.SqlRepository.UpdateLeave(ctx, sqlRepository.UpdateLeaveParams{
		Status: 1,
		ID:     leaveId,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *leaveRepository) GetLeaveById(leaveId int64) (domain.Leave, error) {
	ctx := context.Background()

	leave, err := ur.SqlRepository.GetLeaveRequisitionById(ctx, leaveId)

	if err != nil {
		return domain.Leave{}, err
	}

	var timeStar time.Time
	if leave.StartAt.Valid {
		timeStar = leave.StartAt.Time
	}

	var timeEnd time.Time
	if leave.EndAt.Valid {
		timeEnd = leave.EndAt.Time
	}

	return domain.Leave{
		StartAt: timeStar,
		EndAt:   timeEnd,
	}, nil
}

func (ur *leaveRepository) CreateLeaveRecord(id int64, startTime string, endTime string) error {
	ctx := context.Background()

	startAt, _ := time.Parse("2006-01-02 15:04:05", startTime)

	st := sql.NullTime{
		Time:  startAt,
		Valid: true,
	}

	endAt, _ := time.Parse("2006-01-02 15:04:05", endTime)

	et := sql.NullTime{
		Time:  endAt,
		Valid: true,
	}

	_, err := ur.SqlRepository.CreateLeaveSuccesss(ctx, sqlRepository.CreateLeaveSuccesssParams{
		RequisitionID: id,
		StartAt:       st,
		EndAt:         et,
	})

	if err != nil {
		return err
	}

	return nil
}

func (ur *leaveRepository) DeleteLeave(leaveId int64) error {
	ctx := context.Background()

	err := ur.SqlRepository.DeleteLeave(ctx, leaveId)

	if err != nil {
		return err
	}

	return nil
}

func (ur *leaveRepository) GetAttendanceRecord(departmentID int64) ([]domain.AttendanceRecord, error) {
	ctx := context.Background()

	currentTime := time.Now()
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	endTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location())

	searchParms := sqlRepository.GetAttendanceRecordParams{
		DepartmentID: departmentID,
		StartAt: sql.NullTime{
			Time:  startTime,
			Valid: true,
		},
		EndAt: sql.NullTime{
			Time:  endTime,
			Valid: true,
		},
	}

	attendanceRecord, err := ur.SqlRepository.GetAttendanceRecord(ctx, searchParms)

	if err != nil {
		return []domain.AttendanceRecord{}, err
	}

	var employeeAttendanceList []domain.AttendanceRecord
	for _, l := range attendanceRecord {
		var attendance bool
		if l.Reason == "" {
			attendance = true
		} else {
			attendance = false
		}

		employeeAttendance := domain.AttendanceRecord{
			EmployeeName: l.Name,
			Attendance:   attendance,
		}
		employeeAttendanceList = append(employeeAttendanceList, employeeAttendance)
	}

	return employeeAttendanceList, nil
}

func (ur *leaveRepository) GetEmployeeListByDepartment(departmentId int64) ([]domain.User, error) {
	ctx := context.Background()

	employeeRepository, err := ur.SqlRepository.GetEmployeeListByDepartment(ctx, departmentId)

	if err != nil {
		return []domain.User{}, err
	}

	var employeeList []domain.User
	for _, l := range employeeRepository {
		employeeAttendance := domain.User{
			ID:           l.ID,
			Name:         l.Name,
			Email:        l.Email,
			Password:     l.Password,
			DepartmentID: l.DepartmentID,
		}
		employeeList = append(employeeList, employeeAttendance)
	}

	return employeeList, nil
}

func (ur *leaveRepository) GetLeavesIsSuccess(userId string) ([]domain.Leave, error) {
	ctx := context.Background()

	eID := sql.NullString{
		String: userId,
		Valid:  true,
	}

	startAt := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	leavedata, err := ur.SqlRepository.GetLeavesIsSuccess(ctx, sqlRepository.GetLeavesIsSuccessParams{
		EmployeeID: eID,
		StartAt:    startAt,
	})

	if err != nil {
		return []domain.Leave{}, err
	}

	var leaveList []domain.Leave
	log.Println(leavedata)
	for _, l := range leavedata {
		ld := domain.Leave{
			LeaveTitle: l.LeaveTitle,
			Reason:     l.Reason,
			StartAt:    l.StartAt.Time,
			EndAt:      l.EndAt.Time,
		}
		leaveList = append(leaveList, ld)
	}

	return leaveList, nil
}

func (ur *leaveRepository) GetLeaveType() ([]domain.LeaveType, error) {
	ctx := context.Background()
	leaveType, err := ur.SqlRepository.GetLeaveTypes(ctx)
	if err != nil {
		return nil, err
	}
	var leaveTypeList []domain.LeaveType
	for _, t := range leaveType {
		leaveTypeList = append(
			leaveTypeList,
			domain.LeaveType{
				Id:    t.ID,
				Title: t.Name,
			},
		)
	}

	return leaveTypeList, nil
}

func (ur *leaveRepository) GetNationalHoliday() ([]domain.NationalHoliday, error) {
	ctx := context.Background()
	leaveData, err := ur.SqlRepository.GetNationalHoliday(ctx)
	if err != nil {
		return nil, err
	}
	var nationalHoliday []domain.NationalHoliday
	for _, l := range leaveData {
		var timeStar string = ""
		if l.StartAt.Valid {
			timeStar = l.StartAt.Time.Format("2006-01-02")
		}

		var timeEnd string = ""
		if l.EndAt.Valid {
			timeEnd = l.EndAt.Time.Format("2006-01-02")
		}

		nationalHoliday = append(
			nationalHoliday,
			domain.NationalHoliday{
				Tilte:    l.Reason,
				StarTime: timeStar,
				EndTime:  timeEnd,
			},
		)
	}

	return nationalHoliday, nil
}

func (ur *leaveRepository) GetLeaveByStartDate(departmentId int64, timeStar time.Time, timeAfter time.Time) ([]domain.Leave, error) {
	ctx := context.Background()

	tStar := sql.NullTime{
		Time:  timeStar,
		Valid: true,
	}

	tEnd := sql.NullTime{
		Time:  timeAfter,
		Valid: true,
	}

	leaveData, err := ur.SqlRepository.GetLeaveByStartDate(
		ctx,
		sqlRepository.GetLeaveByStartDateParams{
			DepartmentID: departmentId,
			StartAt:      tStar,
			StartAt_2:    tEnd,
		},
	)
	log.Println(departmentId, timeStar, timeAfter, leaveData)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var leave []domain.Leave
	for _, l := range leaveData {
		var timeStar time.Time
		if l.StartAt.Valid {
			timeStar = l.StartAt.Time
		}

		leave = append(
			leave,
			domain.Leave{
				Name:    l.Name,
				StartAt: timeStar,
			},
		)
	}

	return leave, nil
}
