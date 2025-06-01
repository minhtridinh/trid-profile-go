package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/minhtridinh/trid-profile-go/internal/model"
	"github.com/minhtridinh/trid-profile-go/internal/service"
	"gorm.io/gorm"
	"net/http"
	"strconv"
	"trid-profile-go/internal/repository"
)

type ProfileHandler struct {
	service *service.ProfileService
}

func NewProfileHandler(db interface{}) *ProfileHandler {
	repo := repository.NewProfileRepository(db.(*gorm.DB))
	svc := service.NewProfileService(repo)
	return &ProfileHandler{service: svc}
}

func (h *ProfileHandler) GetProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	profile, err := h.service.GetProfile(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Profile not found"})
		return
	}
	c.JSON(http.StatusOK, profile)
}

func (h *ProfileHandler) CreateProfile(c *gin.Context) {
	var profile model.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.service.CreateProfile(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, profile)
}

func (h *ProfileHandler) UpdateProfile(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var profile model.Profile
	if err := c.ShouldBindJSON(&profile); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	profile.ID = uint(id)
	if err := h.service.UpdateProfile(&profile); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, profile)
}
