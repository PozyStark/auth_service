package api

import (
	"auth-service/internal/models"
	"auth-service/internal/schemas"
	"auth-service/pkg/cryptography"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/samborkent/uuidv7"
)

func RegistrationHandler(cryptParams *cryptography.CryptParams) gin.HandlerFunc {

	op := "RegistrationHandler"

	return func (ctx *gin.Context)  {
			
		var registrationSchema schemas.RegistrationSchema

		if err := ctx.ShouldBind(&registrationSchema); err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusBadRequest,
				gin.H{
					"msg":   "not registered",
					"error": err.Error(),
				},
			)
			return
		}

		newUser := models.User{
			ID:          uuidv7.New().String(),
			Username:    registrationSchema.Username,
			PhoneNumber: registrationSchema.PhoneNumber,
			Email:       registrationSchema.Email,
			FirstName:   registrationSchema.FirstName,
			MiddleName:  registrationSchema.MiddleName,
			LastName:    registrationSchema.LastName,
			Password:    cryptography.HashPassword(registrationSchema.Password, *cryptParams),
			Age:         registrationSchema.Age,
		}

		err := UserService.Save(&newUser)
		if err != nil {
			Logger.Error(fmt.Sprintf("Operation: %s Error: %v", op, err.Error()))
			ctx.JSON(
				http.StatusInternalServerError,
				gin.H{
					"msg":   "registration not successful",
				},
			)
			return
		}

		ctx.JSON(
			http.StatusCreated,
			gin.H{
				"msg": "registration successful",
			},
		)

	}

}
