package usecase_test

import (
	"leaveManager/domain"
	"leaveManager/internal/identification/usecase"
	"testing"

	"leaveManager/internal/identification/repository/mocks"

	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
)

func Test_user_login(t *testing.T) {
	ctl := gomock.NewController(t)
	defer ctl.Finish()

	mockUserRepository := mocks.NewMockUserRepository(ctl)

	userData := domain.User{
		ID:           "1",
		Name:         "test01",
		Email:        "test@mail.com",
		Password:     "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
		DepartmentID: 1,
	}

	mockUserRepository.EXPECT().FetchEmployeeInformation("test@mail.com").Return(userData, nil)
	mockUserRepository.EXPECT().FetchEmployeeInformation("noacconut@mail.com").Return(domain.User{}, nil)
	uc := usecase.NewLoginUsecase(mockUserRepository)

	data := []struct {
		name   string
		input  domain.LoginRequest
		output domain.User
		errMsg string
	}{
		{
			"happy bath",
			domain.LoginRequest{
				Email:    "test@mail.com",
				Password: "123456",
			},
			domain.User{
				ID:           "1",
				Name:         "test01",
				Email:        "test@mail.com",
				Password:     "$2a$14$CJ.vq9pxTyTSKKQBW.UCMO7gkxXtRP4Wcft4ZUtii8te0kD99X7uS",
				DepartmentID: 1,
			},
			"",
		},
		{
			"bad bath",
			domain.LoginRequest{
				Email:    "noacconut@mail.com",
				Password: "123456",
			},
			domain.User{},
			"",
		},
	}

	for _, d := range data {
		t.Run(d.name, func(t *testing.T) {
			user, _ := uc.Login(d.input)
			// if err != nil {
			// 	t.Error(err)
			// }
			// cmp.Diff 值對比
			if diff := cmp.Diff(user, d.output); diff != "" {
				t.Error(diff)
			}
			//var errMsg string
			// if err != nil {
			// 	errMsg = err.Error()
			// }
			// if errMsg != d.errMsg {
			// 	t.Errorf("Expected error `%s`, got `%s`", d.errMsg, errMsg)
			// }
		})
	}

}
