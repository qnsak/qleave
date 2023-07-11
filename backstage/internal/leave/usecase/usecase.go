package usecase

import (
	"fmt"
	"leaveManager/domain"
	"time"
)

type leaveUsecase struct {
	leaveRepository domain.LeaveRepository
}

func NewLeaveUsecase(lr domain.LeaveRepository) domain.LeaveUsecase {
	return &leaveUsecase{
		leaveRepository: lr,
	}
}

func (lu *leaveUsecase) GetLeaveList(employeeID string, pageNow int64) (domain.GetLeave, error) {
	leaveList, err := lu.leaveRepository.GetLeaveByUserId(employeeID, pageNow)

	if err != nil {
		return domain.GetLeave{}, err
	}

	pageTotal := leaveList.Total / 10

	var next int64 = 0
	if pageNow < pageTotal {
		next = pageNow + 1
	}

	var prev int64 = 0
	if pageNow > 1 {
		prev = pageNow - 1
	}

	from := pageNow*10 + 1

	to := pageNow*10 + int64(len(leaveList.Data))

	meta := domain.Page{
		From:  from,
		To:    to,
		Total: leaveList.Total,
	}

	links := domain.Links{
		Prev: prev,
		Next: next,
	}

	return domain.GetLeave{
		Meta:  meta,
		Data:  leaveList.Data,
		Links: links,
	}, nil
}

func (lu *leaveUsecase) ApplyLeave(userId string, al domain.ApplyLeaveRequest) error {
	err := lu.leaveRepository.CreateLeave(domain.Leave{
		TypeID:     al.TypeID,
		EmployeeID: userId,
		DirectorID: al.DirectorID,
		Reason:     al.Reason,
		StartAt:    al.StartAt,
		EndAt:      al.EndAt,
	})

	if err != nil {
		return err
	}

	return nil
}

func (lu *leaveUsecase) CancelLeave(id int64) error {
	err := lu.leaveRepository.DeleteLeave(id)

	if err != nil {
		return err
	}

	return err
}

// 取得時間區間
func getBetweenDates(sdate string, edate string) []string {
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

func (lu *leaveUsecase) ApproveLeave(al domain.ApproveLeaveRequest) error {
	err := lu.leaveRepository.UpdateLeave(al.ID)

	if err != nil {
		return err
	}

	lr, err := lu.leaveRepository.GetLeaveById(al.ID)

	dataS := lr.StartAt.Format("2006-01-02 15:04:05")
	dataE := lr.EndAt.Format("2006-01-02 15:04:05")

	dates := getBetweenDates(dataS, dataE)
	for _, d := range dates {
		startAt := fmt.Sprintf("%s 00:00:00", d)
		endAt := fmt.Sprintf("%s 23:59:59", d)
		lu.leaveRepository.CreateLeaveRecord(al.ID, startAt, endAt)
	}

	return err
}

func (lu *leaveUsecase) GetApproveList(employeeID string, pageNow int64) (domain.GetApproveListResponse, error) {
	leaveList, err := lu.leaveRepository.GetLeaveByDirectorId(employeeID, pageNow)

	if err != nil {
		return domain.GetApproveListResponse{}, err
	}

	pageTotal := leaveList.Total / 10

	var next int64 = 0
	if pageNow < pageTotal {
		next = pageNow + 1
	}

	var prev int64 = 0
	if pageNow > 1 {
		prev = pageNow - 1
	}

	from := pageNow*10 + 1

	to := pageNow*10 + int64(len(leaveList.Data))

	meta := domain.Page{
		From:  from,
		To:    to,
		Total: leaveList.Total,
	}

	links := domain.Links{
		Prev: prev,
		Next: next,
	}

	return domain.GetApproveListResponse{
		Meta:  meta,
		Data:  leaveList.Data,
		Links: links,
	}, nil
}

func (lu *leaveUsecase) GetAttendanceRecord(departmentId int64, currentTime time.Time) (map[string][]domain.LeaveByDay, error) {
	dayDuring := 10
	startTime := time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 0, 0, 0, 0, currentTime.Location())
	afterTime := currentTime.AddDate(0, 0, dayDuring)
	endTime := time.Date(afterTime.Year(), afterTime.Month(), afterTime.Day(), 23, 59, 59, 0, afterTime.Location())
	testday, err := lu.leaveRepository.GetLeaveByStartDate(departmentId, startTime, endTime)

	if err != nil {
		return nil, err
	}

	var days []string
	days = append(
		days,
		currentTime.Format("2006-01-02"),
	)
	for i := 1; i <= dayDuring; i = i + 1 {
		day := currentTime.AddDate(0, 0, i).Format("2006-01-02")
		days = append(
			days,
			day,
		)
	}

	datalist := make(map[string][]domain.LeaveByDay, 0)
	for _, d := range days {
		var nl []domain.LeaveByDay
		for _, a := range testday {
			leaveDay := a.StartAt.Format("2006-01-02")
			if leaveDay == d {
				nl = append(
					nl,
					domain.LeaveByDay{
						EmployeeName: a.Name,
						StarTime:     leaveDay,
					},
				)
			}
		}
		datalist[d] = nl
	}

	return datalist, nil
}

func (lu *leaveUsecase) GetLeaveIsComing(id string) (domain.GetLeaveIsComing, error) {

	leaves, err := lu.leaveRepository.GetLeavesIsSuccess(id)
	if err != nil {
		return domain.GetLeaveIsComing{}, err
	}

	return domain.GetLeaveIsComing{
		Data: leaves,
	}, nil
}

func (lu *leaveUsecase) GetLeaveType() ([]domain.LeaveType, error) {

	leaveType, err := lu.leaveRepository.GetLeaveType()

	if err != nil {
		return nil, err
	}

	return leaveType, nil
}

func (lu *leaveUsecase) GetApplyLeaveInfo() (domain.ApplyLeaveInfoResponse, error) {
	leaveType, err := lu.leaveRepository.GetLeaveType()

	if err != nil {
		return domain.ApplyLeaveInfoResponse{}, err
	}

	nationalHoliday, getNationalHolidayErr := lu.leaveRepository.GetNationalHoliday()

	if getNationalHolidayErr != nil {
		return domain.ApplyLeaveInfoResponse{}, getNationalHolidayErr
	}

	return domain.ApplyLeaveInfoResponse{
		NationalHolidays: nationalHoliday,
		LeaveTypes:       leaveType,
	}, nil
}
