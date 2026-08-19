package main

import (
	db "github.com/HirunikaGunathunga/GO_REST/Database"
	routes "github.com/HirunikaGunathunga/GO_REST/Routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	routes.MainRoutes(server)
	server.Run(":8080")
}
