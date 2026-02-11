package handler

import (
	"Valentine-Proposal-Site-Backend/internals/models"
	"Valentine-Proposal-Site-Backend/internals/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const HOSTURL = "http://localhost:8080/"

type URLHandler struct {
	service *service.PropsalService
}

func NewURLHandler(s *service.PropsalService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var req models.ProposalDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		fmt.Println("❌ Bind error:", err.Error())

		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	fmt.Println("✅ Request received:", req)

	link := h.service.CreateURL(req)

	c.JSON(200, gin.H{
		"link": HOSTURL + link,
	})
}

func (h *URLHandler) GetProposalData(c *gin.Context) {
	code := c.Param("code")

	data, ok := h.service.GetProposal(code)
	if !ok {
		c.JSON(http.StatusNotFound, gin.H{"error": "short URL not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
