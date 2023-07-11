package http

import (
	"database/sql"
	"leaveManager/internal/leave/usecase"
	util "leaveManager/pkg/utils"
	"log"
	"net/http"
	"strconv"
	"time"

	"leaveManager/domain"
	sqlModal "leaveManager/infrastructure/sqlite"
	sqlcRepo "leaveManager/internal/leave/repository/sqlc"

	"github.com/gin-gonic/gin"
)

type leaveController struct {
	LeaveUsecase domain.LeaveUsecase
}

type LeaveConfiguration func(os *leaveController) error

func NewLeaveController(cfgs ...LeaveConfiguration) (*leaveController, error) {
	os := &leaveController{}
	for _, cfg := range cfgs {
		err := cfg(os)
		if err != nil {
			return nil, err
		}
	}
	return os, nil
}

func WithSqliteRepository(db *sql.DB) LeaveConfiguration {
	q := sqlModal.New(db)

	ur := sqlcRepo.NewSqliteRepository(q)

	uc := usecase.NewLeaveUsecase(ur)

	return func(os *leaveController) error {
		os.LeaveUsecase = uc
		return nil
	}
}

// 取得所有假
func (lc *leaveController) GetLeaveList(c *gin.Context) {
	userId := c.GetString("x-user-id")
	pageNow, _ := strconv.Atoi(c.DefaultQuery("page", "0"))
	leaveList, err := lc.LeaveUsecase.GetLeaveList(userId, int64(pageNow))

	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1006, "err GenerateToken")
		return
	}

	var leaveListData []domain.LeaveListData

	for _, l := range leaveList.Data {
		ld := domain.LeaveListData{
			ID:         l.ID,
			Status:     l.Status,
			TypeID:     l.TypeID,
			EmployeeID: l.EmployeeID,
			DirectorID: l.DirectorID,
			Reason:     l.Reason,
			StartAt:    l.StartAt,
			EndAt:      l.EndAt,
		}
		leaveListData = append(leaveListData, ld)
	}

	responce := domain.GetLeaveListResponse{
		Meta:  leaveList.Meta,
		Data:  leaveListData,
		Links: leaveList.Links,
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, responce, nil)
}

// 請假
func (lc *leaveController) ApplyLeave(c *gin.Context) {
	userId := c.GetString("x-user-id")
	var request domain.ApplyLeaveRequest

	err := c.ShouldBind(&request)
	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1007, "err GenerateToken")
		return
	}

	err = lc.LeaveUsecase.ApplyLeave(userId, request)

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1008, "err GenerateToken")
		return
	}

	util.APIResponse(c, "Update student data sucessfully", http.StatusOK, nil, nil)
}

// 取消請假
func (lc *leaveController) CancelLeave(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
	}

	id := int64(idP)

	err = lc.LeaveUsecase.CancelLeave(id)

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1009, "err GenerateToken")
		return
	}

	util.APIResponse(c, "Update student data sucessfully", http.StatusOK, nil, nil)
}

// 核准請假
func (lc *leaveController) ApproveLeave(c *gin.Context) {
	var request domain.ApproveLeaveRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}
	err = lc.LeaveUsecase.ApproveLeave(request)

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1010, "err GenerateToken")
		return
	}

	util.APIResponse(c, "Update student data sucessfully", http.StatusOK, nil, nil)
}

// 取得需要核準的假
func (lc *leaveController) GetApproveList(c *gin.Context) {
	userId := c.GetString("x-user-id")

	pageQ, err := strconv.Atoi(c.DefaultQuery("page", "0"))

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1011, "err GenerateToken")
	}

	page := int64(pageQ)

	leaveList, err := lc.LeaveUsecase.GetApproveList(userId, page)
	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1012, "err GenerateToken")
		return
	}

	from := strconv.FormatInt(leaveList.Meta.From, 10)
	to := strconv.FormatInt(leaveList.Meta.To, 10)
	total := strconv.FormatInt(leaveList.Meta.Total, 10)

	meta := map[string]string{
		"from":  from,
		"to":    to,
		"total": total,
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, leaveList.Data, meta)
}

// 每日請假清單
func (lc *leaveController) GetAttendanceRecord(c *gin.Context) {
	idP, err := strconv.Atoi(c.Param("departmentId"))

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1013, "err GenerateToken")
	}

	departmentId := int64(idP)

	starAt := c.Query("starAt")
	var currentTime time.Time
	if len(starAt) == 0 {
		currentTime = time.Now()
	} else {
		var starAtError error
		currentTime, starAtError = time.Parse("2006-01-02", starAt)

		if starAtError != nil {
			util.APIErrorResponse(c, http.StatusInternalServerError, 1014, "日期格式不正確")
			return
		}
	}

	leaveList, err := lc.LeaveUsecase.GetAttendanceRecord(departmentId, currentTime)

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1014, "err GenerateToken")
		return
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, leaveList, nil)
}

// 取得未到的假
func (lc *leaveController) GetLeaveIsComing(c *gin.Context) {
	id := c.GetString("x-user-id")
	leaveList, err := lc.LeaveUsecase.GetLeaveIsComing(id)

	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1015, "err GenerateToken")
		return
	}
	var leaveListData []domain.LeaveComingList
	for _, l := range leaveList.Data {
		ld := domain.LeaveComingList{
			LeaveTitle: l.LeaveTitle,
			Reason:     l.Reason,
			StartAt:    l.StartAt,
			EndAt:      l.EndAt,
		}
		leaveListData = append(leaveListData, ld)
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, leaveListData, nil)
}

func (lc *leaveController) GetLeaveType(c *gin.Context) {

	leaveType, err := lc.LeaveUsecase.GetLeaveType()

	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1016, "err GenerateToken")
		return
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, leaveType, nil)
}

// 取得假單資訊
func (lc *leaveController) GetApplyLeaveInfo(c *gin.Context) {
	applyLeaveInfo, err := lc.LeaveUsecase.GetApplyLeaveInfo()

	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1017, "err GenerateToken")
		return
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, applyLeaveInfo, nil)
}
