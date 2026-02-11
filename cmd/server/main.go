package main

import (
	"Valentine-Proposal-Site-Backend/internals/handler"
	"Valentine-Proposal-Site-Backend/internals/service"
	"Valentine-Proposal-Site-Backend/internals/store"

	"github.com/gin-gonic/gin"
)

func main() {
	memStore := store.NewMemoryStore()
	r := gin.Default()

	PropsalService := service.NewPropsalService(memStore)
	urlHandler := handler.NewURLHandler(PropsalService)

	r.POST("/", urlHandler.CreateURL)
	r.GET("/:code", urlHandler.GetProposalData)

	r.Run(":8080")
}
