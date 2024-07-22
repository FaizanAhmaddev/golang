package api

import "github.com/gin-gonic/gin"

func Routes(r *gin.Engine) {
	r.POST("calculator/add", AddCalculation)
	r.PUT("calculator/update/:id", UpdateCalculation)
	r.GET("calculator/get/:id", GetCalculation)
	r.DELETE("calculator/delete/:id", DeleteCalculation)
}