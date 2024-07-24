package api

import (
	"github.com/gin-gonic/gin"
	"fmt"
)

func Routes(r *gin.Engine) {
	fmt.Println("Routes called")
	r.POST("calculator/add", AddCalculation)
	r.PUT("calculator/update/:id", UpdateCalculation)
	r.GET("calculator/get/:id", GetCalculation)
	r.DELETE("calculator/delete/:id", DeleteCalculation)
}