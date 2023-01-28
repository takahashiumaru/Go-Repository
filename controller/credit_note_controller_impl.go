package controller

import (
	"net/http"

	"cek/auth"
	"cek/helper"
	"cek/model/web"
	"cek/service"

	"github.com/gin-gonic/gin"
)

type TestControllerImpl struct {
	TestService service.TestService
}

func NewTestController(testService service.TestService) TestController {
	return &TestControllerImpl{
		TestService: testService,
	}
}

func (controller *TestControllerImpl) FindAll(c *gin.Context, auth *auth.AccessDetails) {
	filters := helper.FilterFromQueryString(c, "name.eq")
	customerResponses := controller.TestService.FindAll(auth, &filters)
	webResponse := web.WebResponse{
		Success: true,
		Message: helper.MessageDataFoundOrNot(customerResponses),
		Data:    customerResponses,
	}

	c.JSON(http.StatusOK, webResponse)
}
