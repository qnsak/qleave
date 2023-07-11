package domain

import (
	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID              string
	Name            string
	Email           string
	Password        string
	DepartmentID    int64
	DepartmentTitle string
	DirectorID      string
	Salary          int64
	PaidLeaveTotal  int64
	PaidLeaveUse    int64
}

type UserRepository interface {
	CreateUser(User) error
	FetchEmployeeInformation(employeeEmail string) (User, error)
	GetUserInfo(userId string) (User, error)
}

type LoginRequest struct {
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

type LoginResponse struct {
	ID              string `json:"ID,omitempty"`
	Name            string `json:"name,omitempty"`
	Email           string `json:"email,omitempty"`
	DepartmentID    int64  `json:"departmentID,omitempty"`
	DepartmentTitle string `json:"departmentTitle,omitempty"`
	DirectorID      string `json:"directorID,omitempty"`
	AccessToken     string `json:"accessToken,omitempty"`
	RefreshToken    string `json:"refreshToken,omitempty"`
}

type CreateUserRequest struct {
	Name         string `form:"name" binding:"required"`
	Email        string `form:"email" binding:"required,email"`
	Password     string `form:"password" binding:"required"`
	DepartmentID int64  `form:"departmentID" binding:"required"`
	Salary       int64  `form:"salary" binding:"required"`
}

type Response struct {
	ID              string `json:"ID"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	DepartmentID    int64  `json:"departmentID"`
	DepartmentTitle string `json:"departmentTitle"`
	DirectorID      string `json:"directorID"`
	AccessToken     string `json:"accessToken"`
	RefreshToken    string `json:"refreshToken"`
}

type GetUserInfoResponse struct {
	Name            string `json:"name"`
	Email           string `json:"email"`
	DepartmentTitle string `json:"departmentTitle"`
	PaidLeaveTotal  int64  `json:"paidLeaveTotal"`
	PaidLeaveUse    int64  `json:"paidLeaveUse"`
}

type LoginUsecase interface {
	Login(LoginRequest) (User, error)
	CreateAccessToken(User) (string, error)
	CreateRefreshToken(User) (string, error)
	GetPasswordHash(string) (string, error)
	CheckPasswordHash(string, string) error
	CreateUser(CreateUserRequest) error
	GetUserInfo(userId string) (GetUserInfoResponse, error)
}
