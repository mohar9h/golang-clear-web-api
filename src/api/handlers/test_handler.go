package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/mohar9h/golang-clear-web-api/api/helpers"
	"net/http"
)

type header struct {
	UserId  string `json:"userId"`
	Browser string `json:"browser"`
	Mobile  string `json:"mobile" binding:"required,mobile"`
}

type TestHandler struct{}

func NewTestHandler() *TestHandler {
	return &TestHandler{}
}

func (h *TestHandler) Test(c *gin.Context) {

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(gin.H{
		"result": "Test",
	}, true, 0))
}

func (h *TestHandler) TestById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(gin.H{
		"result": "Test",
		"id":     id,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(gin.H{
		"result":  "Test",
		"user_id": userId,
	}, true, 0))
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	err := c.BindHeader(&header)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest,
			helpers.GenerateBaseResponseWithError(nil, false, -1, err))
		return
	}

	c.JSON(http.StatusOK, helpers.GenerateBaseResponse(gin.H{
		"result": "Test",
		"header": header,
	}, true, 0))
}
