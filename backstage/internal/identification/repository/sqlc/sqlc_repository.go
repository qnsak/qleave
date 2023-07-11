package sqlcRepo

import (
	"context"

	"leaveManager/domain"
	sqlRepository "leaveManager/infrastructure/sqlite"

	_ "github.com/mattn/go-sqlite3"
)

type userRepository struct {
	SqlRepository *sqlRepository.Queries
}

// 初始化
func NewSqliteRepository(sqlRepository *sqlRepository.Queries) *userRepository {
	return &userRepository{
		SqlRepository: sqlRepository,
	}
}

// 取得員工資料
func (ur *userRepository) FetchEmployeeInformation(email string) (domain.User, error) {
	ctx := context.Background()

	employee, err := ur.SqlRepository.GetUserByEmail(ctx, email)

	if err != nil {
		return domain.User{}, err
	}

	var directorID string
	if employee.DirectorID.Valid {
		directorID = employee.DirectorID.String
	}

	return domain.User{
		ID:              employee.ID,
		Name:            employee.Name,
		Email:           employee.Email,
		Password:        employee.Password,
		DepartmentID:    employee.DepartmentID,
		DirectorID:      directorID,
		DepartmentTitle: employee.DepartmentTitle,
	}, nil
}

// 創建新的帳號
func (ur *userRepository) CreateUser(user domain.User) error {
	ctx := context.Background()
	_, err := ur.SqlRepository.CreateUser(ctx, sqlRepository.CreateUserParams{
		Name:         user.Name,
		Email:        user.Email,
		Password:     user.Password,
		DepartmentID: user.DepartmentID,
		Salary:       user.Salary,
	})

	if err != nil {
		return err
	}

	return nil
}

// 取得員工資料
func (ur *userRepository) GetUserInfo(userId string) (domain.User, error) {
	ctx := context.Background()
	user, err := ur.SqlRepository.GetUserInfo(ctx, sqlRepository.GetUserInfoParams{
		ID:   userId,
		Year: "2023",
	})

	if err != nil {
		return domain.User{}, err
	}

	return domain.User{
		Name:            user.Name,
		Email:           user.Email,
		DepartmentTitle: user.DepartmentTitle,
		PaidLeaveTotal:  user.PaidLeaveTotal,
		PaidLeaveUse:    user.PaidLeaveUse,
	}, nil
}
