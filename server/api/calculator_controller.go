package api
// "github.com/go-playground/validator/v10"

import (
	"database/sql"
	"fmt"
	"net/http"
	"parking/lib/utils"
	"parking/models"
	"strings"

	"github.com/gin-gonic/gin"
	// "github.com/jinzhu/copier"
)

func AddCalculation(ctx *gin.Context) {
	var body models.AddCalculation
	var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	copier.Copy(&user, body)
	
	// reqBody:= 
var result int
	if body.Operation == "*" {
		result = reqBody.firstvalue * reqBody.secondvalue

	}
	else if body.Operation == "+" {
		result = reqBody.firstvalue + reqBody.secondvalue

	}
	else if body.Operation == "-" {
		result = reqBody.firstvalue - reqBody.secondvalue

	}
	else if body.Operation == "/" {
		result = reqBody.firstvalue / reqBody.secondvalue

	}
	else {
		http.Error(w, "Unsupported operation", http.StatusBadRequest)
	}

	user.result :=result

	userDb, err := c.DB.AddCalculation(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "email already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user added successfully"})

}

func DeleteCalculation(ctx *gin.Context) {
	var body models.AddCalculation
	var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	copier.Copy(&user, body)

	userDb, err := c.DB.DeleteCalculation(user)

	if err == nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "already deleted")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "Data deleted successfully"})

}


func GetCalculation(ctx *gin.Context) {
	var body models.AddCalculation
	var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	copier.Copy(&user, body)
	
	userDb, err := c.DB.GetCalculation(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "data already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "Get Data successfully", Data: resp})

}


func UpdateCalculation(ctx *gin.Context) {
	var body models.AddCalculation
	var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	copier.Copy(&user, body)
	
	// reqBody:= 
	var result int
	if body.Operation == "*" {
		result = reqBody.firstvalue * reqBody.secondvalue

	}
	else if body.Operation == "+" {
		result = reqBody.firstvalue + reqBody.secondvalue

	}
	else if body.Operation == "-" {
		result = reqBody.firstvalue - reqBody.secondvalue

	}
	else if body.Operation == "/" {
		result = reqBody.firstvalue / reqBody.secondvalue

	}
	else {
		http.Error(w, "Unsupported operation", http.StatusBadRequest)
	}

	user.result :=result

	userDb, err := c.DB.AddCalculation(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "email already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user added successfully"})

}
