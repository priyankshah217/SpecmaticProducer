package controllers

import (
	"SpecmaticProducer/models"
	"errors"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	invalidSchemaError     = errors.New("invalid schema")
	invalidQueryParamError = errors.New("invalid query param")
	inventoryLimitError    = errors.New("inventory should be less than 9999")
)

type CreateProductRequest struct {
	Name      string `json:"name" binding:"required"`
	Type      string `json:"type" binding:"required"`
	Inventory int    `json:"inventory" binding:"required"`
	Cost      int    `json:"cost" binding:"required"`
}

type CreateProductResponse struct {
	ID int `json:"id"`
}

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Status    int       `json:"status"`
	Error     string    `json:"error"`
	Path      string    `json:"path"`
}

var productMap = make(map[int]models.Product)

// CreateProduct create product using gin
func CreateProduct(c *gin.Context) {
	var request CreateProductRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(400, ErrorResponse{
			Timestamp: time.Now(),
			Status:    0,
			Error:     invalidSchemaError.Error(),
			Path:      c.Request.URL.Path,
		})
		return
	}
	if request.Inventory > 9999 {
		c.JSON(400, ErrorResponse{
			Timestamp: time.Now(),
			Status:    0,
			Error:     inventoryLimitError.Error(),
			Path:      c.Request.URL.Path})
		return
	}
	count := len(productMap) + 1
	product := models.Product{ID: count, Name: request.Name, Type: request.Type, Inventory: request.Inventory, Cost: request.Cost}
	productMap[count] = product
	c.JSON(201, CreateProductResponse{ID: count})
}

// GetProductsByQuery get products by query
func GetProductsByQuery(c *gin.Context) {
	typeParam := c.Query("type")
	if isNonStringParam(typeParam) {
		c.JSON(400, ErrorResponse{
			Timestamp: time.Now(),
			Status:    0,
			Error:     invalidQueryParamError.Error(),
			Path:      c.Request.URL.Path,
		})
		return
	}
	// if typeParam is not type string then it should return 400
	name := c.Query("name")
	inventory := c.Query("inventory")
	var products []models.Product
	for _, product := range productMap {
		if typeParam != "" && product.Type != typeParam {
			continue
		}
		if name != "" && product.Name != name {
			continue
		}
		if inventory != "" && product.Inventory != 0 {
			continue
		}
		products = append(products, product)
	}
	c.JSON(200, products)
}

func isNonStringParam(param string) bool {
	if _, err := strconv.Atoi(param); err == nil {
		return true
	} else if _, err := strconv.ParseBool(param); err == nil {
		return true
	} else if _, err := strconv.ParseFloat(param, 64); err == nil {
		return true
	}
	return false
}
