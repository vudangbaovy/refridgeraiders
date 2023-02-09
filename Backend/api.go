package main

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"errors"
)

type Server struct {
	dataBase	*gorm.DB
	router  	*gin.Engine
}

type Account struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
	Allergies string `json:"Allergies" binding:"required"`
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}

func (server *Server) createAccount(ctx *gin.Context) {
	// user validation
	var req Account
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// need to update the add use function to accept diff parameters
	userAdded, account := newUserProfile(req.Name, req.Password, req.Allergies, server.dataBase)
	if !userAdded {
		err := errors.New("usernameTaken")
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func routerSetup (db *gorm.DB) *Server {

	server := &Server{dataBase: db}
	router1 := gin.Default()

	router1.POST("/accounts", server.createAccount)

	server.router = router1

	return server
}

func (server *Server) Start(address string) error {
    return server.router.Run(address)
}

func serverStart(db *gorm.DB) {
	server := routerSetup(db)
	err := server.Start("localhost:8080")
	if err != nil {
		fmt.Println("cannot start server:", err)
	}
}
