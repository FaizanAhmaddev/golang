package api

import (
	"calculator/models"
	"calculator/db"
	"fmt"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
)

// @Summary get all items in the todo list
// @ID get-all-todos
// @Produce json
// @Success 200 {object} models.Response
// @Router /calculator/add [post]
func AddCalculation(ctx *gin.Context) {

	var body models.AddCal
	var afterCal models.Calculation
	var resp models.Responsed

	err := ctx.Bind(&body)

	stringOperation := body.Operation

	var result int 
	
	switch stringOperation {
	case "*":
		result = body.FirstValue * body.SecondValue
	case "+":
		result = body.FirstValue + body.SecondValue
	case "-":
		result = body.FirstValue - body.SecondValue
	case "/":
		if body.SecondValue == 0 {
            ctx.JSON(http.StatusBadRequest, gin.H{"error": "Division by zero"})
            return
        }
        result = body.FirstValue / body.SecondValue
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported operation"})
		return
	}

	afterCal.Result = result
	afterCal.FirstValue = "4"
	afterCal.SecondValue = "3"
	afterCal.ID = "1"
	afterCal.Operation = "/"

	db.AddCalculation(afterCal)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "email already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "user added successfully", Data: resp})
}

func DeleteCalculation(ctx *gin.Context) {
	var body models.AddCal
	var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := ctx.Bind(&body)

	copier.Copy(&user, body)

	if err == nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "already deleted")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, models.Response{Message: "Data deleted successfully"})
}

func GetCalculation(ctx *gin.Context) {
	var body models.AddCal
	// var user models.Calculation
	var resp models.Responsed
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := ctx.Bind(&body)

	// copier.Copy(&user, body)
	
	// userDb, err := c.DB.GetCalculation(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "data already exists")
			return
		}	

		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	
	// copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "Get Data successfully", Data: resp})

}


func UpdateCalculation(ctx *gin.Context) {
	var body models.AddCal
	// var user models.Calculation

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}
	err := ctx.Bind(&body)

	// copier.Copy(&user, body)
	
	// body:= 
	// stringOperation := strconv.Itoa(body.Operation)
	stringOperation := body.Operation

	var result int
	switch stringOperation {
	case "*":
		result = body.FirstValue * body.SecondValue
	case "+":
		result = body.FirstValue + body.SecondValue
	case "-":
		result = body.FirstValue - body.SecondValue
	case "/":
		result = body.FirstValue / body.SecondValue
	default:
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Unsupported operation"})
		return
	}
	fmt.Println("result:					", result)
	// user.result =result

	// userDb, err := c.DB.AddCalculation(user)

	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			ctx.JSON(http.StatusConflict, "email already exists")
			return
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	// copier.Copy(&resp, userDb)

	ctx.JSON(http.StatusOK, models.Response{Message: "user added successfully"})

}
