package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/saeede-bellefille/national-code/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Request struct {
	ID string `json:"id"`
}

type Response struct {
	Valid bool   `json:"valid"`
	City  string `json:"city"`
}
type ErrorResponse struct {
	Message string `json:"message"`
}

func start(address string) error {
	r := gin.Default()
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		v1.GET("/check", check)
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.POST("/check", check)
	return r.Run(address)
}

// check godoc
// @Summary : check national code
// @Description do check national code
// @Tags national code
// @Accept json
// @Produce json
// @Success 200 {Request} {Response}
// @Router /check [get]
func check(c *gin.Context) {
	var request Request
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Message: "invalid input"})
		return
	}
	isValid := validate(request.ID)
	var ct string
	if isValid {
		ct = city(request.ID)
	}
	c.JSON(http.StatusOK, Response{
		Valid: isValid,
		City:  ct,
	})
}

func validate(id string) bool {
	if len(id) != 10 {
		return false
	}
	sum := 0
	for i := 0; i < 9; i++ {
		num, err := strconv.Atoi(string(id[i]))
		if err != nil {
			return false
		}
		sum += num * (10 - i)
	}
	checksum := sum % 11
	if checksum >= 2 {
		checksum = 11 - checksum
	}
	num, err := strconv.Atoi(string(id[9]))
	if err != nil {
		return false
	}
	if num != checksum {
		return false
	}
	return true
}

func city(id string) string {
	index := id[0:3]
	return cities[index]
}
