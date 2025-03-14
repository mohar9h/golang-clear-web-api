package handlers

import (
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
	})
}

func (h *TestHandler) TestById(c *gin.Context) {
	id := c.Param("id")
	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
		"id":     id,
	})
}

func (h *TestHandler) HeaderBinder(c *gin.Context) {
	userId := c.GetHeader("UserId")

	c.JSON(http.StatusOK, gin.H{
		"result":  "Test",
		"user_id": userId,
	})
}

func (h *TestHandler) HeaderBinder2(c *gin.Context) {
	header := header{}
	err := c.BindHeader(&header)
	if err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"result": "Test",
		"header": header,
	})
}
