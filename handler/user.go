package handler

import (
	"net/http"
	"tpm_7_HendriHeryanto/helpers"
	"tpm_7_HendriHeryanto/models"
	"tpm_7_HendriHeryanto/service"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService *service.UserService
}

func (u *UserHandler) Login(ctx *gin.Context) {
	userLogin := &models.User{}

	if err := ctx.Bind(userLogin); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	user, err := u.UserService.GetByEmail(userLogin.Email)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	comparePass := helpers.CheckPasswordHash(userLogin.Password, user.Password)

	if !comparePass {
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"status":  "error",
			"message": "invalid email or password",
		})
		return
	}

	token, _ := helpers.GenerateUserJWT(user.ID, user.Email)

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"token": token,
		},
	})
}

func (u *UserHandler) Register(ctx *gin.Context) {
	userRegister := &models.User{}

	if err := ctx.Bind(userRegister); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	user, err := u.UserService.Create(userRegister)

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data": gin.H{
			"id":    user.ID,
			"email": user.Email,
		},
	})
}
