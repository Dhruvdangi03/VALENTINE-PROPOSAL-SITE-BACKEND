package main

import (
	"Valentine-Proposal-Site-Backend/internals/handler"
	"Valentine-Proposal-Site-Backend/internals/service"
	"Valentine-Proposal-Site-Backend/internals/store"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	memStore := store.NewMemoryStore()
	r := gin.Default()

	// ✅ CORS MUST BE HERE (before routes)
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	PropsalService := service.NewPropsalService(memStore)
	urlHandler := handler.NewURLHandler(PropsalService)

	r.POST("/", urlHandler.CreateURL)
	r.GET("/:code", urlHandler.GetProposalData)

	r.Run(":8080")
}
