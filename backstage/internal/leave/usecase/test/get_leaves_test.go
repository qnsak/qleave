package leave_test

import (
	"leaveManager/domain"
	"leaveManager/internal/leave/repository/mocks"
	"leaveManager/internal/leave/usecase"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_get_leave_list(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockLeaveRepository := mocks.NewMockLeaveRepository(ctl)

	leaveList := []domain.Leave{
		{
			ID:         1,
			Status:     2,
			TypeID:     1,
			EmployeeID: "I0002",
			DirectorID: "I0001",
			Reason:     "無",
			StartAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			EndAt:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			ID:         1,
			Status:     2,
			TypeID:     1,
			EmployeeID: "I0002",
			DirectorID: "I0001",
			Reason:     "無",
			StartAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
			EndAt:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
		},
	}

	returnData := domain.GetLeavePage{
		Data:  leaveList,
		Total: 2,
	}
	mockLeaveRepository.EXPECT().GetLeaveByUserId("I0002", int64(0)).Return(returnData, nil)

	mockLeaveRepository.EXPECT().GetLeaveByUserId("I0003", int64(0)).Return(domain.GetLeavePage{}, nil)

	// 初始化
	uc := usecase.NewLeaveUsecase(mockLeaveRepository)

	data := []struct {
		name   string
		input  string
		output domain.GetLeave
		errMsg string
	}{
		{
			"happy path",
			"I0002",
			domain.GetLeave{
				Meta: domain.Page{
					From:  1,
					To:    2,
					Total: 2,
				},
				Data: []domain.Leave{
					{
						ID:         1,
						Status:     2,
						TypeID:     1,
						EmployeeID: "I0002",
						DirectorID: "I0001",
						Reason:     "無",
						StartAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						EndAt:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:         1,
						Status:     2,
						TypeID:     1,
						EmployeeID: "I0002",
						DirectorID: "I0001",
						Reason:     "無",
						StartAt:    time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
						EndAt:      time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC),
					},
				},
				Links: domain.Links{
					Prev: 0,
					Next: 0,
				},
			},
			"",
		},
		{
			"bad path",
			"I0003",
			domain.GetLeave{
				Meta: domain.Page{
					From:  1,
					To:    0,
					Total: 0,
				},
				Data:  nil,
				Links: domain.Links{},
			},
			"",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			leaveLiat, _ := uc.GetLeaveList(d.input, 0)
			// cmp.Diff 值對比
			if diff := cmp.Diff(leaveLiat, d.output); diff != "" {
				t.Error(diff)
			}
		})
	}
}

func Test_get_leave_is_coming_soon(t *testing.T) {}

func Test_approve_leave(t *testing.T) {}

func Test_apply_leave(t *testing.T) {

}
