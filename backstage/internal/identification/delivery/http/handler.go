package http

import (
	"database/sql"
	"fmt"
	"leaveManager/internal/identification/usecase"
	util "leaveManager/pkg/utils"
	"leaveManager/pkg/validator"
	"log"
	"net/http"

	"leaveManager/domain"
	sqlModal "leaveManager/infrastructure/sqlite"
	sqlcRepo "leaveManager/internal/identification/repository/sqlc"

	"github.com/gin-gonic/gin"
)

type loginController struct {
	LoginUsecase domain.LoginUsecase
}

func NewLoginController(db *sql.DB) *loginController {
	q := sqlModal.New(db)

	ur := sqlcRepo.NewSqliteRepository(q)

	uc := usecase.NewLoginUsecase(ur)

	return &loginController{
		LoginUsecase: uc,
	}
}

func (lc *loginController) Login(c *gin.Context) {
	var request domain.LoginRequest
	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ValidatorErrorResponse{
			Message: validator.Translate(err),
		})
		return
	}
	employee, err := lc.LoginUsecase.Login(request)
	if err != nil {
		log.Println(err)
		util.APIErrorResponse(c, http.StatusInternalServerError, 1001, "err GenerateToken")
		return
	}
	userJwtData := domain.User{
		ID:       employee.ID,
		Name:     employee.Name,
		Password: employee.Password,
	}
	accessToken, err := lc.LoginUsecase.CreateAccessToken(userJwtData)
	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1002, "err GenerateToken")
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(userJwtData)

	if err != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1003, "err GenerateToken")
		return
	}

	loginResponse := domain.LoginResponse{
		ID:              employee.ID,
		Name:            employee.Name,
		Email:           employee.Email,
		DepartmentID:    employee.DepartmentID,
		DirectorID:      employee.DirectorID,
		DepartmentTitle: employee.DepartmentTitle,
		AccessToken:     accessToken,
		RefreshToken:    refreshToken,
	}

	util.APIResponse(c, "Update student data sucessfully", http.StatusOK, loginResponse, nil, nil)
}

func (lc *loginController) CreateUser(c *gin.Context) {
	var request domain.CreateUserRequest

	err := c.ShouldBind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	errCu := lc.LoginUsecase.CreateUser(request)
	fmt.Println(errCu)
	if errCu != nil {
		util.APIErrorResponse(c, http.StatusInternalServerError, 1004, "err GenerateToken")
		return
	}

	util.APIResponse(c, "Update student data sucessfully", http.StatusOK, nil, nil, nil)
}

func (lc *loginController) GetUserInfo(c *gin.Context) {
	userId := c.GetString("x-user-id")
	userInfoResponse, err := lc.LoginUsecase.GetUserInfo(userId)

	if err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	util.APIResponse(c, "sucessfully", http.StatusOK, userInfoResponse, nil, nil)
}
