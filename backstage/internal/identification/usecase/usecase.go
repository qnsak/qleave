package usecase

import (
	"leaveManager/domain"
	"leaveManager/internal/jwtToken"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type loginUsecase struct {
	UserRepository domain.UserRepository
	// contextTimeout time.Duration
}

func NewLoginUsecase(userRepository domain.UserRepository) domain.LoginUsecase {
	return &loginUsecase{
		UserRepository: userRepository,
	}
}

func (lu *loginUsecase) CreateUser(lr domain.CreateUserRequest) error {
	err := lu.UserRepository.CreateUser(domain.User{
		Name:         lr.Name,
		Email:        lr.Email,
		Password:     lr.Password,
		DepartmentID: lr.DepartmentID,
		Salary:       lr.Salary,
	})

	if err != nil {
		return err
	}

	return err
}

func (lu *loginUsecase) GetUserInfo(userId string) (domain.GetUserInfoResponse, error) {
	userInfo, err := lu.UserRepository.GetUserInfo(userId)

	if err != nil {
		return domain.GetUserInfoResponse{}, err
	}

	return domain.GetUserInfoResponse{
		Name:            userInfo.Name,
		Email:           userInfo.Email,
		DepartmentTitle: userInfo.DepartmentTitle,
		PaidLeaveTotal:  userInfo.PaidLeaveTotal,
		PaidLeaveUse:    userInfo.PaidLeaveUse,
	}, nil
}

func (lu *loginUsecase) Login(lr domain.LoginRequest) (domain.User, error) {
	user, err := lu.UserRepository.FetchEmployeeInformation(lr.Email)
	log.Println(user)
	if err != nil {
		log.Println(err)
		return domain.User{}, err
	}

	checkPasswordError := lu.CheckPasswordHash(lr.Password, user.Password)

	if checkPasswordError != nil {
		return domain.User{}, checkPasswordError
	}

	return user, err
}

func (lu *loginUsecase) GetPasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func (lu *loginUsecase) CheckPasswordHash(password string, passwordHash string) error {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err
}

func (lu *loginUsecase) CreateAccessToken(user domain.User) (accessToken string, err error) {
	return jwtToken.GenerateToken(user)
}

func (lu *loginUsecase) CreateRefreshToken(user domain.User) (refreshToken string, err error) {
	return jwtToken.CreateRefreshToken()
}
