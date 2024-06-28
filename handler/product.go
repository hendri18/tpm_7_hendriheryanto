package handler

import (
	"fmt"
	"net/http"
	"strconv"
	"tpm_7_HendriHeryanto/models"
	"tpm_7_HendriHeryanto/service"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)

type ProductHandler struct {
	ProductService *service.ProductService
}

func (p *ProductHandler) Get(ctx *gin.Context) {
	products, err := p.ProductService.Get()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   products,
	})
}

func (p *ProductHandler) Create(ctx *gin.Context) {

	productCreate := &models.Product{}
	if err := ctx.Bind(productCreate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}
	userData := ctx.MustGet("userData").(jwt.MapClaims)
	userID := uint64(userData["id"].(float64))

	product, err := p.ProductService.Create(&models.Product{
		Name:   productCreate.Name,
		Price:  productCreate.Price,
		UserID: userID,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (p *ProductHandler) Update(ctx *gin.Context) {

	idx := ctx.Param("id")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	productUpdate := &models.Product{}
	if err := ctx.Bind(productUpdate); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	product, err := p.ProductService.Update(uint64(id), &models.Product{
		Name:  productUpdate.Name,
		Price: productUpdate.Price,
	})

	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   product,
	})
}

func (p *ProductHandler) Delete(ctx *gin.Context) {

	idx := ctx.Param("id")

	if idx == "" {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "ID not found",
		})
		return
	}
	id, _ := strconv.Atoi(idx)

	err := p.ProductService.Delete(uint64(id))
	fmt.Printf("%v %T\n", err, err)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"status":  "error",
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   nil,
	})

}
