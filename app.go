package main

import (
	routes "github.com/HirunikaGunathunga/GO_REST/Routes"
	db "github.com/HirunikaGunathunga/GO_REST/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.MainRoutes(server)
	server.Run(":8080")
}
