package controller

import (
	"admin/app/utils/result"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	c.JSON(http.StatusOK, result.SUCCESS.WithMsg("hello"))
}
