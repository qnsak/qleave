package domain

import (
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Leave struct {
	ID         int64     `json:"id"`
	Status     int64     `json:"status"`
	TypeID     int64     `json:"typeId"`
	LeaveTitle string    `json:"typeTitle"`
	EmployeeID string    `json:"employeeId"`
	DirectorID string    `json:"directorId"`
	Reason     string    `json:"reason"`
	Name       string    `json:"name"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
	DockRate   int64     `json:"dockRate"`
}

type LeaveType struct {
	Id    int64  `json:"id"`
	Title string `json:"title"`
}

type AttendanceRecord struct {
	EmployeeName string
	Attendance   bool
}

type GetLeaveIsComing struct {
	Data []Leave `json:"data"`
}

type LeaveUsecase interface {
	GetLeaveList(employeeID string, page int64) (GetLeave, error)
	ApplyLeave(string, ApplyLeaveRequest) error
	CancelLeave(LeaveId int64) error
	ApproveLeave(ApproveLeaveRequest) error
	GetApproveList(employeeID string, page int64) (GetApproveListResponse, error)
	GetAttendanceRecord(departmentId int64, currentTime time.Time) (map[string][]LeaveByDay, error)
	GetLeaveIsComing(id string) (GetLeaveIsComing, error)
	GetLeaveType() ([]LeaveType, error)
	GetApplyLeaveInfo() (ApplyLeaveInfoResponse, error)
}

type LeaveRepository interface {
	GetLeaveByUserId(userId string, page int64) (GetLeavePage, error)
	GetLeaveByDirectorId(userId string, page int64) (GetLeavePage, error)
	CreateLeave(Leave) error
	UpdateLeave(leaveId int64) error
	DeleteLeave(leaveId int64) error
	GetAttendanceRecord(departmentId int64) ([]AttendanceRecord, error)
	GetEmployeeListByDepartment(departmentId int64) ([]User, error)
	GetLeavesIsSuccess(userId string) ([]Leave, error)
	GetLeaveType() ([]LeaveType, error)
	GetNationalHoliday() ([]NationalHoliday, error)
	GetLeaveByStartDate(departmentId int64, timeStar time.Time, timeAfter time.Time) ([]Leave, error)
	GetLeaveById(int64) (Leave, error)
	CreateLeaveRecord(int64, string, string) error
}

type GetLeavePage struct {
	Data  []Leave
	Total int64
}

type NationalHoliday struct {
	Tilte    string `json:"title"`
	StarTime string `json:"starTime"`
	EndTime  string `json:"endTime"`
}

type ApplyLeaveInfoResponse struct {
	NationalHolidays []NationalHoliday `json:"nationalHoliday"`
	LeaveTypes       []LeaveType       `json:"leaveType"`
}
type LeaveByDay struct {
	EmployeeName string `json:"name"`
	StarTime     string `json:"date"`
}

type CancelLeaveRequest struct {
	ID int64 `form:"id" binding:"required"`
}

type ApproveLeaveRequest struct {
	ID int64 `form:"id" binding:"required"`
}

type ApplyLeaveRequest struct {
	TypeID     int64     `form:"typeId" binding:"required"`
	DirectorID string    `form:"directorID" binding:"required"`
	Reason     string    `form:"reason" binding:"required"`
	StartAt    time.Time `form:"startAt" binding:"required"`
	EndAt      time.Time `form:"endAt" binding:"required"`
}

type GetApproveListRequest struct {
	Email string `form:"Email" binding:"required,email"`
}

type Page struct {
	From  int64 `json:"from"`
	To    int64 `json:"to"`
	Total int64 `json:"total"`
}

type Links struct {
	Prev int64 `json:"prev"`
	Next int64 `json:"next"`
}

type GetApproveListResponse struct {
	Meta  Page    `json:"meta"`
	Data  []Leave `json:"data"`
	Links Links   `json:"links"`
}

type LeaveListData struct {
	ID         int64     `json:"id"`
	Status     int64     `json:"status"`
	TypeID     int64     `json:"typeId"`
	EmployeeID string    `json:"employeeId"`
	DirectorID string    `json:"directorId"`
	Reason     string    `json:"reason"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
}

type GetLeave struct {
	Meta  Page    `json:"meta"`
	Data  []Leave `json:"data"`
	Links Links   `json:"links"`
}

type GetLeaveListResponse struct {
	Meta  Page            `json:"meta"`
	Data  []LeaveListData `json:"data"`
	Links Links           `json:"links"`
}

type LeaveComingList struct {
	LeaveTitle string    `json:"leaveTitle"`
	Reason     string    `json:"reason"`
	StartAt    time.Time `json:"startAt"`
	EndAt      time.Time `json:"endAt"`
}
