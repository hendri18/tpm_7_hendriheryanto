package middleware

import (
	"net/http"
	"strconv"
	"strings"
	"tpm_7_HendriHeryanto/helpers"
	"tpm_7_HendriHeryanto/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		headerAuth := ctx.GetHeader("Authorization")
		// {Authorization: Bearer jwt_token}
		// get the encoded string
		splitToken := strings.Split(headerAuth, " ")
		if len(splitToken) != 2 {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid authorization header",
			})
			return
		}

		// check basic
		if splitToken[0] != "Bearer" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid authorization method",
			})
			return
		}
		// validate jwt
		valid, claims := helpers.ValidateUserJWT(splitToken[1])
		if !valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "malformed token",
			})
			return
		}
		// get claim
		// add value in context
		ctx.Set("userData", claims)
		ctx.Next()
	}
}

func ProductAuthorization(productService *service.ProductService) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idx := ctx.Param("id")

		if idx == "" {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "ID not found",
			})
			return
		}
		id, _ := strconv.Atoi(idx)
		product, err := productService.GetById(uint64(id))

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
				"status":  "error",
				"message": "product not found",
			})
			return
		}

		userData := ctx.MustGet("userData").(jwt.MapClaims)
		userID := uint64(userData["id"].(float64))

		if product.UserID != userID {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"status":  "error",
				"message": "invalid authorization product",
			})
			return
		}
		ctx.Next()
	}
}
