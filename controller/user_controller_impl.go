package controller

import (
	"net/http"
	"strconv"

	"cek/auth"
	"cek/helper"
	"cek/model/web"
	"cek/service"

	"github.com/gin-gonic/gin"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func NewUserController(userService service.UserService) UserController {
	return &UserControllerImpl{
		UserService: userService,
	}
}

func (controller *UserControllerImpl) RefreshToken(c *gin.Context) {
	request := web.RefreshTokenCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	userAgent := c.GetHeader("User-Agent")
	remoteAddress := c.Request.RemoteAddr

	userResponses := controller.UserService.RefreshToken(&request, &userAgent, &remoteAddress)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(userResponses),
		Data:    userResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Login(c *gin.Context) {
	nip, password, _ := c.Request.BasicAuth()
	userAgent := c.GetHeader("User-Agent")
	remoteAddress := c.Request.RemoteAddr

	tokenResponse := controller.UserService.Login(&nip, &password, &userAgent, &remoteAddress)
	webResponse := web.WebResponse{
		Success: true,
		Message: "OK",
		Data:    tokenResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Create(c *gin.Context, _ *auth.AccessDetails) {
	request := &web.UserCreateRequest{}
	helper.ReadFromRequestBody(c, &request)

	userResponse := controller.UserService.Create(request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User created successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) ChangePassword(c *gin.Context, auth *auth.AccessDetails) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	request := &web.UserChangePasswordRequest{}
	helper.ReadFromRequestBody(c, &request)

	userResponse := controller.UserService.ChangePassword(auth, &userID, request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User change password successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) ResetPassword(c *gin.Context, auth *auth.AccessDetails) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	userResponse := controller.UserService.ResetPassword(auth, &userID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User reset password successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Delete(c *gin.Context, _ *auth.AccessDetails) {
	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	controller.UserService.Delete(&userID)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User deleted successfully",
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) Update(c *gin.Context, _ *auth.AccessDetails) {
	request := &web.UserUpdateRequest{}
	helper.ReadFromRequestBody(c, &request)

	id := c.Param("id")
	userID, err := strconv.Atoi(id)
	helper.PanicIfError(err)

	userResponse := controller.UserService.Update(&userID, request)
	webResponse := web.WebResponse{
		Success: true,
		Message: "User updated successfully",
		Data:    userResponse,
	}

	c.JSON(http.StatusOK, webResponse)
}

func (controller *UserControllerImpl) FindAll(c *gin.Context, _ *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "dept.eq", "nip.eq")

	userResponses := controller.UserService.FindAll(&filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(userResponses),
		Data:    userResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}
