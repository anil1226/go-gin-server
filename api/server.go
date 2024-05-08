package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const HTTP_SERVER_ADDRESS = ":8090"

func Start() error {

	router := gin.Default()

	store := Store{}

	router.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		user, err := store.GetUser(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, user)
	})

	router.POST("/users", func(ctx *gin.Context) {
		var req User
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err := store.CreateUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	router.PUT("/users", func(ctx *gin.Context) {
		var req User
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(http.StatusBadRequest, err)
			return
		}
		err := store.UpdateUser(req)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	router.DELETE("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		err := store.DeleteUser(name)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, nil)
			return
		}
		ctx.JSON(http.StatusOK, nil)
	})

	log.Printf("Starting server at %v", HTTP_SERVER_ADDRESS)
	err := router.Run(HTTP_SERVER_ADDRESS)
	if err != nil {
		return err
	}

	return nil
}
