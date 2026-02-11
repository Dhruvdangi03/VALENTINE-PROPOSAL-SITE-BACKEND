package handler

import (
	"Valentine-Proposal-Site-Backend/internals/models"
	"Valentine-Proposal-Site-Backend/internals/service"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

const HOSTURL = "http://localhost:8080"

type URLHandler struct {
	service *service.PropsalService
}

func NewURLHandler(s *service.PropsalService) *URLHandler {
	return &URLHandler{service: s}
}

func (h *URLHandler) CreateURL(c *gin.Context) {
	var body struct {
		Proposal models.ProposalDTO `json:"proposal"`
	}

	if err := c.BindJSON(&body); err != nil {
		fmt.Println("Bind Error:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	fmt.Println("Received:", body.Proposal)

	if body.Proposal.Email == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "email required"})
		return
	}

	url := h.service.CreateURL(body.Proposal)

	c.JSON(http.StatusOK, gin.H{
		"url": HOSTURL + "/" + url,
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
